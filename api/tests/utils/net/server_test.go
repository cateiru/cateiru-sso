package net_test

import (
	"errors"
	"net/http"

	"github.com/cateiru/cateiru-sso/api/logging"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/cateiru/go-http-error/httperror/status"
)

type TestApp struct {
	*http.ServeMux
}

type PostForm struct {
	Name string `json:"name"`
}

type ResponseOK struct {
	Status string `json:"status"`
}

// cookieテスト用のサーバ
func NewTestApp() *TestApp {
	mux := http.NewServeMux()
	app := &TestApp{mux}

	mux.HandleFunc("/set", app.TestSetCookiehandler)
	mux.HandleFunc("/get", app.TestGetCookieHandler)
	mux.HandleFunc("/delete", app.TestDeleteCookieHandler)

	mux.HandleFunc("/form", app.TestPostFormHandler)

	mux.HandleFunc("/ok", app.TestResponseOKHandler)
	mux.HandleFunc("/error", app.TestResponseErrorHandler)
	mux.HandleFunc("/notfound", app.TestNotfoundErrorHandler)

	mux.HandleFunc("/query", app.TestURLQuery)

	return app
}

// cookieを追加する
func (c *TestApp) TestSetCookiehandler(w http.ResponseWriter, r *http.Request) {
	exp := net.NewCookieMinutsExp(10)
	cookie := net.NewCookie("", false, http.SameSiteDefaultMode, true)
	cookie.Set(w, cookieKey, cookieValue, exp)

	w.Write([]byte("OK"))
}

// cookieを取得する
func (c *TestApp) TestGetCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := net.NewCookie("", true, http.SameSiteNoneMode, true)
	value, err := cookie.Get(r, cookieKey)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if value != cookieValue {
		w.Write([]byte("setされたcookieが違う"))
		return
	}

	w.Write([]byte("OK"))
}

// cookieを削除
func (c *TestApp) TestDeleteCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := net.NewCookie("", true, http.SameSiteNoneMode, true)
	err := cookie.Delete(w, r, cookieKey)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("OK"))
}

// postのjson formを取得し、name要素を返す
func (c *TestApp) TestPostFormHandler(w http.ResponseWriter, r *http.Request) {
	if !net.CheckContentType(r) {
		w.WriteHeader(http.StatusBadRequest)
		logging.Sugar.Error("content-type is not application/json")
		return
	}
	var postForm PostForm

	if err := net.GetJsonForm(w, r, &postForm); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logging.Sugar.Error(err.Error())
		return
	}

	w.Write([]byte(postForm.Name))
}

func (c *TestApp) TestResponseOKHandler(w http.ResponseWriter, r *http.Request) {
	body := ResponseOK{
		Status: "OK",
	}

	net.ResponseOK(w, body)
}

func (c *TestApp) TestResponseErrorHandler(w http.ResponseWriter, r *http.Request) {
	err := errors.New("Dummy error")

	net.ResponseError(w, err)
}

func (c *TestApp) TestNotfoundErrorHandler(w http.ResponseWriter, r *http.Request) {
	err := status.NewNotFoundError(errors.New("Dummy error")).Wrap()

	net.ResponseError(w, err)
}

func (c *TestApp) TestURLQuery(w http.ResponseWriter, r *http.Request) {
	query, err := net.GetQuery(r, "sample")
	if err != nil {
		net.ResponseError(w, err)
		return
	}

	w.Write([]byte(query))
}
