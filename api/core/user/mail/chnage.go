package mail

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/logging"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/cateiru/cateiru-sso/api/utils/mail"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

const VERIFY_MAIL_TEMPLATE_PATH = "verify_change_mail"

type ChangeMailRequest struct {
	Type string `json:"type"` // `change` or `verify`

	NewMail string `json:"new_mail"`

	MailToken string `json:"mail_token"`
}

// テンプレートに適用する用の型
type VerifyMailTemplate struct {
	VerifyURL string
	Mail      string
}

func CangeMailHandler(w http.ResponseWriter, r *http.Request) error {
	// contents-type: application/json 以外では400エラーを返す
	if !net.CheckContentType(r) {
		return status.NewBadRequestError(errors.New("requests contets-type is not application/json")).Caller()
	}

	ctx := r.Context()

	var request ChangeMailRequest

	if err := net.GetJsonForm(w, r, &request); err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	db, err := database.NewDatabase(ctx)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	defer db.Close()

	// メールアドレス変更はログイン状態でのみ可
	userId, err := common.GetUserID(ctx, db, w, r)
	if err != nil {
		return err
	}

	switch request.Type {
	case "change":
		// メールアドレス認証のリクエストを送信する
		return ChangeMail(ctx, db, request.NewMail, userId)
	case "verify":
		// メールトークンを使用して自分のアカウントのメールアドレスを変更します
		return VerifyNewMail(ctx, db, request.MailToken, userId)
	default:
		return status.NewBadRequestError(errors.New("parse failed")).Caller()
	}
}

// メールアドレス変更リクエストを受け付けます
func ChangeMail(ctx context.Context, db *database.Database, newMail string, userId string) error {
	if err := createVerifyChangeMail(ctx, db, newMail, userId); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	return nil
}

// メールトークンからメールアドレスを更新する
func VerifyNewMail(ctx context.Context, db *database.Database, token string, userId string) error {
	entity, err := models.GetMailCertificationByMailToken(ctx, db, token)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if entity == nil {
		return status.NewBadRequestError(errors.New("mail cert is empty")).Caller()
	}

	// 違うアカウントで認証しようとしたら400を返す
	if entity.UserId != userId {
		return status.NewBadRequestError(errors.New("bad account")).Caller()
	}

	// 有効期限が切れている場合、400を返す
	if common.CheckExpired(&entity.Period) {
		return status.NewBadRequestError(errors.New("expired")).AddCode(net.TimeOutError).Caller()
	}

	// ---- Certを変更する

	cert, err := models.GetCertificationByUserID(ctx, db, entity.UserId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if cert == nil {
		return status.NewInternalServerErrorError(errors.New("cert is empty")).Caller()
	}

	cert.Mail = entity.Mail // certのメールアドレスを更新

	if err := cert.Add(ctx, db); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	/// ---- UserInfoを変更する

	info, err := models.GetUserDataByUserID(ctx, db, cert.UserId.UserId)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	if info == nil {
		return status.NewInternalServerErrorError(errors.New("user info is empty")).Caller()
	}

	info.Mail = entity.Mail // user infoのメールアドレスを更新

	if err := info.Add(ctx, db); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	return nil
}

// メール認証を開始します
//
// client_check_token(wsを接続するのに使用するトークンを返します)
func createVerifyChangeMail(ctx context.Context, db *database.Database, newMail string, userId string) error {
	mailToken := utils.CreateID(20)

	mailVerify := &models.MailCertification{
		MailToken:        mailToken,
		ClientCheckToken: utils.CreateID(0), // 使わないが一応keyを指定しておく

		OpenNewWindow:  false,
		Verify:         false,
		ChangeMailMode: true, // メールアドレス変更なので

		UserMailPW: models.UserMailPW{
			Mail: newMail,
		},

		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 30,
		},

		UserId: userId,
	}

	if err := mailVerify.Add(ctx, db); err != nil {
		return err
	}

	// send mail
	// SendGrid APIをテストでは使用しないため、
	// DEPLOY_MODEがproductionのときのみ送信します
	if config.Defs.DeployMode == "production" {
		if err := sendVerifyMail(newMail, mailToken); err != nil {
			return err
		}
	} else {
		logging.Sugar.Debugf(
			"create mail token. url: https://%s/mail/change?m=%s", config.Defs.SiteDomain, mailToken)
	}

	return nil
}

// メールアドレス認証メールを送信する
func sendVerifyMail(mailAddress string, mailToken string) error {
	template := VerifyMailTemplate{
		VerifyURL: fmt.Sprintf("https://%s/mail/change?m=%s", config.Defs.SiteDomain, mailToken),
		Mail:      mailAddress,
	}

	mailClient, err := mail.NewMail(
		"", mailAddress, "メールアドレス変更認証のお知らせ").AddContentsFromTemplate(VERIFY_MAIL_TEMPLATE_PATH, template)
	if err != nil {
		return err
	}

	return mailClient.Send()
}
