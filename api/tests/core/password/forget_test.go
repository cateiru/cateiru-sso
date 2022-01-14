package password_test

import (
	"context"
	"testing"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/core/password"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/cateiru/cateiru-sso/api/utils/secure"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestFrogetPW(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	newPassword := "hogehoge"

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserCert(ctx, db)
	require.NoError(t, err)

	err = password.CreateChangeMail(ctx, db, dummy.Mail)
	require.NoError(t, err)

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

	form := password.AccpetFortgetRequest{
		ForgetToken: forgetToken,
		NewPassword: newPassword,
	}

	err = password.ChangePWAccept(ctx, db, &form)
	require.NoError(t, err)

	// --- チェック

	goretry.Retry(t, func() bool {
		cert, err := models.GetCertificationByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return cert != nil && secure.ValidatePW(newPassword, cert.Password, cert.Salt)
	}, "パスワードが変更されている")

	goretry.Retry(t, func() bool {
		forget, err := models.GetPWForgetByToken(ctx, db, forgetToken)
		require.NoError(t, err)

		return forget == nil
	}, "forgetが削除されている")
}
