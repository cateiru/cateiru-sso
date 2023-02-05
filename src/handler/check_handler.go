package handler

import (
	"net/http"

	"github.com/cateiru/cateiru-sso/api/core/check"
	"github.com/cateiru/cateiru-sso/api/utils/net"
)

func CheckHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		checkUserNameHandler(w, r)
	default:
		RootHandler(w, r)
	}
}

func checkUserNameHandler(w http.ResponseWriter, r *http.Request) {
	if err := check.CheckUserNameHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}
