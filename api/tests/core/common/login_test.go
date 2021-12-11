package common_test

import (
	"context"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
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

func TestLogin(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	userId := utils.CreateID(20)

	login, err := common.LoginByUserID(ctx, db, userId)
	require.NoError(t, err)

	time.Sleep(1 * time.Second)

	entry, err := models.GetSessionToken(ctx, db, login.SessionToken)
	require.NoError(t, err)
	require.Equal(t, entry.UserId.UserId, userId)

	entryR, err := models.GetRefreshToken(ctx, db, login.RefreshToken)
	require.NoError(t, err)
	require.Equal(t, entryR.UserId.UserId, userId)
}
