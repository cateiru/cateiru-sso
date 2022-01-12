package password

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
	mail_ctrl "github.com/cateiru/cateiru-sso/api/utils/mail"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

const CHANGE_PW_MAIL_TEMPLATE_PATH = "pw_change_mail"

type ForgetRequest struct {
	Mail string `json:"mail"`
}

// テンプレートに適用する用の型
type ChangePWMailTemplate struct {
	VerifyURL string
	Mail      string
}

func ForgetPasswordRequestHandler(w http.ResponseWriter, r *http.Request) error {
	// contents-type: application/json 以外では400エラーを返す
	if !net.CheckContentType(r) {
		return status.NewBadRequestError(errors.New("requests contets-type is not application/json")).Caller()
	}

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

	return CreateChangeMail(ctx, db, form.Mail)
}

// パスワードリセットメールを送信する
func CreateChangeMail(ctx context.Context, db *database.Database, mail string) error {
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

	if config.Defs.DeployMode == "production" {
		if err := sendPwMail(mail, forgetToken); err != nil {
			return err
		}
	} else {
		logging.Sugar.Debugf(
			"create pw_forget token. url: https://%s/pw/change?m=%s", config.Defs.SiteDomain, forgetToken)
	}

	return nil
}

func sendPwMail(mail string, forgetToken string) error {
	template := ChangePWMailTemplate{
		VerifyURL: fmt.Sprintf("https://%s/pw/change?m=%s", config.Defs.SiteDomain, forgetToken),
		Mail:      mail,
	}

	mailClient, err := mail_ctrl.NewMail(
		"", mail, "パスワード変更のお知らせ").AddContentsFromTemplate(CHANGE_PW_MAIL_TEMPLATE_PATH, template)
	if err != nil {
		return err
	}

	return mailClient.Send()
}
