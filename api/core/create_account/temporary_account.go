package createaccount

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
	"github.com/cateiru/cateiru-sso/api/utils/mail"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/cateiru-sso/api/utils/secure"
	"github.com/cateiru/go-http-error/httperror/status"
)

const VERIFY_MAIL_TEMPLATE_PATH = "verify_mail"

// POSTのformの型
type PostForm struct {
	Mail       string `json:"mail"`
	Password   string `json:"password"`
	ReCHAPTCHA string `json:"re_chaptcha"`
}

// レスポンスの型
type Response struct {
	ClientCheckToken string `json:"client_check_token"`
}

// テンプレートに適用する用の型
type VerifyMailTemplate struct {
	VerifyURL string
	Mail      string
}

// 一時的にアカウントを作成します
// メールアドレス、パスワードをfromで送信することで、そのメールアドレスに確認用URLを送信します。
// さらに、Websocketでメールアドレスが認証されたか確認するためのトークンを返します。
//
// Post Form (application/json):
//	{
//		"mail": "example@example.com",
//		"password": "**********",
//		"re_chaptcha": "********",
//	}
func CreateTemporaryHandler(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()

	// contents-type: application/json 以外では403エラーを返す
	if net.CheckContentType(r) {
		return status.NewForbiddenError(errors.New("requests contets-type is not application/json")).Caller("core/create_account/temporary_account.go", 24)
	}

	postForm := new(PostForm)
	if err := net.GetJsonForm(w, r, postForm); err != nil {
		return err
	}

	ip := r.Header.Get("x-forwarded-for")

	// reCHAPTCHA
	if utils.DEPLOY_MODE == "production" {
		isOk, err := secure.NewReCaptcha().Validate(postForm.ReCHAPTCHA, ip)
		if err != nil {
			return status.NewInternalServerErrorError(err).Caller(
				"core/create_account/temporary_account.go", 68).Wrap()
		}
		// reCHAPTCHAが認証できなかった場合、403を返す
		if !isOk {
			return status.NewForbiddenError(errors.New("reCHAPTCHA is failed")).Caller("core/create_account/temporary_account.go", 71)
		}
	}

	clientCheckToken, err := CreateTemporaryAccount(ctx, postForm, ip)
	if err != nil {
		return err
	}

	response := Response{
		ClientCheckToken: clientCheckToken,
	}

	net.ResponseOK(w, response)

	return nil
}

func CreateTemporaryAccount(ctx context.Context, form *PostForm, ip string) (string, error) {
	db, err := database.NewDatabase(ctx)
	if err != nil {
		return "", status.NewInternalServerErrorError(err).Caller(
			"core/create_account/temporary_account.go", 35).Wrap()
	}
	defer db.Close()

	// IPアドレス、メールアドレスがブロックされているか確認
	isBlocked, err := common.ChaeckBlock(ctx, db, ip, form.Mail)
	if err != nil {
		return "", status.NewInternalServerErrorError(err).Caller(
			"core/create_account/temporary_account.go", 35).Wrap()
	}
	if isBlocked {
		return "", status.NewForbiddenError(errors.New("ip is blocked")).Caller("core/create_account/temporary_account.go", 47).Wrap()
	}

	// メールアドレスが既に存在するかチェック
	isMailExist, err := common.CheckExistMail(ctx, db, form.Mail)
	if err != nil {
		return "", status.NewInternalServerErrorError(err).Caller(
			"core/create_account/temporary_account.go", 41).Wrap()
	}
	// メールアドレスがすでに存在している = そのメールアドレスを持ったアカウントが作られている場合、
	// あたらにそのメールアドレスでアカウントを作成することはできないため、403エラーを返す
	if isMailExist {
		return "", status.NewForbiddenError(errors.New("email already exists")).Caller("core/create_account/temporary_account.go", 47).Wrap()
	}
	// Adminのメールアドレスは既に定義されており、ログインできるため弾く
	if common.CheckAdminMail(form.Mail) {
		return "", status.NewForbiddenError(errors.New("email is admin")).Caller("core/create_account/temporary_account.go", 47).Wrap()
	}

	user := models.UserMailPW{
		Mail:     form.Mail,
		Password: utils.PWHash(form.Password),
	}
	clientCheckToken, err := createVerifyMail(ctx, db, user)
	if err != nil {
		return "", status.NewInternalServerErrorError(err).Caller(
			"core/create_account/temporary_account.go", 55).Wrap()
	}

	return clientCheckToken, nil
}

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

	if err := mailVerify.Add(ctx, db); err != nil {
		return "", err
	}

	// send mail
	// SendGrid APIをテストでは使用しないため、
	// DEPLOY_MODEがproductionのときのみ送信します
	if utils.DEPLOY_MODE == "production" {
		if err := sendVerifyMail(user.Mail, mailToken); err != nil {
			return "", err
		}
	} else {
		logging.Sugar.Debugf("create mail token. url: https://%s/create?m=%s", os.Getenv("SITE_DOMAIN"), mailToken)
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
