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
	"github.com/cateiru/cateiru-sso/api/utils/secure"
	"github.com/cateiru/go-http-error/httperror/status"
)

type InfoRequestForm struct {
	ClientToken string `json:"client_token"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`

	Password string `json:"password"`

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

	ctx := r.Context()

	cert := common.NewCert(w, r).AddUser()

	userInfo, err := InsertUserInfo(ctx, userData, cert)
	if err != nil {
		return err
	}

	// ログイン用のトークンをcookieにセットする
	cert.SetCookie()

	net.ResponseOK(w, userInfo)

	return nil
}

// ユーザ情報を入力し、アカウントを正式に登録します
//
// 登録後、userIdを返します
func InsertUserInfo(ctx context.Context, user InfoRequestForm, cert *common.Cert) (*models.User, error) {
	db, err := database.NewDatabase(ctx)
	if err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}
	defer db.Close()

	buffer, err := models.GetMailCertificationByClientToken(ctx, db, user.ClientToken)
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

	// UserIDはユニークであるためすでに存在している場合は400を返す
	existUserName, err := common.CheckUsername(ctx, db, user.UserName)
	if err != nil {
		return nil, err
	}
	if existUserName {
		return nil, status.NewBadRequestError(errors.New("user id is already exists")).Caller().AddCode(net.ExistUserName)
	}

	userId := utils.CreateID(30)

	hashedPW, err := secure.PWHash(user.Password)
	if err != nil {
		return nil, status.NewBadRequestError(err).Caller()
	}

	// ユーザ認証情報追加
	certification := &models.Certification{
		AccountCreateDate: time.Now(),

		// アカウント作成後はOTPは設定しない
		// 設定ページから追加する
		OnetimePasswordSecret:  "",
		OnetimePasswordBackups: []string{},

		UserMailPW: models.UserMailPW{
			Mail:     buffer.Mail,
			Password: hashedPW.Key,
			Salt:     hashedPW.Salt,
		},

		UserId: models.UserId{
			UserId: userId,
		},
	}
	if err = certification.Add(ctx, db); err != nil {
		return nil, status.NewInternalServerErrorError(err).Caller()
	}

	if !utils.CheckUserName(user.UserName) {
		return nil, status.NewBadRequestError(errors.New("incorrect username")).Caller().AddCode(net.IncorrectUserName)
	}

	// ユーザ情報追加
	userInfo := &models.User{
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		UserName:          user.UserName,
		UserNameFormatted: utils.FormantUserName(user.UserName),
		Theme:             user.Theme,
		AvatarUrl:         user.AvatarUrl,

		Mail: buffer.Mail,

		// デフォルトは`user`のみ
		Role: []string{"user"},

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

	if err := cert.NewLogin(ctx, db, userId); err != nil {
		return nil, err
	}

	return userInfo, nil
}
