package createaccount

import (
	"context"
	"errors"
	"net/http"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

type VerifyRequestForm struct {
	MailToken string `json:"mail_token"`
}

type VerifyResponse struct {
	IsKeepThisPage bool   `json:"keep_this_page"`
	ClientToken    string `json:"client_token"`
}

// mail tokenを受け取り、該当するメールアドレスを認証済みにします。
//
// Request Form (application/json):
//	{
//		"mail_token": "*******",
//	}
func CreateVerifyHandler(w http.ResponseWriter, r *http.Request) error {
	// contents-type: application/json 以外では400エラーを返す
	if !net.CheckContentType(r) {
		return status.NewBadRequestError(errors.New("requests contets-type is not application/json")).Caller()
	}

	postForm := new(VerifyRequestForm)
	err := net.GetJsonForm(w, r, postForm)
	if err != nil {
		return status.NewBadRequestError(errors.New("parse not find")).Caller()
	}

	ctx := r.Context()

	verify, err := CreateVerify(ctx, postForm.MailToken)
	if err != nil {
		return err
	}

	net.ResponseOK(w, verify)

	return nil
}

func CreateVerify(ctx context.Context, mailToken string) (*VerifyResponse, error) {
	db, err := database.NewDatabase(ctx)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}
	defer db.Close()

	certificationEntry, err := models.GetMailCertificationByMailToken(ctx, db, mailToken)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	// 既に削除されている場合、400を返す
	if certificationEntry == nil {
		return nil, status.NewBadRequestError(errors.New("deleted entry")).Caller()
	}

	// 既に認証済みの場合、400を返す
	if certificationEntry.Verify {
		return nil, status.NewBadRequestError(errors.New("verified")).Caller().AddCode(net.AlreadyDone)
	}

	// 有効期限が切れている場合は、400を返す
	if common.CheckExpired(&certificationEntry.Period) {
		return nil, status.NewBadRequestError(errors.New("Expired")).Caller().AddCode(net.TimeOutError)
	}

	// 認証: trueにする
	certificationEntry.Verify = true
	if err := certificationEntry.Add(ctx, db); err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	return &VerifyResponse{
		IsKeepThisPage: certificationEntry.OpenNewWindow,
		ClientToken:    certificationEntry.ClientToken,
	}, nil
}
