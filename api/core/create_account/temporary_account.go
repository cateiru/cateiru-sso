package createaccount

import (
	"context"
	"errors"
	"net/http"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

type PostForm struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type Response struct {
	ClientCheckToken string `json:"client_check_token"`
}

// 一時的にアカウントを作成します
// メールアドレス、パスワードをfromで送信することで、そのメールアドレスに確認用URLを送信します。
// さらに、Websocketでメールアドレスが認証されたか確認するためのトークンを返します。
//
// Post Form (application/json):
//	{
//		"mail": "example@example.com",
//		"password": "**********",
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

	clientCheckToken, err := createTemporaryAccount(ctx, postForm)
	if err != nil {
		return err
	}

	response := Response{
		ClientCheckToken: clientCheckToken,
	}

	net.ResponseOK(w, response)

	return nil
}

func createTemporaryAccount(ctx context.Context, form *PostForm) (string, error) {
	db, err := database.NewDatabase(ctx)
	if err != nil {
		return "", status.NewInternalServerErrorError(err).Caller(
			"core/create_account/temporary_account.go", 35).Wrap()
	}

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
