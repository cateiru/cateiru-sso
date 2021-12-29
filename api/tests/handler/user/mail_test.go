package user_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/core/user/mail"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/handler"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/cateiru/cateiru-sso/api/utils/net"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
	"google.golang.org/api/iterator"
)

func mailServer() *http.ServeMux {
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

	app := mailServer()
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

func TestChangeMail(t *testing.T) {
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

	newMail := tools.NewDummyUser().Mail

	app := mailServer()
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

	// --- 認証リクエスト

	changeForm := mail.ChangeMailRequest{
		Type:    "change",
		NewMail: newMail,
	}
	form, err := json.Marshal(changeForm)
	require.NoError(t, err)

	resp, err := client.Post(server.URL+"/", "application/json", bytes.NewBuffer(form))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	// --- メールトークンをDBから抜いてくる

	var mailToken string

	query := datastore.NewQuery("MailCertification").Filter("mail =", newMail)
	goretry.Retry(t, func() bool {
		iter := db.Run(ctx, query)
		var entity models.MailCertification
		_, err := iter.Next(&entity)

		// 要素がなにもない場合
		if err == iterator.Done {
			return false
		}
		require.NoError(t, err)

		mailToken = entity.MailToken

		return entity.ChangeMailMode && entity.UserId == dummy.UserID
	}, "dbに要素がある")

	require.NotEmpty(t, mailToken)

	// --- 認証とメールアドレス変更

	changeForm = mail.ChangeMailRequest{
		Type:      "verify",
		MailToken: mailToken,
	}
	form, err = json.Marshal(changeForm)
	require.NoError(t, err)

	resp, err = client.Post(server.URL+"/", "application/json", bytes.NewBuffer(form))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 200)

	// --- 確認

	goretry.Retry(t, func() bool {
		cert, err := models.GetCertificationByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return cert != nil && cert.Mail == newMail
	}, "certのメールアドレスが変更されている")

	goretry.Retry(t, func() bool {
		info, err := models.GetUserDataByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return info != nil && info.Mail == newMail
	}, "userInfoのメールアドレスが変更されている")
}
