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
	ClientToken string `json:"client_token"`

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
		return status.NewBadRequestError(errors.New("core/create_account/info.go")).Caller()
	}

	var userData InfoRequestForm
	err := net.GetJsonForm(w, r, &userData)
	if err != nil {
		return status.NewBadRequestError(errors.New("parse not failed")).Caller()
	}

	ip := net.GetIPAddress(r)
	userAgent := net.GetUserAgent(r)

	ctx := r.Context()

	login, err := InsertUserInfo(ctx, userData.ClientToken, userData, ip, userAgent)
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
func InsertUserInfo(ctx context.Context, clientToken string, user InfoRequestForm, ip string, userAgent string) (*common.LoginTokens, error) {
	db, err := database.NewDatabase(ctx)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}
	defer db.Close()

	buffer, err := models.GetMailCertificationByClientToken(ctx, db, clientToken)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	// bufferのentryがなかった場合、400を返す
	if buffer == nil {
		return nil, status.NewBadRequestError(errors.New("buffer is not exist")).Caller()
	}

	// 有効期限が切れている場合は、400を返す
	if common.CheckExpired(&buffer.Period) {
		return nil, status.NewBadRequestError(errors.New("expired")).Caller().AddCode(net.TimeOutError)
	}

	// メールアドレスが未認証の場合は400を返す
	if !buffer.Verify {
		return nil, status.NewBadRequestError(errors.New("email address is unauthenticated")).Caller()
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
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	// ユーザ情報追加
	userInfo := &models.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Theme:     user.Theme,
		AvatarUrl: user.AvatarUrl,

		Mail: buffer.Mail,

		UserId: models.UserId{
			UserId: userId,
		},
	}

	if err = userInfo.Add(ctx, db); err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	// ユーザの権限
	role := &models.Role{
		// デフォルトは`user`のみ
		Role: []string{"user"},

		UserId: models.UserId{
			UserId: userId,
		},
	}

	if err := role.Add(ctx, db); err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	if err := models.DeleteMailCertification(ctx, db, buffer.MailToken); err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	return common.LoginByUserID(ctx, db, userId, ip, userAgent)
}
