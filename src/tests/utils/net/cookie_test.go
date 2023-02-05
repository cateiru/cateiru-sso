package net_test

import (
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/cateiru/cateiru-sso/src/tests/tools"
	"github.com/cateiru/cateiru-sso/src/utils/net"
	"github.com/stretchr/testify/require"
)

var cookieKey string = "KEY"
var cookieValue string = "hoge"

func cookieTestServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/set", testSetCookiehandler)
	mux.HandleFunc("/get", testGetCookieHandler)
	mux.HandleFunc("/delete", testDeleteCookieHandler)

	return mux
}

// cookieを追加する
func testSetCookiehandler(w http.ResponseWriter, r *http.Request) {
	exp := net.NewCookieMinutsExp(10)
	cookie := net.NewCookie("", false, http.SameSiteDefaultMode, true)
	cookie.Set(w, cookieKey, cookieValue, exp)

	w.Write([]byte("OK"))
}

// cookieを取得する
func testGetCookieHandler(w http.ResponseWriter, r *http.Request) {
	value, err := net.GetCookie(r, cookieKey)
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
func testDeleteCookieHandler(w http.ResponseWriter, r *http.Request) {
	err := net.DeleteCookie(w, r, cookieKey)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("OK"))
}

// レスポンスを調べる
func respMessage(t *testing.T, res *http.Response, message string) {
	require.Equal(t, tools.ConvertResp(res), "OK", message)
}

// Cookieのテスト
func TestCookie(t *testing.T) {
	app := cookieTestServer()
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
