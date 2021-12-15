package common_test

import (
	"bytes"
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
	"github.com/cateiru/cateiru-sso/api/utils/net"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

const SESSION_TOKEN = "hogehoge"
const REFRESH_TOKEN = "hugahuga"

func loginServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", testHandler)
	mux.HandleFunc("/u", getUserHandler)

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

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	db, err := database.NewDatabase(ctx)
	if err != nil {
		net.ResponseError(w, err)
		return
	}
	defer db.Close()

	userId, err := common.GetUserID(ctx, db, w, r)
	if err != nil {
		net.ResponseError(w, err)
		return
	}

	w.Write([]byte(userId))
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

	ip := "198.51.100.0"
	userAgent := "Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion"

	login, err := common.LoginByUserID(ctx, db, userId, ip, userAgent)
	require.NoError(t, err)

	// 初回のみ
	goretry.Retry(t, func() bool {
		entry, err := models.GetSessionToken(ctx, db, login.SessionToken)
		require.NoError(t, err)

		return entry.UserId.UserId == userId
	}, "sessionTokenがある")

	entryR, err := models.GetRefreshToken(ctx, db, login.RefreshToken)
	require.NoError(t, err)
	require.Equal(t, entryR.UserId.UserId, userId)
}

// session-tokenのcookieからuser idを取得する
func TestGetUserIdSuccess(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	userId := utils.CreateID(30)

	sessionToken, refreshToken, err := setToken(ctx, t, db, userId, time.Now())
	require.NoError(t, err)

	// -----

	app := loginServer()
	server := httptest.NewServer(app)
	defer server.Close()

	jar, err := cookiejar.New(nil)
	require.NoError(t, err, "cookiejarでエラー")
	client := &http.Client{Jar: jar}
	url, err := url.Parse(server.URL + "/")
	require.NoError(t, err)

	// cookieをセットする

	sessionExp := net.NewCookieMinutsExp(3)
	setCookie(jar, "session-token", sessionToken, sessionExp, url)

	refreshExp := net.NewCookieMinutsExp(3)
	setCookie(jar, "refresh-token", refreshToken, refreshExp, url)

	// -----

	resp, err := client.Get(server.URL + "/u")
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	body := convertResp(resp)
	require.Equal(t, body, userId)
}

// refresh-tokenからsession-tokenを作成し、user idを取得する
func TestGetUseIdRefresh(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	userId := utils.CreateID(30)

	_, refreshToken, err := setToken(ctx, t, db, userId, time.Now())
	require.NoError(t, err)

	// -----

	app := loginServer()
	server := httptest.NewServer(app)
	defer server.Close()

	jar, err := cookiejar.New(nil)
	require.NoError(t, err, "cookiejarでエラー")
	client := &http.Client{Jar: jar}
	url, err := url.Parse(server.URL + "/")
	require.NoError(t, err)

	// refresh-tokenのみ、cookieをセットする
	refreshExp := net.NewCookieMinutsExp(3)
	setCookie(jar, "refresh-token", refreshToken, refreshExp, url)

	// -----

	resp, err := client.Get(server.URL + "/u")
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	body := convertResp(resp)
	require.Equal(t, body, userId)

	cookies := jar.Cookies(url)

	sessionTokenFindFlag := false
	refreshTokenFindFlag := false
	for _, cookie := range cookies {
		if cookie.Name == "session-token" {
			sessionTokenFindFlag = true
			require.NotEmpty(t, cookie.Value)
			break
		} else if cookie.Name == "refresh-token" {
			refreshTokenFindFlag = true
			require.NotEmpty(t, cookie.Value)
			// あたらくsession-tokenを作成したのでrefresh-tokenの値は更新される
			require.NotEqual(t, cookie.Value, refreshToken, "refresh-tokenの値が更新されている")
		}
	}
	require.True(t, sessionTokenFindFlag, "session-tokenのcookieがある")
	require.True(t, refreshTokenFindFlag, "refresh-tokenのcookieがある")
}

// session-token、refresh-tokenは存在しない
func TestGetUserNoTokens(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	app := loginServer()
	server := httptest.NewServer(app)
	defer server.Close()

	jar, err := cookiejar.New(nil)
	require.NoError(t, err, "cookiejarでエラー")
	client := &http.Client{Jar: jar}
	require.NoError(t, err)

	resp, err := client.Get(server.URL + "/u")
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 403)
}

// session-tokenは存在するが中の値が違う
func TestGetUserNotSession(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	userId := utils.CreateID(30)

	_, refreshToken, err := setToken(ctx, t, db, userId, time.Now())
	require.NoError(t, err)

	// -----

	app := loginServer()
	server := httptest.NewServer(app)
	defer server.Close()

	jar, err := cookiejar.New(nil)
	require.NoError(t, err, "cookiejarでエラー")
	client := &http.Client{Jar: jar}
	url, err := url.Parse(server.URL + "/")
	require.NoError(t, err)

	// cookieをセットする

	sessionExp := net.NewCookieMinutsExp(3)
	setCookie(jar, "session-token", "hogehoge", sessionExp, url) // session-tokenに違う値をセットする

	refreshExp := net.NewCookieMinutsExp(3)
	setCookie(jar, "refresh-token", refreshToken, refreshExp, url)

	// -----

	resp, err := client.Get(server.URL + "/u")
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200) // refresh-tokenからログインする

	body := convertResp(resp)
	require.Equal(t, body, userId)

	cookies := jar.Cookies(url)

	sessionTokenFindFlag := false
	refreshTokenFindFlag := false
	for _, cookie := range cookies {
		t.Logf("%s : %s", cookie.Name, cookie.Value)
		if cookie.Name == "session-token" {
			sessionTokenFindFlag = true
			require.NotEmpty(t, cookie.Value)
		} else if cookie.Name == "refresh-token" {
			refreshTokenFindFlag = true
			require.NotEmpty(t, cookie.Value)
			// あたらくsession-tokenを作成したのでrefresh-tokenの値は更新される
			require.NotEqual(t, cookie.Value, refreshToken, "refresh-tokenの値が更新されている")
		}
	}
	require.True(t, sessionTokenFindFlag, "session-tokenのcookieがある")
	require.True(t, refreshTokenFindFlag, "refresh-tokenのcookieがある")
}

// session-tokenはなく、refresh-tokenはあるが中の値が違う
func TestNotExistSession(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	userId := utils.CreateID(30)

	_, refreshToken, err := setToken(ctx, t, db, userId, time.Now())
	require.NoError(t, err)

	// -----

	app := loginServer()
	server := httptest.NewServer(app)
	defer server.Close()

	jar, err := cookiejar.New(nil)
	require.NoError(t, err, "cookiejarでエラー")
	client := &http.Client{Jar: jar}
	url, err := url.Parse(server.URL + "/")
	require.NoError(t, err)

	// cookieをセットする（refresh-tokenのみ）
	refreshExp := net.NewCookieMinutsExp(3)
	setCookie(jar, "refresh-token", refreshToken, refreshExp, url)

	// -----

	resp, err := client.Get(server.URL + "/u")
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200) // refresh-tokenからログインする

	body := convertResp(resp)
	require.Equal(t, body, userId)

	cookies := jar.Cookies(url)

	sessionTokenFindFlag := false
	refreshTokenFindFlag := false
	for _, cookie := range cookies {
		if cookie.Name == "session-token" {
			sessionTokenFindFlag = true
			require.NotEmpty(t, cookie.Value)
			break
		} else if cookie.Name == "refresh-token" {
			refreshTokenFindFlag = true
			require.NotEmpty(t, cookie.Value)
			// あたらくsession-tokenを作成したのでrefresh-tokenの値は更新される
			require.NotEqual(t, cookie.Value, refreshToken, "refresh-tokenの値が更新されている")
		}
	}
	require.True(t, sessionTokenFindFlag, "session-tokenのcookieがある")
	require.True(t, refreshTokenFindFlag, "refresh-tokenのcookieがある")
}

// session-tokenの有効期限がサーバー上で切れている
func TestExpiredSession(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	userId := utils.CreateID(30)

	// 有効期限はsession-tokenは6時間、refresh-tokenは7日間であるため、
	// 現在の時間を+24hしてsession-tokenのみ有効期限切れにする
	now := time.Now().Add(time.Duration(-24) * time.Hour)
	sessionToken, refreshToken, err := setToken(ctx, t, db, userId, now)
	require.NoError(t, err)

	// -----

	app := loginServer()
	server := httptest.NewServer(app)
	defer server.Close()

	jar, err := cookiejar.New(nil)
	require.NoError(t, err, "cookiejarでエラー")
	client := &http.Client{Jar: jar}
	url, err := url.Parse(server.URL + "/")
	require.NoError(t, err)

	// cookieをセットする

	sessionExp := net.NewCookieMinutsExp(3)
	setCookie(jar, "session-token", sessionToken, sessionExp, url)

	refreshExp := net.NewCookieMinutsExp(3)
	setCookie(jar, "refresh-token", refreshToken, refreshExp, url)

	// -----

	resp, err := client.Get(server.URL + "/u")
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200) // session-tokenが有効期限切れでもrefresh-tokenを使用してログインする

	body := convertResp(resp)
	require.Equal(t, body, userId)
}

// session-token、refresh-tokenの有効期限がサーバー上で切れている
func TestExpiredRefresh(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	userId := utils.CreateID(30)

	// 有効期限はsession-tokenは6時間、refresh-tokenは7日間であるため、
	// 現在の時間を+10*24hしてどちらも有効期限切れにする
	now := time.Now().Add(time.Duration(-10*24) * time.Hour)
	sessionToken, refreshToken, err := setToken(ctx, t, db, userId, now)
	require.NoError(t, err)

	// -----

	app := loginServer()
	server := httptest.NewServer(app)
	defer server.Close()

	jar, err := cookiejar.New(nil)
	require.NoError(t, err, "cookiejarでエラー")
	client := &http.Client{Jar: jar}
	url, err := url.Parse(server.URL + "/")
	require.NoError(t, err)

	// cookieをセットする

	sessionExp := net.NewCookieMinutsExp(3)
	setCookie(jar, "session-token", sessionToken, sessionExp, url)

	refreshExp := net.NewCookieMinutsExp(3)
	setCookie(jar, "refresh-token", refreshToken, refreshExp, url)

	// -----

	resp, err := client.Get(server.URL + "/u")
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 403)

	cookies := jar.Cookies(url)

	sessionTokenFindFlag := false
	refreshTokenFindFlag := false
	for _, cookie := range cookies {
		t.Logf("%s : %s", cookie.Name, cookie.Value)
		if cookie.Name == "session-token" && len(cookie.Value) != 0 {
			sessionTokenFindFlag = true
		} else if cookie.Name == "refresh-token" && len(cookie.Value) != 0 {
			refreshTokenFindFlag = true
		}
	}
	require.False(t, sessionTokenFindFlag, "session-tokenのcookieは削除済")
	require.False(t, refreshTokenFindFlag, "refresh-tokenのcookieは削除済")
}

// トークンをセットする
//
// 有効期限:
//	session-token: 6時間
//	refresh-token: 7日間
func setToken(ctx context.Context, t *testing.T, db *database.Database, userId string, now time.Time) (string, string, error) {
	sessionToken := utils.CreateID(0)
	refreshToken := utils.CreateID(0)

	session := &models.SessionInfo{
		SessionToken: sessionToken,

		Period: models.Period{
			CreateDate: now,
			PeriodHour: 6,
		},
		UserId: models.UserId{
			UserId: userId,
		},
	}
	refresh := &models.RefreshInfo{
		RefreshToken: refreshToken,
		SessionToken: sessionToken,

		Period: models.Period{
			CreateDate: now,
			PeriodDay:  7,
		},
		UserId: models.UserId{
			UserId: userId,
		},
	}

	if err := session.Add(ctx, db); err != nil {
		return "", "", err
	}
	if err := refresh.Add(ctx, db); err != nil {
		return "", "", err
	}

	return sessionToken, refreshToken, nil
}

// cookieを設定する
func setCookie(jar *cookiejar.Jar, key string, value string, exp *net.CookieExp, url *url.URL) {
	cookie := &http.Cookie{
		Name:  key,
		Value: value,

		Secure:   false,
		Path:     "/",
		Domain:   "",
		HttpOnly: false,
		SameSite: http.SameSiteDefaultMode,
	}

	if !exp.IsSession {
		cookie.Expires = time.Now().Add(exp.GetTime())
		cookie.MaxAge = exp.GetNum()
	}

	jar.SetCookies(url, []*http.Cookie{cookie})
}

// responseをstringに変換する
func convertResp(resp *http.Response) string {
	defer resp.Body.Close()

	buf := &bytes.Buffer{}
	buf.ReadFrom(resp.Body)

	return buf.String()
}
