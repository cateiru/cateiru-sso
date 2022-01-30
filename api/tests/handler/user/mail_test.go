package user_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/core/user/mail"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/handler"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
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
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, mailServer(), true)
	s.AddSession(ctx, db, dummy)

	// ---

	resp := s.Get(t, "/")

	var element mail.ResponseMail

	err = json.Unmarshal(tools.ConvertByteResp(resp), &element)
	require.NoError(t, err)

	require.Equal(t, dummy.Mail, element.Mail)
}

func TestChangeMail(t *testing.T) {
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

	newMail := tools.NewDummyUser().Mail

	s := tools.NewTestServer(t, mailServer(), true)
	s.AddSession(ctx, db, dummy)

	// --- 認証リクエスト

	changeForm := mail.ChangeMailRequest{
		Type:    "change",
		NewMail: newMail,
	}

	s.Post(t, "/", changeForm)

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

	resp := s.Post(t, "/", changeForm)

	var respBody mail.VerifyMailResponse
	json.Unmarshal(tools.ConvertByteResp(resp), &respBody)

	require.Equal(t, newMail, respBody.NewMail)

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

	goretry.Retry(t, func() bool {
		mailCert, err := models.GetMailCertificationByMailToken(ctx, db, mailToken)
		require.NoError(t, err)

		return mailCert == nil
	}, "mail certが削除されている")
}
