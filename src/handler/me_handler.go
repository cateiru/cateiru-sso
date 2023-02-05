package handler

import (
	"net/http"

	"github.com/cateiru/cateiru-sso/src/core/me"
	"github.com/cateiru/cateiru-sso/src/utils/net"
)

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
	if err := me.MeHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}
