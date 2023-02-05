package handler

import (
	"net/http"

	"github.com/cateiru/cateiru-sso/src/core/password"
	"github.com/cateiru/cateiru-sso/src/utils/net"
)

// パスワード忘れの再登録用
func PasswordForgetHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		passwordForgetPostHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

func PasswordForgetAcceptHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		passwordForgetAcceptPostHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// パスワードを忘れた場合の再登録
//
// メールアドレスを送信して、そのメールアドレスの持ったアカウントが存在する場合に、
// トークンをパラメータに付与したURLをメール送信
func passwordForgetPostHandler(w http.ResponseWriter, r *http.Request) {
	if err := password.ForgetPasswordRequestHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

// パスワードの再登録
func passwordForgetAcceptPostHandler(w http.ResponseWriter, r *http.Request) {
	if err := password.ForgetPasswordAcceptHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}
