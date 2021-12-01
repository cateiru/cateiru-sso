package net_test

import (
	"bytes"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

var cookieKey string = "KEY"
var cookieValue string = "hoge"

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
