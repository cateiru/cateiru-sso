package handler_test

import (
	"context"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/handler"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func logoutServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.LogoutHandler)

	return mux
}

func TestLogout(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)
	session, refresh, err := dummy.AddLoginToken(ctx, db, time.Now())
	require.NoError(t, err)

	app := logoutServer()
	server := httptest.NewServer(app)
	defer server.Close()

	jar, err := cookiejar.New(nil)
	require.NoError(t, err, "cookiejarでエラー")
	client := &http.Client{Jar: jar}

	url, err := url.Parse(server.URL + "/")
	require.NoError(t, err)

	exp := net.NewCookieMinutsExp(3)
	tools.SetCookie(jar, "session-token", session, exp, url)
	tools.SetCookie(jar, "refresh-token", refresh, exp, url)

	resp, err := client.Get(server.URL + "/")
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	// --- チェックする

	// cookieが削除されているか確認する
	for _, cookie := range jar.Cookies(url) {
		require.NotEqual(t, cookie.Name, "session-token")
		require.NotEqual(t, cookie.Name, "refresh-token")
	}

	goretry.Retry(t, func() bool {
		sessionToken, err := models.GetSessionToken(ctx, db, session)
		require.NoError(t, err)
		refreshToken, err := models.GetRefreshToken(ctx, db, refresh)
		require.NoError(t, err)

		return sessionToken == nil && refreshToken == nil
	}, "トークン類はDBから削除されている")
}

func TestDelete(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)
	_, err = dummy.AddUserCert(ctx, db)
	require.NoError(t, err)

	session, refresh, err := dummy.AddLoginToken(ctx, db, time.Now())
	require.NoError(t, err)

	app := logoutServer()
	server := httptest.NewServer(app)
	defer server.Close()

	jar, err := cookiejar.New(nil)
	require.NoError(t, err, "cookiejarでエラー")
	client := &http.Client{Jar: jar}

	url, err := url.Parse(server.URL + "/")
	require.NoError(t, err)

	exp := net.NewCookieMinutsExp(3)
	tools.SetCookie(jar, "session-token", session, exp, url)
	tools.SetCookie(jar, "refresh-token", refresh, exp, url)

	// --- 削除する

	req, err := http.NewRequest("DELETE", server.URL+"/", nil)
	require.NoError(t, err)
	resp, err := client.Do(req)
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	// --- チェックする

	// cookieが削除されているか確認する
	for _, cookie := range jar.Cookies(url) {
		require.NotEqual(t, cookie.Name, "session-token")
		require.NotEqual(t, cookie.Name, "refresh-token")
	}

	goretry.Retry(t, func() bool {
		sessionToken, err := models.GetSessionToken(ctx, db, session)
		require.NoError(t, err)
		refreshToken, err := models.GetRefreshToken(ctx, db, refresh)
		require.NoError(t, err)

		return sessionToken == nil && refreshToken == nil
	}, "トークン類はDBから削除されている")

	goretry.Retry(t, func() bool {
		info, err := models.GetUserDataByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)
		cert, err := models.GetCertificationByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return info == nil && cert == nil
	}, "ユーザの認証情報が消えている")
}
