package password_test

import (
	"context"
	"testing"

	"github.com/cateiru/cateiru-sso/api/core/user/password"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/cateiru/cateiru-sso/api/utils/secure"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestChangePassword(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	_, err = dummy.AddUserCert(ctx, db)
	require.NoError(t, err)

	form := password.ChangePasswordRequest{
		OldPassword: "password",
		NewPassword: "hogehoge",
	}

	goretry.Retry(t, func() bool {
		err = password.ChangePassword(ctx, db, dummy.UserID, &form)
		if err != nil {
			t.Log(err)
			return false
		}
		return true
	}, "")

	goretry.Retry(t, func() bool {
		cert, err := models.GetCertificationByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return cert != nil && secure.ValidatePW("hogehoge", cert.Password, cert.Salt)
	}, "パスワードが変更されている")
}
