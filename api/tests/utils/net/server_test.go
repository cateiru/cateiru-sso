package net_test

import (
	"net/http"

	"github.com/cateiru/cateiru-sso/api/logging"
	"github.com/cateiru/cateiru-sso/api/utils/net"
)

type TestApp struct {
	*http.ServeMux
}

type PostForm struct {
	Name string `json:"name"`
}

// cookieテスト用のサーバ
func NewTestApp() *TestApp {
	mux := http.NewServeMux()
	app := &TestApp{mux}

	mux.HandleFunc("/set", app.TestSetCookiehandler)
	mux.HandleFunc("/get", app.TestGetCookieHandler)
	mux.HandleFunc("/delete", app.TestDeleteCookieHandler)

	mux.HandleFunc("/form", app.TestPostFormHandler)

	return app
}

// cookieを追加する
func (c *TestApp) TestSetCookiehandler(w http.ResponseWriter, r *http.Request) {
	exp := net.NewCookieMinutsExp(10)
	cookie := net.NewCookie("", false, http.SameSiteDefaultMode)
	cookie.Set(w, cookieKey, cookieValue, exp)

	w.Write([]byte("OK"))
}

// cookieを取得する
func (c *TestApp) TestGetCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := net.NewCookie("", true, http.SameSiteNoneMode)
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
	cookie := net.NewCookie("", true, http.SameSiteNoneMode)
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
