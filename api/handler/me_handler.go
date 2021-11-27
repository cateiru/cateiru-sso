package handler

import "net/http"

// ユーザ情報を取得する
func MeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		meGetHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

// ユーザ情報を取得する
// cookieを見る
func meGetHandler(w http.ResponseWriter, r *http.Request) {
}
