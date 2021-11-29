package handler

import "net/http"

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

// SSOでログインする
func LoginSSOHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		loginSSOPostHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// メアド、パスワード、ワンタイムパスワードを送信してcookieを作成
// userがadminの場合で初回ログインの場合はワンタイムパスワードはいらない
func loginPostHandler(w http.ResponseWriter, r *http.Request) {
}

// ワンタイムパスワードを入力（必要な場合）
func loginOnetimePostHandler(w http.ResponseWriter, r *http.Request) {
}

// sso_public_keyを送信してトークンを作成する
func loginSSOPostHandler(w http.ResponseWriter, r *http.Request) {
}
