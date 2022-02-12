package handler

import (
	"net/http"
	"net/url"

	"github.com/cateiru/cateiru-sso/api/core/oauth"
	"github.com/cateiru/cateiru-sso/api/utils/net"
)

func OAuthPreview(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		oauthPerviewPost(w, r)
	default:
		RootHandler(w, r)
	}
}

func OAuthLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		oauthLoginPost(w, r)
	default:
		RootHandler(w, r)
	}
}

func OAuthToken(w http.ResponseWriter, r *http.Request) {
	var query url.Values

	switch r.Method {
	case http.MethodGet:
		query = r.URL.Query()
	case http.MethodPost:
		query = r.PostForm
	default:
		RootHandler(w, r)
		return
	}

	if err := oauth.TokenEndpoint(w, r, query); err != nil {
		net.ResponseError(w, err)
	}
}

func OAuthJWTKey(w http.ResponseWriter, r *http.Request) {
	if err := oauth.JWTPublicHandler(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

func oauthPerviewPost(w http.ResponseWriter, r *http.Request) {
	if err := oauth.ServicePreview(w, r); err != nil {
		net.ResponseError(w, err)
	}
}

func oauthLoginPost(w http.ResponseWriter, r *http.Request) {
	if err := oauth.ServiceLogin(w, r); err != nil {
		net.ResponseError(w, err)
	}
}
