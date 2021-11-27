package handler

import "net/http"

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		logoutGetHandler(w, r)
	case http.MethodDelete:
		logoutDeleteHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// ログアウト
func logoutGetHandler(w http.ResponseWriter, r *http.Request) {
}

// アカウント削除
func logoutDeleteHandler(w http.ResponseWriter, r *http.Request) {
}
