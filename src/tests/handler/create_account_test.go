package handler_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/src/config"
	createaccount "github.com/cateiru/cateiru-sso/src/core/create_account"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/handler"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/tests/tools"
	"github.com/cateiru/cateiru-sso/src/utils"
	goretry "github.com/cateiru/go-retry"
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
	config.TestInit(t)

	var (
		Email     = fmt.Sprintf("%s@example.com", utils.CreateID(5))
		Password  = "password"
		FirstName = "first"
		LastName  = "last"
		UserName  = utils.CreateID(5)
		Theme     = "dark"
	)

	ctx := context.Background()

	s := tools.NewTestServer(t, createAccountServer(), true)

	// Step.1 ----

	createForm := createaccount.PostForm{
		Mail:      Email,
		ReCAPTCHA: "", // dev modeのため検証しない
	}

	// 最初に一時的にアカウントを作成する（メール認証はまだ）
	resp := s.Post(t, "/create", createForm)

	var response createaccount.Response
	err := respToJson(resp, &response)
	require.NoError(t, err)

	t.Log(response.ClientToken)

	require.NotEqual(t, len(response.ClientToken), 0, "ちゃんとclientTokenが返ってくる")

	mailToken := getMailToken(t, ctx, response.ClientToken)

	// Step.2 ----

	verifyForm := createaccount.VerifyRequestForm{
		MailToken: mailToken,
	}

	// メール認証URLにアクセスする & bufferTokenのcookieが適用される
	s.Post(t, "/create/verify", verifyForm)

	time.Sleep(1 * time.Second)

	// Step.4 ----

	userForm := createaccount.InfoRequestForm{
		ClientToken: response.ClientToken,

		FirstName: FirstName,
		LastName:  LastName,
		UserName:  UserName,
		Theme:     Theme,
		AvatarUrl: "",

		Password: Password,
	}

	s.Post(t, "/create/info", userForm)

	//////
	// 入力した値、セッショントークンが正しいか検証する

	s.FindCookies(t, []string{"session-token", "refresh-token"})

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	// セッショントークンからUserIDを取得する
	session, err := models.GetSessionToken(ctx, db, s.GetCookie("session-token"))
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
	require.Equal(t, userInfo.Theme, Theme)
	require.Equal(t, userInfo.UserName, UserName)
	require.Equal(t, userInfo.AvatarUrl, "")

	// ユーザのroleを取得する
	role, err := models.GetRoleByUserID(ctx, db, userID)
	require.NoError(t, err)
	require.NotNil(t, role)

	require.Equal(t, role.Role[0], "user")

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
	config.TestInit(t)

	ctx := context.Background()

	form := &createaccount.PostForm{
		Mail:      fmt.Sprintf("%s@example.com", utils.CreateID(4)),
		ReCAPTCHA: "",
	}
	ip := "192.168.1.1"

	clientToken, err := createaccount.CreateTemporaryAccount(ctx, form, ip)
	require.NoError(t, err)

	////

	server := createAccountServer()

	d := wstest.NewDialer(server)

	c, resp, err := d.Dial(fmt.Sprintf("ws://whatever/create/verify?cct=%s", clientToken), nil)
	require.NoError(t, err)
	got, want := resp.StatusCode, http.StatusSwitchingProtocols
	require.Equal(t, got, want)

	go verifyMail(ctx, t, clientToken)

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

// clientTokenからmailTokenを取得する
func getMailToken(t *testing.T, ctx context.Context, clientToken string) string {
	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	var mailToken string

	goretry.Retry(t, func() bool {
		entry, err := models.GetMailCertificationByClientToken(ctx, db, clientToken)
		require.NoError(t, err)

		if entry != nil {
			mailToken = entry.MailToken
			return true
		}

		return false
	}, "")

	return mailToken
}

// responseをjsonにパースする
func respToJson(resp *http.Response, obj interface{}) error {
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("no 200")
	}

	return json.Unmarshal(tools.ConvertByteResp(resp), obj)
}

func verifyMail(ctx context.Context, t *testing.T, clientToken string) {
	// 3秒間待機する: WSで待機するため
	time.Sleep(3 * time.Second)

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	cert, err := models.GetMailCertificationByClientToken(ctx, db, clientToken)
	require.NoError(t, err)
	require.NotNil(t, cert)

	t.Logf("verify mailToken: %s", cert.MailToken)
	cert.Verify = true

	err = cert.Add(ctx, db)
	require.NoError(t, err)
}
