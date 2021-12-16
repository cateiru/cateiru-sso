package handler_test

import (
	"context"
	"encoding/json"
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
	"github.com/stretchr/testify/require"
)

func meServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.MeHandler)

	return mux
}

func TestMe(t *testing.T) {
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

	app := meServer()
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

	var userInfo models.User
	err = json.Unmarshal(tools.ConvertByteResp(resp), &userInfo)
	require.NoError(t, err)

	require.Equal(t, userInfo.FirstName, "TestFirstName")
	require.Equal(t, userInfo.LastName, "TestLastName")
	require.Equal(t, userInfo.UserName, "TestUserName")
	require.Equal(t, userInfo.UserId.UserId, dummy.UserID)
	require.Equal(t, userInfo.Mail, dummy.Mail)
}

func TestMeNotVerify(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	app := meServer()
	server := httptest.NewServer(app)
	defer server.Close()

	jar, err := cookiejar.New(nil)
	require.NoError(t, err, "cookiejarでエラー")
	client := &http.Client{Jar: jar}

	resp, err := client.Get(server.URL + "/")
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 403, "認証情報がないので何も返さない")
}
