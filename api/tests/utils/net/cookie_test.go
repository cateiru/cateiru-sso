package net_test

import (
	"bytes"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/stretchr/testify/require"
)

var cookieKey string = "KEY"
var cookieValue string = "hoge"

type TestApp struct {
	*http.ServeMux
}

// cookieテスト用のサーバ
func NewTestApp() *TestApp {
	mux := http.NewServeMux()
	app := &TestApp{mux}

	mux.HandleFunc("/set", app.TestSetCookiehandler)
	mux.HandleFunc("/get", app.TestGetCookieHandler)
	mux.HandleFunc("/delete", app.TestDeleteCookieHandler)

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

// レスポンスを調べる
func respMessage(t *testing.T, res *http.Response, message string) {
	defer res.Body.Close()

	buf := &bytes.Buffer{}
	buf.ReadFrom(res.Body)

	require.Equal(t, buf.String(), "OK", message)
}

// Cookieのテスト
func TestCookie(t *testing.T) {
	app := NewTestApp()
	server := httptest.NewServer(app)
	defer server.Close()

	jar, err := cookiejar.New(nil)
	require.NoError(t, err, "cookiejarでエラー")
	client := &http.Client{Jar: jar}

	resp, err := client.Get(server.URL + "/set")
	require.NoError(t, err, "cookieをsetできない")
	respMessage(t, resp, "cookieをsetしたけどエラーが起きた")

	set_cookie_url, err := url.Parse(server.URL + "/set")
	require.NoError(t, err, "cookieを取得できない")
	cookies := jar.Cookies(set_cookie_url)
	require.Len(t, cookies, 1)
	require.Equal(t, cookies[0].Value, cookieValue, "setされたcookieが違う")

	resp, err = client.Get(server.URL + "/get")
	require.NoError(t, err, "cookieをgetできない")
	respMessage(t, resp, "cookieをgetしたけどエラーが起きた")

	resp, err = client.Get(server.URL + "/delete")
	require.NoError(t, err, "cookieを削除できない")
	respMessage(t, resp, "cookieをdeleteしたけどエラーが起きた")
}
