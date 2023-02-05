package user_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru-sso/src/config"
	"github.com/cateiru/cateiru-sso/src/core/user/password"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/handler"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/tests/tools"
	"github.com/cateiru/cateiru-sso/src/utils/secure"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func changePWServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.UserPasswordHandler)

	return mux
}

func TestPasswordChange(t *testing.T) {
	config.TestInit(t)

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

	s := tools.NewTestServer(t, changePWServer(), true)
	s.AddSession(ctx, db, dummy)

	form := password.ChangePasswordRequest{
		NewPassword: newPassword,
		OldPassword: "password",
	}

	s.Post(t, "/", form)

	// --- 確認する

	goretry.Retry(t, func() bool {
		cert, err := models.GetCertificationByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return cert != nil && secure.ValidatePW(newPassword, cert.Password, cert.Salt)
	}, "パスワードが変更されている")
}
