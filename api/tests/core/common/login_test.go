package common_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/logging"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func server() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/s", success)

	return mux
}

func success(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	c := common.NewCert(w, r)
	c.Login(ctx, db)

	// アクセスID指定されている
	if c.AccessID == "" {
		logging.Sugar.Error("acessId")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// ユーザID取得できている
	if c.UserId == "" {
		logging.Sugar.Error("acessId")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(c.UserId))
}

func TestNewLogin(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	c := &common.Cert{
		Ip:        "192.168.1.0",
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36",
	}

	err = c.NewLogin(ctx, db, dummy.UserID)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		session, err := models.GetSessionToken(ctx, db, c.SessionToken)
		require.NoError(t, err)

		return session != nil
	}, "session tokenが保存されている")

	goretry.Retry(t, func() bool {
		refresh, err := models.GetRefreshToken(ctx, db, c.RefreshToken)
		require.NoError(t, err)

		return refresh != nil
	}, "refres tokenが保存されている")

	goretry.Retry(t, func() bool {
		histories, err := models.GetAllLoginHistory(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return len(histories) == 1
	}, "ログイン履歴が格納されている")
}

func TestSuccessLogin(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, server(), true)
	defer s.Close()
	err = s.AddSession(ctx, db, dummy)
	require.NoError(t, err)

	resp := s.Get(t, "/s")

	userId := tools.ConvertResp(resp)
	require.Equal(t, userId, dummy.UserID)

	s.FindCookies(t, []string{"session-token", "refresh-token"})
}

func TestSuccessRefresh(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, server(), true)
	defer s.Close()

	// refresh-tokenのみ有効する
	_, refresh, err := dummy.AddLoginToken(ctx, db, time.Now())
	require.NoError(t, err)

	exp := net.NewCookieMinutsExp(3)
	tools.SetCookie(s.Jar, "refresh-token", refresh, exp, s.Url)

	resp := s.Get(t, "/s")

	userId := tools.ConvertResp(resp)
	require.Equal(t, userId, dummy.UserID)

	s.FindCookies(t, []string{"session-token", "refresh-token"})
}

func TestFailed(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, server(), true)
	defer s.Close()

	resp, err := s.Client.Get(s.Server.URL + "/s")
	require.NoError(t, err)

	require.Equal(t, resp.StatusCode, 500)
}

func TestExistSessionRefreshLogin(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, server(), true)
	defer s.Close()

	// refresh-tokenのみ有効する
	_, refresh, err := dummy.AddLoginToken(ctx, db, time.Now())
	require.NoError(t, err)

	exp := net.NewCookieMinutsExp(3)
	tools.SetCookie(s.Jar, "refresh-token", refresh, exp, s.Url)
	tools.SetCookie(s.Jar, "session-token", "dummy-token", exp, s.Url)

	resp := s.Get(t, "/s")

	userId := tools.ConvertResp(resp)
	require.Equal(t, userId, dummy.UserID)

	s.FindCookies(t, []string{"session-token", "refresh-token"})
}

func TestExpiredSession(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, server(), true)
	defer s.Close()

	// refresh-tokenのみ有効する
	session, refresh, err := dummy.AddLoginToken(ctx, db, time.Now())
	require.NoError(t, err)

	exp := net.NewCookieMinutsExp(3)
	expS := net.NewCookieMinutsExp(-10)
	tools.SetCookie(s.Jar, "refresh-token", refresh, exp, s.Url)
	tools.SetCookie(s.Jar, "session-token", session, expS, s.Url)

	resp := s.Get(t, "/s")

	userId := tools.ConvertResp(resp)
	require.Equal(t, userId, dummy.UserID)

	s.FindCookies(t, []string{"session-token", "refresh-token"})
}

func TestExpiredRefresh(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, server(), true)
	defer s.Close()

	// refresh-tokenのみ有効する
	_, refresh, err := dummy.AddLoginToken(ctx, db, time.Now())
	require.NoError(t, err)

	exp := net.NewCookieMinutsExp(3)
	expS := net.NewCookieMinutsExp(-10)
	tools.SetCookie(s.Jar, "refresh-token", refresh, expS, s.Url)
	tools.SetCookie(s.Jar, "session-token", "dummy", exp, s.Url)

	resp, err := s.Client.Get(s.Server.URL + "/s")
	require.NoError(t, err)

	require.Equal(t, resp.StatusCode, 500)
}
