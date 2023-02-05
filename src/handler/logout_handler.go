package handler

import (
	"net/http"

	"github.com/cateiru/cateiru-sso/src/core/logout"
	"github.com/cateiru/cateiru-sso/src/utils/net"
)

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
	if err := logout.LogoutHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

// アカウント削除
func logoutDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if err := logout.DeleteHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}
