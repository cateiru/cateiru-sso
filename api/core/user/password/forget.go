package password

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/logging"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	mail_ctrl "github.com/cateiru/cateiru-sso/api/utils/mail"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/cateiru-sso/api/utils/secure"
	"github.com/cateiru/go-http-error/httperror/status"
)

const CHANGE_PW_MAIL_TEMPLATE_PATH = "pw_change_mail"

type ForgetRequest struct {
	Mail string
}

type AccpetFortgetRequest struct {
	ForgetToken string
	NewPassword string
}

// テンプレートに適用する用の型
type ChangePWMailTemplate struct {
	VerifyURL string
	Mail      string
}

func ForgetPasswordRequestHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	db, err := database.NewDatabase(ctx)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	defer db.Close()

	var form ForgetRequest
	if err := net.GetJsonForm(w, r, &form); err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	return createChangeMail(ctx, db, form.Mail)
}

func ForgetPasswordAcceptHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	db, err := database.NewDatabase(ctx)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	defer db.Close()

	var form AccpetFortgetRequest
	if err := net.GetJsonForm(w, r, &form); err != nil {
		return status.NewBadRequestError(err).Caller()
	}

	return changePWAccept(ctx, db, &form)
}

func changePWAccept(ctx context.Context, db *database.Database, form *AccpetFortgetRequest) error {
	buffer, err := models.GetPWForgetByToken(ctx, db, form.ForgetToken)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	// entityが無い = tokenが無効な場合は400を返す
	if buffer == nil {
		return status.NewBadRequestError(errors.New("pw forget is no entity")).Caller()
	}
	// 有効期限の場合は400を返す
	if common.CheckExpired(&buffer.Period) {
		return status.NewBadRequestError(errors.New("expired")).Caller()
	}

	cert, err := models.GetCertificationByMail(ctx, db, buffer.Mail)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	pw := secure.PWHash(form.NewPassword)

	cert.Password = pw.Key
	cert.Salt = pw.Salt

	if err := cert.Add(ctx, db); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	return nil
}

// パスワードリセットメールを送信する
func createChangeMail(ctx context.Context, db *database.Database, mail string) error {
	existMail, err := common.CheckExistMail(ctx, db, mail)
	if err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}
	// メールアドレスが登録されていない場合は400を返す
	if !existMail {
		return status.NewBadRequestError(errors.New("mail is not exist")).Caller()
	}

	forgetToken := utils.CreateID(20)

	forgetBuffer := models.PWForget{
		ForgetToken: forgetToken,
		Mail:        mail,

		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 30,
		},
	}

	if err := forgetBuffer.Add(ctx, db); err != nil {
		return status.NewInternalServerErrorError(err).Caller()
	}

	if utils.DEPLOY_MODE == "production" {
		if err := sendPwMail(mail, forgetToken); err != nil {
			return err
		}
	} else {
		logging.Sugar.Debugf(
			"create pw_forget token. url: https://%s/pw/change?m=%s", os.Getenv("SITE_DOMAIN"), forgetToken)
	}

	return nil
}

func sendPwMail(mail string, forgetToken string) error {
	template := ChangePWMailTemplate{
		VerifyURL: fmt.Sprintf("https://%s/pw/change?m=%s", os.Getenv("SITE_DOMAIN"), forgetToken),
		Mail:      mail,
	}

	mailClient, err := mail_ctrl.NewMail(
		"", mail, "パスワード変更のお知らせ").AddContentsFromTemplate(CHANGE_PW_MAIL_TEMPLATE_PATH, template)
	if err != nil {
		return err
	}

	return mailClient.Send()
}
