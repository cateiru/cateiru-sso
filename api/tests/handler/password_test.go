package handler_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru-sso/api/core/password"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/handler"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/cateiru/cateiru-sso/api/utils/secure"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func forgetServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/forget", handler.PasswordForgetHandler)
	mux.HandleFunc("/forget/accept", handler.PasswordForgetAcceptHandler)

	return mux
}

func TestForgetPassword(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	newPassword := "hogehoge"

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserCert(ctx, db)
	require.NoError(t, err)
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, forgetServer(), false)

	form := password.ForgetRequest{
		Mail: dummy.Mail,
	}
	s.Post(t, "/forget", form)

	var forgetToken string

	goretry.Retry(t, func() bool {
		entity, err := models.GetPWForgetByMail(ctx, db, dummy.Mail)
		require.NoError(t, err)

		if len(entity) == 0 {
			return false
		}
		forgetToken = entity[0].ForgetToken
		return true
	}, "")
	require.NotEmpty(t, forgetToken)

	acceptForm := password.AccpetFortgetRequest{
		ForgetToken: forgetToken,
		NewPassword: newPassword,
	}
	s.Post(t, "/forget/accept", acceptForm)

	// --- 確認する

	goretry.Retry(t, func() bool {
		cert, err := models.GetCertificationByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return cert != nil && secure.ValidatePW(newPassword, cert.Password, cert.Salt)
	}, "パスワードが変更されている")
}
