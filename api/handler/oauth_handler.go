package handler

import "net/http"

// ssoのセッショントークンを使用してユーザを認証し、情報を取得する
func OAuthCertHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		oauthCertPostHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// ssoのセッショントークンをリフレッシュトークンを使用して更新する
func OAuthUpdateHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		oauthUpdatePostHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// セッショントークンでユーザ情報取得
// login_onlyがfalseのみ返す
// /user/accessでユーザはこれを停止できる
func oauthCertPostHandler(w http.ResponseWriter, r *http.Request) {
}

// ssoのセッショントークンをリフレッシュトークンを使用して更新する
func oauthUpdatePostHandler(w http.ResponseWriter, r *http.Request) {
}
