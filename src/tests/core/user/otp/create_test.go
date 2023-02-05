package otp_test

import (
	"context"
	"testing"

	"github.com/cateiru/cateiru-sso/src/config"
	"github.com/cateiru/cateiru-sso/src/core/user/otp"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/tests/tools"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestCreateOTP(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	var otpData *otp.GetOTPTokenResponse

	goretry.Retry(t, func() bool {
		otpData, err = otp.GenerateOTPToken(ctx, db, dummy.UserID)
		return err == nil
	}, "OTPがbufferに格納できる")

	goretry.Retry(t, func() bool {
		entity, err := models.GetOTPBufferByID(ctx, db, otpData.Id)
		require.NoError(t, err)

		return entity != nil && entity.UserId.UserId == dummy.UserID
	}, "ちゃんと格納できている")
}
