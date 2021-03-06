package handler

import (
	"net/http"

	"github.com/cateiru/cateiru-sso/api/core/user"
	"github.com/cateiru/cateiru-sso/api/core/user/history"
	"github.com/cateiru/cateiru-sso/api/core/user/info"
	"github.com/cateiru/cateiru-sso/api/core/user/mail"
	"github.com/cateiru/cateiru-sso/api/core/user/otp"
	"github.com/cateiru/cateiru-sso/api/core/user/password"
	"github.com/cateiru/cateiru-sso/api/utils/net"
)

// 自分のメールアドレスの操作
func UserMailHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userMailGetHandler(w, r)
	case http.MethodPost:
		userMailPostHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// パスワード変更
func UserPasswordHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		userPasswordPostHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// ワンタイムパスワードの無効化、有効化
func UserOnetimePWHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		userOnetimePWPostHandler(w, r)
	case http.MethodGet:
		createOnetimePWGetHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// ワンタイムパスワードのバックアップコードを表示する
func UserOnetimePWBackupHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		userOnetimePWBackupPostHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// ログインしているSSOの操作
func UserAccessHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userAccessGetHandler(w, r)
	case http.MethodPost:
		userAccessPostHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// アカウントのログイン履歴
func UserHistoryHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		userHistoryGetHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// アバターの設定
func UserAvatarHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		AvatarPostHandler(w, r)
	case http.MethodDelete:
		AvatarDeleteHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

func UserInfoChangeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		ChangeUserInfoHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

func UserOTPMeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		OTPMeHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

func UserOAuthHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		oauthGetHandler(w, r)
	case http.MethodDelete:
		oauthDeleteHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// メールアドレス取得
// 全ユーザ
// `/me`でも取得できる
func userMailGetHandler(w http.ResponseWriter, r *http.Request) {
	if err := mail.GetMailHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

// メールアドレスの更新
func userMailPostHandler(w http.ResponseWriter, r *http.Request) {
	if err := mail.CangeMailHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

// パスワード更新
// 現在のパスワードを送信するか、パスワード忘れ用の再登録トークンを送信
func userPasswordPostHandler(w http.ResponseWriter, r *http.Request) {
	if err := password.PasswordChangeHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

// ワンタイムパスワードのトークンURLを取得する
func createOnetimePWGetHandler(w http.ResponseWriter, r *http.Request) {
	if err := otp.GetOTPTokenURL(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

// ワンタイムパスワードの無効化、有効化
func userOnetimePWPostHandler(w http.ResponseWriter, r *http.Request) {
	if err := otp.OTPHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

// ワンタイムパスワードのバックアップコードを返す
func userOnetimePWBackupPostHandler(w http.ResponseWriter, r *http.Request) {
	if err := otp.BackupHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

// ログインしているSSOを取得
func userAccessGetHandler(w http.ResponseWriter, r *http.Request) {
}

// ログインしているSSOからログアウト
func userAccessPostHandler(w http.ResponseWriter, r *http.Request) {
}

// アカウントのログイン履歴取得
func userHistoryGetHandler(w http.ResponseWriter, r *http.Request) {
	if err := history.UserLoginHistoryHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

// アバター設定
func AvatarPostHandler(w http.ResponseWriter, r *http.Request) {
	if err := user.AvatarSetHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

// アバター削除
func AvatarDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if err := user.DeleteAvatarHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

func ChangeUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	if err := info.ChangeInfoHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

func OTPMeHandler(w http.ResponseWriter, r *http.Request) {
	if err := otp.OTPMeHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

func oauthGetHandler(w http.ResponseWriter, r *http.Request) {
	if err := user.OAuthShow(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

func oauthDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if err := user.DeleteOAth(w, r); err != nil {
		net.ResponseError(w, err)
	}
}
