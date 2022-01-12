package handler_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/handler"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func logoutServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.LogoutHandler)

	return mux
}

func TestLogout(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, logoutServer(), true)
	s.AddSession(ctx, db, dummy)

	s.Get(t, "/")

	// --- チェックする

	// cookieが削除されているか確認する
	for _, cookie := range s.Jar.Cookies(s.Url) {
		require.NotEqual(t, cookie.Name, "session-token")
		require.NotEqual(t, cookie.Name, "refresh-token")
	}

	goretry.Retry(t, func() bool {
		sessionToken, err := models.GetSessionToken(ctx, db, s.SessionToken)
		require.NoError(t, err)
		refreshToken, err := models.GetRefreshToken(ctx, db, s.RefreshToken)
		require.NoError(t, err)

		return sessionToken == nil && refreshToken == nil
	}, "トークン類はDBから削除されている")
}

func TestDelete(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)
	_, err = dummy.AddUserCert(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, logoutServer(), true)
	s.AddSession(ctx, db, dummy)

	// --- 削除する

	s.Delete(t, "/")

	// --- チェックする

	// cookieが削除されているか確認する
	for _, cookie := range s.Jar.Cookies(s.Url) {
		require.NotEqual(t, cookie.Name, "session-token")
		require.NotEqual(t, cookie.Name, "refresh-token")
	}

	goretry.Retry(t, func() bool {
		sessionToken, err := models.GetSessionToken(ctx, db, s.SessionToken)
		require.NoError(t, err)
		refreshToken, err := models.GetRefreshToken(ctx, db, s.RefreshToken)
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
