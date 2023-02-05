package handler

import (
	"net/http"

	"github.com/cateiru/cateiru-sso/api/core/login"
	"github.com/cateiru/cateiru-sso/api/utils/net"
)

// ログインする
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		loginPostHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

func LoginOnetimeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		loginOnetimePostHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// メアド、パスワードを送信してcookieを作成
// userがadminの場合で初回ログインの場合はワンタイムパスワードはいらない
func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	if err := login.LoginHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

// ワンタイムパスワードを入力（必要な場合）
func loginOnetimePostHandler(w http.ResponseWriter, r *http.Request) {
	if err := login.OTPLoginHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}
