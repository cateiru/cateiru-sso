package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	createaccount "github.com/cateiru/cateiru-sso/api/core/create_account"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/handler"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/posener/wstest"
	"github.com/stretchr/testify/require"
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

	var (
		Email     = fmt.Sprintf("%s@example.com", utils.CreateID(5))
		Password  = "password"
		FirstName = "first"
		LastName  = "last"
		UserName  = "cateiru"
		Theme     = "dark"
	)

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
	require.Equal(t, resp.StatusCode, 200)

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
	resp, err = client.Post(server.URL+"/create/verify", "application/json", bytes.NewBuffer(form))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	// Step.3 ----

	// cookieを設定する
	resp, err = client.Head(fmt.Sprintf("%s/create/verify?token=%s", server.URL, response.ClientCheckToken))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	// Step.4 ----

	userForm := createaccount.InfoRequestForm{
		FirstName: FirstName,
		LastName:  LastName,
		UserName:  UserName,
		Theme:     Theme,
		AvatarUrl: "",
	}
	form, err = json.Marshal(userForm)
	require.NoError(t, err)

	resp, err = client.Post(server.URL+"/create/info", "application/json", bytes.NewBuffer(form))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	//////
	// 入力した値、セッショントークンが正しいか検証する

	set_cookie_url, err := url.Parse(server.URL + "/create/info")
	require.NoError(t, err)

	cookies := jar.Cookies(set_cookie_url)
	var sessionToken string
	for _, cookie := range cookies {
		if cookie.Name == "session-token" {
			sessionToken = cookie.Value
			break
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

func TestOther(t *testing.T) {
	app := createAccountServer()
	server := httptest.NewServer(app)
	defer server.Close()

	client := &http.Client{}

	resp, err := http.Get(server.URL + "/create")
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	req, err := http.NewRequest("DELETE", server.URL+"/create/verify", nil)
	require.NoError(t, err)
	resp, err = client.Do(req)
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)
}

func TestObserve(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	form := &createaccount.PostForm{
		Mail:       fmt.Sprintf("%s@example.com", utils.CreateID(4)),
		Password:   "password",
		ReCHAPTCHA: "",
	}
	ip := "192.168.1.1"

	clientCheckToken, err := createaccount.CreateTemporaryAccount(ctx, form, ip)
	require.NoError(t, err)

	////

	server := createAccountServer()

	d := wstest.NewDialer(server)

	c, resp, err := d.Dial(fmt.Sprintf("ws://whatever/create/verify?cct=%s", clientCheckToken), nil)
	require.NoError(t, err)
	got, want := resp.StatusCode, http.StatusSwitchingProtocols
	require.Equal(t, got, want)

	go verifyMail(ctx, t, clientCheckToken)

	// 受信待機
	var respm bool
	err = c.ReadJSON(&respm)
	require.NoError(t, err)

	// 返ってくる = メール認証が完了したためclient側からwsをcloseする
	err = c.Close()
	require.NoError(t, err)

	// response messageは`true`が返る
	require.True(t, respm)
}

// clientCheckTokenからmailTokenを取得する
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

func verifyMail(ctx context.Context, t *testing.T, clientCheckToken string) {
	// 3秒間待機する: WSで待機するため
	time.Sleep(3 * time.Second)

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	cert, err := models.GetMailCertificationByCheckToken(ctx, db, clientCheckToken)
	require.NoError(t, err)
	require.NotNil(t, cert)

	t.Logf("verify mailToken: %s", cert.MailToken)
	cert.Verify = true

	err = cert.Add(ctx, db)
	require.NoError(t, err)
}
