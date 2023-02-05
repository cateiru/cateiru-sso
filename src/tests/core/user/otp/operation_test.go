package otp_test

import (
	"context"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/src/config"
	"github.com/cateiru/cateiru-sso/src/core/user/otp"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/tests/tools"
	goretry "github.com/cateiru/go-retry"
	_otp "github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/require"
)

func TestOTPEnable(t *testing.T) {
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

	goretry.Retry(t, func() bool {
		buff, err := models.GetOTPBufferByID(ctx, db, otpData.Id)
		require.NoError(t, err)

		return buff == nil
	}, "bufferは削除されている")
}

func TestOTPDisable(t *testing.T) {
	config.TestInit(t)

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
