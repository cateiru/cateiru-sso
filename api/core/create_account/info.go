package createaccount

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

type InfoRequestForm struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`

	Theme     string `json:"theme"`
	AvatarUrl string `json:"avatar_url"`
}

// ユーザ情報を設定し、ログイン状態にします
func CreateInfoHandler(w http.ResponseWriter, r *http.Request) error {
	// contents-type: application/json 以外では400エラーを返す
	if !net.CheckContentType(r) {
		return status.NewBadRequestError(errors.New("core/create_account/info.go")).Caller(
			"core/create_account/temporary_account.go", 56)
	}

	var userData InfoRequestForm
	err := net.GetJsonForm(w, r, &userData)
	if err != nil {
		return status.NewBadRequestError(errors.New("parse not failed")).Caller(
			"core/create_account/info.go", 62)
	}

	bufferToken, err := net.GetCookie(r, "buffer-token")
	if err != nil || len(bufferToken) == 0 {
		// cookieが存在しない、valueが存在しない場合は403を返す
		return status.NewForbiddenError(errors.New("cookie is not found")).Caller(
			"core/create_account/info.go", 36)
	}

	ctx := r.Context()

	login, err := InsertUserInfo(ctx, bufferToken, userData)
	if err != nil {
		return err
	}

	// ログイン用のトークンをcookieにセットする
	common.LoginSetCookie(w, login)

	return nil
}

// ユーザ情報を入力し、アカウントを正式に登録します
//
// 登録後、userIdを返します
func InsertUserInfo(ctx context.Context, bufferToken string, user InfoRequestForm) (*common.LoginTokens, error) {
	db, err := database.NewDatabase(ctx)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller(
			"core/create_account/info.go", 100).Wrap()
	}
	defer db.Close()

	buffer, err := models.GetCreateAccountBufferByBufferToken(ctx, db, bufferToken)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller(
			"core/create_account/info.go", 100).Wrap()
	}

	// bufferのentryがなかった場合、400を返す
	if buffer == nil {
		return nil, status.NewBadRequestError(errors.New("buffer is not exist")).Caller(
			"core/create_account/info.go", 69)
	}

	// 有効期限が切れている場合は、400を返す
	if common.CheckExpired(&buffer.VerifyPeriod) {
		return nil, status.NewBadRequestError(errors.New("expired")).Caller(
			"core/create_account/verify.go", 67).AddCode(net.TimeOutError).Wrap()
	}

	userId := utils.CreateID(30)

	// ユーザ認証情報追加
	certification := &models.Certification{
		AccountCreateDate: time.Now(),

		// アカウント作成後はOTPは設定しない
		// 設定ページから追加する
		OnetimePasswordSecret:  "",
		OnetimePasswordBackups: []string{},

		UserMailPW: buffer.UserMailPW,
		UserId: models.UserId{
			UserId: userId,
		},
	}
	if err = certification.Add(ctx, db); err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller(
			"core/create_account/info.go", 100).Wrap()
	}

	// ユーザ情報追加
	userInfo := &models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Theme:     user.Theme,
		AvatarUrl: user.AvatarUrl,

		// デフォルトは`user`のみ
		Role: []string{"user"},

		Mail: buffer.Mail,

		UserId: models.UserId{
			UserId: userId,
		},
	}

	if err = userInfo.Add(ctx, db); err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller(
			"core/create_account/info.go", 100).Wrap()
	}

	// CreateAccontBufferは削除する
	if err := models.DeleteCreateAccountBuffer(ctx, db, buffer.BufferToken); err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller(
			"core/create_account/info.go", 107).Wrap()
	}

	return common.LoginByUserID(ctx, db, userId)
}
