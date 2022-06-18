package tools

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/stretchr/testify/require"
)

type TestServer struct {
	Server       *httptest.Server
	Client       *http.Client
	Jar          *cookiejar.Jar
	Dummy        *DummyUser
	Url          *url.URL
	SessionToken string
	RefreshToken string
}

func NewTestServer(t *testing.T, handler http.Handler, isCookie bool) *TestServer {
	server := httptest.NewServer(handler)

	self := &TestServer{
		Server: server,
	}

	if isCookie {
		jar, err := cookiejar.New(nil)
		require.NoError(t, err)
		client := &http.Client{Jar: jar}

		self.Jar = jar
		self.Client = client

		// URLをセット
		url, err := url.Parse(server.URL + "/")
		require.NoError(t, err)
		self.Url = url
	} else {
		self.Client = &http.Client{}
	}

	return self
}

// cookieにセッション情報を追加する
func (c *TestServer) AddSession(ctx context.Context, db *database.Database, dummy *DummyUser) error {
	session, refresh, err := dummy.AddLoginToken(ctx, db, time.Now())
	if err != nil {
		return err
	}

	c.SessionToken = session
	c.RefreshToken = refresh

	exp := net.NewCookieMinutsExp(3)
	SetCookie(c.Jar, "session-token", session, exp, c.Url)
	SetCookie(c.Jar, "refresh-token", refresh, exp, c.Url)

	return nil
}

// cookieが存在するかをチェックする
func (c *TestServer) FindCookies(t *testing.T, cookies []string) {
	flag := make([]bool, len(cookies))

	for _, cookie := range c.Jar.Cookies(c.Url) {
		for index, target := range cookies {
			if cookie.Name == target && len(cookie.Value) != 0 {
				flag[index] = true
			}
		}
	}

	for index, isExist := range flag {
		require.True(t, isExist, fmt.Sprintf("cookie %v is exist.", cookies[index]))
	}
}

func (c *TestServer) GetCookie(key string) string {
	for _, cookie := range c.Jar.Cookies(c.Url) {
		if cookie.Name == key {
			return cookie.Value
		}
	}

	return ""
}

func (c *TestServer) Get(t *testing.T, path string) *http.Response {
	resp, err := c.Client.Get(c.Server.URL + path)

	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	return resp
}

func (c *TestServer) Post(t *testing.T, path string, form interface{}) *http.Response {
	requestForm, err := json.Marshal(form)
	require.NoError(t, err)

	resp, err := c.Client.Post(c.Server.URL+path, "application/json", bytes.NewBuffer(requestForm))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	return resp
}

func (c *TestServer) PostForm(t *testing.T, path string, form url.Values) *http.Response {
	resp, err := c.Client.PostForm(c.Server.URL+path, form)

	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	return resp
}

func (c *TestServer) Head(t *testing.T, path string) *http.Response {
	resp, err := c.Client.Head(c.Server.URL + path)
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	return resp
}

func (c *TestServer) Delete(t *testing.T, path string) *http.Response {
	req, err := http.NewRequest("DELETE", c.Server.URL+path, nil)
	require.NoError(t, err)
	resp, err := c.Client.Do(req)
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	return resp
}

func (c *TestServer) Close() {
	c.Server.Close()
}
