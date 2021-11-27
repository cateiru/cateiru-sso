package handler

import "net/http"

// SSOの管理
func ProSSOHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		proSSOGetHandler(w, r)
	case http.MethodPost:
		proSSOPostHandler(w, r)
	case http.MethodDelete:
		proSSODeleteHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// 自分のSSO情報取得
// pro以上のユーザのみ
func proSSOGetHandler(w http.ResponseWriter, r *http.Request) {
}

// SSO追加
// pro以上のユーザのみ
func proSSOPostHandler(w http.ResponseWriter, r *http.Request) {
}

// SSO削除
// pro以上のユーザのみ
func proSSODeleteHandler(w http.ResponseWriter, r *http.Request) {
}
