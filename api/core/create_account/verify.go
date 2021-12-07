package createaccount

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

type VerifyRequestForm struct {
	MailToken string `json:"mail_token"`
}

type VerifyStatus struct {
	BufferToken string
	IsSetCookie bool
}

type VerifyResponse struct {
	IsKeepThisPage bool `json:"keep_this_page"`
}

// mail tokenを受け取り、該当するメールアドレスを認証済みにします。
// さらに、CreateAccountBufferにアップデートし、openNewWindowがtrueの場合は、BufferTokenをcookieに入れます。
//
// Request Form (application/json):
//	{
//		"mail_token": "*******",
//	}
func CreateVerifyHandler(w http.ResponseWriter, r *http.Request) error {
	// contents-type: application/json 以外では400エラーを返す
	if net.CheckContentType(r) {
		return status.NewBadRequestError(errors.New("requests contets-type is not application/json")).Caller(
			"core/create_account/verify.go", 26)
	}

	postForm := new(VerifyRequestForm)
	err := net.GetJsonForm(w, r, postForm)
	if err != nil {
		return status.NewBadRequestError(errors.New("parse not failed")).Caller(
			"core/create_account/verify.go", 33)
	}

	ctx := r.Context()

	verify, err := CreateVerify(ctx, postForm.MailToken)
	if err != nil {
		return err
	}

	// 開いたベージでそのまま続ける場合はcookieを設定する
	if verify.IsSetCookie {
		// secure属性はproductionのみにする（テストが通らないため）
		secure := false
		if utils.DEPLOY_MODE == "production" {
			secure = true
		}
		cookie := net.NewCookie(os.Getenv("COOKIE_DOMAIN"), secure, http.SameSiteDefaultMode)

		// 有効期限1時間
		cookieExp := net.NewCookieHourExp(1)
		cookie.Set(w, "buffer_token", verify.BufferToken, cookieExp)
	}

	res := VerifyResponse{
		// cookieをセットする = そのページで続ける
		IsKeepThisPage: verify.IsSetCookie,
	}
	net.ResponseOK(w, res)

	return nil
}

func CreateVerify(ctx context.Context, mailToken string) (*VerifyStatus, error) {
	db, err := database.NewDatabase(ctx)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller(
			"core/create_account/verify.go", 100).Wrap()
	}
	defer db.Close()

	certificationEntry, err := models.GetMailCertificationByMailToken(ctx, db, mailToken)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller(
			"core/create_account/verify.go", 100).Wrap()
	}

	// 既に削除されている場合、400を返す
	if certificationEntry == nil {
		return nil, status.NewBadRequestError(err).Caller("core/create_account/verify.go", 98).Wrap()
	}

	// 既に認証済みの場合、400を返す
	if certificationEntry.Verify {
		return nil, status.NewBadRequestError(err).Caller(
			"core/create_account/verify.go", 103).AddCode(net.AlreadyDone).Wrap()
	}

	// 有効期限が切れている場合は、400を返す
	now := time.Now()
	if now.Sub(certificationEntry.CreateDate) >= time.Duration(certificationEntry.PeriodMinute)*time.Minute {
		return nil, status.NewBadRequestError(err).Caller(
			"core/create_account/verify.go", 67).AddCode(net.TimeOutError).Wrap()
	}

	var bufferToken string

	if certificationEntry.OpenNewWindow {
		// Websocketの監視は終わっているため、ユーザ情報をCreateAccountBufferに移行してこのentryは削除する
		bufferToken = utils.CreateID(20)
		buffer := &models.CreateAccountBuffer{
			BufferToken: bufferToken,
			VerifyPeriod: models.VerifyPeriod{
				CreateDate:   time.Now(),
				PeriodMinute: 60,
			},
			UserMailPW: certificationEntry.UserMailPW,
		}
		if err := buffer.Add(ctx, db); err != nil {
			return nil, status.NewInternalServerErrorError(err).Caller(
				"core/create_account/verify.go", 100).Wrap()
		}

		if err := models.DeleteMailCertification(ctx, db, mailToken); err != nil {
			return nil, status.NewInternalServerErrorError(err).Caller(
				"core/create_account/verify.go", 133).Wrap()
		}

	} else {
		// 認証: trueにする
		certificationEntry.Verify = true
		if err := certificationEntry.Add(ctx, db); err != nil {
			return nil, status.NewInternalServerErrorError(err).Caller(
				"core/create_account/verify.go", 100).Wrap()
		}
	}

	return &VerifyStatus{
		IsSetCookie: certificationEntry.OpenNewWindow,
		BufferToken: bufferToken,
	}, nil
}
