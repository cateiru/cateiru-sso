package common_test

import (
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/stretchr/testify/require"
)

const SESSION_TOKEN = "hogehoge"
const REFRESH_TOKEN = "hugahuga"

func loginServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", testHandler)

	return mux
}

// session tokenとrefresh tokenをcookieにセットする
func testHandler(w http.ResponseWriter, r *http.Request) {
	login := &common.LoginTokens{
		SessionToken: SESSION_TOKEN,
		RefreshToken: REFRESH_TOKEN,
	}

	common.LoginSetCookie(w, login)
}

func TestLoginSetCookie(t *testing.T) {
	app := loginServer()
	server := httptest.NewServer(app)
	defer server.Close()

	u, err := url.Parse(server.URL)
	require.NoError(t, err)

	jar, err := cookiejar.New(nil)
	require.NoError(t, err, "cookiejarでエラー")
	client := &http.Client{Jar: jar}

	_, err = client.Get(server.URL + "/")
	require.NoError(t, err)

	cookies := jar.Cookies(u)

	sessionTokenFlag := false
	refreshTokenFlag := false
	for _, cookie := range cookies {
		if cookie.Name == "session-token" {
			require.Equal(t, cookie.Value, SESSION_TOKEN)
			sessionTokenFlag = true
		}
		if cookie.Name == "refresh-token" {
			require.Equal(t, cookie.Value, REFRESH_TOKEN)
			refreshTokenFlag = true
		}
	}

	require.True(t, sessionTokenFlag, "sessionToken")
	require.True(t, refreshTokenFlag, "refreshToken")
}
