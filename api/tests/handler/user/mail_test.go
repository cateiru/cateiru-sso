package user_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/core/user/mail"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/handler"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	"github.com/stretchr/testify/require"
)

func getMailServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.UserMailHandler)

	return mux
}

func TestGetMail(t *testing.T) {
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

	app := getMailServer()
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

	// ---

	resp, err := client.Get(server.URL + "/")
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	var element mail.ResponseMail

	err = json.Unmarshal(tools.ConvertByteResp(resp), &element)
	require.NoError(t, err)

	require.Equal(t, dummy.Mail, element.Mail)
}
