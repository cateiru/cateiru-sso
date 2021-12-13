package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"testing"

	createaccount "github.com/cateiru/cateiru-sso/api/core/create_account"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/handler"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/stretchr/testify/require"
)

const (
	Email     = "example@example.com"
	Password  = "password"
	FirstName = "first"
	LastName  = "last"
	UserName  = "cateiru"
	Theme     = "dark"
)

func createAccountServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/create", handler.CreateHandler)
	mux.HandleFunc("/create/verify", handler.CreateVerifyHandler)
	mux.HandleFunc("/create/info", handler.CreateInfoHandler)

	return mux
}

// アカウント作成のテスト
//
//	メールアドレス、PWをpostしてメール認証開始
//	↓
//	メールトークンをクエリパラメータに含めたリンク踏む（GET）
//	↓
//	buffer tokenを使用してユーザ情報を入力し、アカウント作成
func TestCreateAccount(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	app := createAccountServer()
	server := httptest.NewServer(app)
	defer server.Close()

	jar, err := cookiejar.New(nil)
	require.NoError(t, err, "cookiejarでエラー")
	client := &http.Client{Jar: jar}

	// Step.1 ----

	createForm := createaccount.PostForm{
		Mail:       Email,
		Password:   Password,
		ReCHAPTCHA: "", // dev modeのため検証しない
	}
	form, err := json.Marshal(createForm)
	require.NoError(t, err)

	// 最初に一時的にアカウントを作成する（メール認証はまだ）
	resp, err := client.Post(server.URL+"/create", "application/json", bytes.NewBuffer(form))
	require.NoError(t, err)

	var response createaccount.Response
	err = respToJson(resp, &response)
	require.NoError(t, err)

	t.Log(response.ClientCheckToken)

	require.NotEqual(t, len(response.ClientCheckToken), 0, "ちゃんとclientCheckTokenが返ってくる")

	mailToken, err := getMailToken(ctx, response.ClientCheckToken)
	require.NoError(t, err)

	// Step.2 ----

	verifyForm := createaccount.VerifyRequestForm{
		MailToken: mailToken,
	}
	form, err = json.Marshal(verifyForm)
	require.NoError(t, err)

	// メール認証URLにアクセスする & bufferTokenのcookieが適用される
	_, err = client.Post(server.URL+"/create/verify", "application/json", bytes.NewBuffer(form))
	require.NoError(t, err)

	// Step.3 ----

	userForm := createaccount.InfoRequestForm{
		FirstName: FirstName,
		LastName:  LastName,
		UserName:  UserName,
		Theme:     Theme,
		AvatarUrl: "",
	}
	form, err = json.Marshal(userForm)
	require.NoError(t, err)

	_, err = client.Post(server.URL+"/create/info", "application/json", bytes.NewBuffer(form))
	require.NoError(t, err)

	//////
	// 入力した値、セッショントークンが正しいか検証する

	set_cookie_url, err := url.Parse(server.URL + "/create/info")
	require.NoError(t, err)

	cookies := jar.Cookies(set_cookie_url)
	var sessionToken string
	for _, cookie := range cookies {
		if cookie.Name == "session-token" {
			sessionToken = cookie.Value
		}
	}
	require.NotEmpty(t, sessionToken)

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	// セッショントークンからUserIDを取得する
	session, err := models.GetSessionToken(ctx, db, sessionToken)
	require.NoError(t, err)
	require.NotNil(t, session)

	userID := session.UserId.UserId

	// ユーザinfoを取得する
	userInfo, err := models.GetUserDataByUserID(ctx, db, userID)
	require.NoError(t, err)
	require.NotNil(t, userInfo)

	require.Equal(t, userInfo.FirstName, FirstName)
	require.Equal(t, userInfo.LastName, LastName)
	require.Equal(t, userInfo.Mail, Email)
	require.Equal(t, userInfo.Role[0], "user")
	require.Equal(t, userInfo.Theme, Theme)
	require.Equal(t, userInfo.UserName, UserName)
	require.Equal(t, userInfo.AvatarUrl, "")

	// ユーザの認証情報を取得する
	userCert, err := models.GetCertificationByMail(ctx, db, userInfo.Mail)
	require.NoError(t, err)
	require.NotNil(t, userCert)

	require.Equal(t, userCert.Mail, Email)
	require.Equal(t, userCert.UserId.UserId, userID)
	require.Equal(t, userCert.OnetimePasswordSecret, "")      // 設定していないので空
	require.Equal(t, len(userCert.OnetimePasswordBackups), 0) // 設定していないので空

}

// clientCheckTokenからmailTokenを取得し、opennewwindowをtrueにする
func getMailToken(ctx context.Context, clientCheckToken string) (string, error) {
	db, err := database.NewDatabase(ctx)
	if err != nil {
		return "", err
	}
	defer db.Close()

	entry, err := models.GetMailCertificationByCheckToken(ctx, db, clientCheckToken)
	if err != nil {
		return "", err
	}

	entry.OpenNewWindow = true

	err = entry.Add(ctx, db)
	if err != nil {
		return "", err
	}

	return entry.MailToken, nil
}

// responseをjsonにパースする
func respToJson(resp *http.Response, obj interface{}) error {
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("no 200")
	}

	buf := &bytes.Buffer{}
	buf.ReadFrom(resp.Body)

	return json.Unmarshal(buf.Bytes(), obj)
}
