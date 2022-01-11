package otp_test

import (
	"context"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/core/user/otp"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	goretry "github.com/cateiru/go-retry"
	_otp "github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/require"
)

func TestOTPEnable(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	t.Setenv("ISSUER", "TestIssuer")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)
	_, err = dummy.AddUserCert(ctx, db)
	require.NoError(t, err)

	var otpData *otp.GetOTPTokenResponse

	goretry.Retry(t, func() bool {
		otpData, err = otp.GenerateOTPToken(ctx, db, dummy.UserID)
		return err == nil
	}, "OTPがbufferに格納できる")

	w, err := _otp.NewKeyFromURL(otpData.OtpToken)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		code, err := totp.GenerateCode(w.Secret(), time.Now().UTC())
		require.NoError(t, err)

		_, err = otp.SetOTP(ctx, db, dummy.UserID, otpData.Id, code)
		if err != nil {
			t.Log(err)
		}
		return err == nil
	}, "OTPをセットできる")

	goretry.Retry(t, func() bool {
		cert, err := models.GetCertificationByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return cert != nil && len(cert.OnetimePasswordBackups) == 10 && len(cert.OnetimePasswordSecret) != 0
	}, "")
}

func TestOTPDisable(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	t.Setenv("ISSUER", "TestIssuer")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy, err := tools.NewDummyUser().NewOTP()
	require.NoError(t, err)

	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)
	_, err = dummy.AddUserCert(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		passcode, err := dummy.GenOTPCode()
		require.NoError(t, err)

		err = otp.DeleteOTP(ctx, db, dummy.UserID, passcode)
		if err != nil {
			t.Log(err)
		}

		return err == nil
	}, "OTPを削除できる")

	goretry.Retry(t, func() bool {
		cert, err := models.GetCertificationByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return cert != nil && len(cert.OnetimePasswordSecret) == 0 && len(cert.OnetimePasswordBackups) == 0
	}, "OTPが正しく削除されている")
}