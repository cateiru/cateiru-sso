package login_test

import (
	"context"
	"testing"

	"github.com/cateiru/cateiru-sso/api/core/login"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	_, err = dummy.AddUserCert(ctx, db)
	require.NoError(t, err)
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	ip := "192.0.2.0"
	userAgent := "Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion"
	form := &login.RequestFrom{
		Mail:     dummy.Mail,
		Password: "password",
	}

	var loginState *login.LoginState

	goretry.Retry(t, func() bool {
		loginState, err = login.Login(ctx, form, ip, userAgent)
		if err != nil {
			t.Log(err)
			return false
		}
		return true
	}, "ログインできる")

	require.False(t, loginState.IsOTP)
	require.NotEmpty(t, loginState.RefreshToken)
	require.NotEmpty(t, loginState.SessionToken)

	// -----

	// PWが違う場合
	form = &login.RequestFrom{
		Mail:     dummy.Mail,
		Password: "asd3as",
	}
	_, err = login.Login(ctx, form, ip, userAgent)
	require.Error(t, err)
}

func TestLoginAdmin(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	adminMail := tools.NewDummyUser().Mail
	adminPW := "ddsfe0w3sa"

	t.Setenv("ADMIN_MAIL", adminMail)
	t.Setenv("ADMIN_PASSWORD", adminPW)

	ctx := context.Background()

	ip := "192.0.2.0"
	userAgent := "Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion"
	form := &login.RequestFrom{
		Mail:     adminMail,
		Password: adminPW,
	}

	loginState, err := login.Login(ctx, form, ip, userAgent)
	require.NoError(t, err)

	require.False(t, loginState.IsOTP)
	require.NotEmpty(t, loginState.RefreshToken)
	require.NotEmpty(t, loginState.SessionToken)

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	cert, err := models.GetCertificationByMail(ctx, db, adminMail)
	require.NoError(t, err)
	require.NotNil(t, cert)
	require.Equal(t, cert.Mail, adminMail)
	user, err := models.GetUserDataByUserID(ctx, db, cert.UserId.UserId)
	require.NoError(t, err)
	require.NotNil(t, user)
	require.Equal(t, user.Mail, adminMail)
}

func TestLoginOTP(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	t.Setenv("ISSUER", "TestIssuer")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy, err := tools.NewDummyUser().NewOTP()
	require.NoError(t, err)

	_, err = dummy.AddUserCert(ctx, db)
	require.NoError(t, err)
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	ip := "192.0.2.0"
	userAgent := "Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion"
	form := &login.RequestFrom{
		Mail:     dummy.Mail,
		Password: "password",
	}

	var loginState *login.LoginState

	goretry.Retry(t, func() bool {
		loginState, err = login.Login(ctx, form, ip, userAgent)
		if err != nil {
			t.Log(err)
			return false
		}
		return true
	}, "ログインできる")

	require.NotNil(t, loginState)
	require.True(t, loginState.IsOTP)
	require.NotEmpty(t, loginState.OTPId)
}