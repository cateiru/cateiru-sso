package createaccount

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/badoux/checkmail"
	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/logging"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/cateiru/cateiru-sso/api/utils/mail"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/cateiru-sso/api/utils/secure"
	"github.com/cateiru/go-http-error/httperror/status"
)

const VERIFY_MAIL_TEMPLATE_PATH = "verify_mail"

// POSTのformの型
type PostForm struct {
	Mail      string `json:"mail"`
	ReCAPTCHA string `json:"re_captcha"`
}

// レスポンスの型
type Response struct {
	ClientToken string `json:"client_token"`
}

// テンプレートに適用する用の型
type VerifyMailTemplate struct {
	VerifyURL string
	Mail      string
}

// 一時的にアカウントを作成します
// メールアドレス、をfromで送信することで、そのメールアドレスに確認用URLを送信します。
// さらに、Websocketでメールアドレスが認証されたか確認するためのトークンを返します。
//
// Request Form (application/json):
//	{
//		"mail": "example@example.com",
//		"re_chaptcha": "********",
//	}
func CreateTemporaryHandler(w http.ResponseWriter, r *http.Request) error {
	// contents-type: application/json 以外では400エラーを返す
	if !net.CheckContentType(r) {
		return status.NewBadRequestError(errors.New("requests contets-type is not application/json")).Caller()
	}

	postForm := new(PostForm)
	if err := net.GetJsonForm(w, r, postForm); err != nil {
		status.NewBadRequestError(errors.New("parse failed")).Caller()
	}

	ip := net.GetIPAddress(r)
	ctx := r.Context()

	clientToken, err := CreateTemporaryAccount(ctx, postForm, ip)
	if err != nil {
		return err
	}

	response := Response{
		ClientToken: clientToken,
	}

	net.ResponseOK(w, response)

	return nil
}

func CreateTemporaryAccount(ctx context.Context, form *PostForm, ip string) (string, error) {
	// reCAPTCHA
	if config.Defs.DeployMode == "production" {
		isOk, err := secure.NewReCaptcha().Validate(form.ReCAPTCHA, ip)
		if err != nil {
			return "", err
		}
		// reCAPTCHAが認証できなかった場合、400を返す
		if !isOk {
			return "", status.NewBadRequestError(errors.New("reCAPTCHA is failed")).Caller().AddCode(net.BotError)
		}
	}

	db, err := database.NewDatabase(ctx)
	if err != nil {
		return "", status.NewInternalServerErrorError(err).Caller()
	}
	defer db.Close()

	// IPアドレス、メールアドレスがブロックされているか確認
	isBlocked, err := common.ChaeckBlock(ctx, db, ip, form.Mail)
	if err != nil {
		return "", status.NewInternalServerErrorError(err).Caller()
	}
	if isBlocked {
		return "", status.NewForbiddenError(errors.New("ip is blocked")).Caller().AddCode(net.BlockedError)
	}

	// メールアドレスをチェックする
	if err := checkmail.ValidateFormat(form.Mail); err != nil {
		return "", status.NewBadRequestError(err).Caller().AddCode(net.IncorrectMail)
	}

	// メールアドレスが既に存在するかチェック
	isMailExist, err := common.CheckExistMail(ctx, db, form.Mail)
	if err != nil {
		return "", status.NewInternalServerErrorError(err).Caller()
	}
	// メールアドレスがすでに存在している = そのメールアドレスを持ったアカウントが作られている場合、
	// あたらにそのメールアドレスでアカウントを作成することはできないため、403エラーを返す
	if isMailExist {
		return "", status.NewForbiddenError(errors.New("email already exists")).Caller().AddCode(net.IncorrectMail)
	}
	// Adminのメールアドレスは既に定義されており、ログインできるため弾く
	if common.CheckAdminMail(form.Mail) {
		return "", status.NewForbiddenError(errors.New("email is admin")).Caller().AddCode(net.IncorrectMail)
	}

	// ログを保存する
	log := &models.TryCreateAccountLog{
		LogId:      utils.CreateID(0),
		IP:         ip,
		TryDate:    time.Now(),
		TargetMail: form.Mail,
	}
	if err := log.Add(ctx, db); err != nil {
		return "", status.NewInternalServerErrorError(err).Caller()
	}

	clientToken, err := createVerifyMail(ctx, db, form.Mail)
	if err != nil {
		return "", status.NewInternalServerErrorError(err).Caller()
	}

	return clientToken, nil
}

// メール認証を開始します
//
// client_token(wsを接続するのに使用する&認証後ユーザ情報を設定するためのトークンを返します)
func createVerifyMail(ctx context.Context, db *database.Database, mail string) (string, error) {
	mailToken := utils.CreateID(20)
	clientToken := utils.CreateID(20)

	mailVerify := &models.MailCertification{
		MailToken:   mailToken,
		ClientToken: clientToken,

		OpenNewWindow:  false,
		Verify:         false,
		ChangeMailMode: false,

		Mail: mail,

		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 30,
		},
	}

	if err := mailVerify.Add(ctx, db); err != nil {
		return "", err
	}

	// send mail
	// SendGrid APIをテストでは使用しないため、
	// DEPLOY_MODEがproductionのときのみ送信します
	if config.Defs.DeployMode == "production" {
		if err := sendVerifyMail(mail, mailToken); err != nil {
			return "", err
		}
	} else {
		logging.Sugar.Debugf(
			"create mail token. url: https://%s/create?m=%s", config.Defs.SiteDomain, mailToken)
	}

	logging.Sugar.Debugf(
		"Send verify email. mail: %s, client check token: %s", mail, clientToken)

	return clientToken, nil
}

// メールアドレス認証メールを送信する
func sendVerifyMail(mailAddress string, mailToken string) error {
	template := VerifyMailTemplate{
		VerifyURL: fmt.Sprintf("https://%s/create?m=%s", config.Defs.SiteDomain, mailToken),
		Mail:      mailAddress,
	}

	mailClient, err := mail.NewMail(
		"", mailAddress, "メールアドレス認証のお知らせ").AddContentsFromTemplate(VERIFY_MAIL_TEMPLATE_PATH, template)
	if err != nil {
		return err
	}

	return mailClient.Send()
}
