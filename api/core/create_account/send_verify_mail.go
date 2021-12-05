package createaccount

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/logging"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/cateiru/cateiru-sso/api/utils/mail"
)

type VerifyMailTemplate struct {
	VerifyURL string
	Mail      string
}

const VERIFY_MAIL_TEMPLATE_PATH = "verify_mail"

// メール認証を開始します
//
// client_check_token(wsを接続するのに使用するトークンを返します)
func createVerifyMail(ctx context.Context, db *database.Database, user models.UserMailPW) (string, error) {
	mailToken := utils.CreateID(0)
	clientCheckToken := utils.CreateID(0)

	mailVerify := &models.MailCertification{
		MailToken:        mailToken,
		ClientCheckToken: clientCheckToken,
		CreateDate:       time.Now(),
		PeriodMinute:     30,

		OpenNewWindow:  false,
		Verify:         false,
		ChangeMailMode: false,

		UserMailPW: user,
	}

	key := database.CreateNameKey("MailCertification", mailToken)

	if err := db.Put(ctx, key, mailVerify); err != nil {
		return "", err
	}

	// send mail
	// SendGrid APIをテストでは使用しないため、
	// DEPLOY_MODEがproductionのときのみ送信します
	if utils.DEPLOY_MODE == "production" {
		if err := sendVerifyMail(user.Mail, mailToken); err != nil {
			return "", err
		}
	}

	logging.Sugar.Debugf("Send verify email. mail: %s, client check token: %s", user.Mail, clientCheckToken)

	return clientCheckToken, nil
}

// メールアドレス認証メールを送信する
func sendVerifyMail(mailAddress string, mailToken string) error {
	template := VerifyMailTemplate{
		VerifyURL: fmt.Sprintf("https://%s/create?m=%s", os.Getenv("SITE_DOMAIN"), mailToken),
		Mail:      mailAddress,
	}

	mailClient, err := mail.NewMail("", mailAddress, "メールアドレス認証のお知らせ").AddContentsFromTemplate(VERIFY_MAIL_TEMPLATE_PATH, template)
	if err != nil {
		return err
	}

	return mailClient.Send()
}
