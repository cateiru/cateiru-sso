package login_test

import (
	"context"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/core/login"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/cateiru/cateiru-sso/api/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestOTPLogin(t *testing.T) {
	config.TestInit(t)

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

	id := utils.CreateID(0)

	buffer := &models.OnetimePasswordBuffer{
		Id:      id,
		IsLogin: true,

		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 10,
		},

		UserId: models.UserId{
			UserId: dummy.UserID,
		},
	}

	err = buffer.Add(ctx, db)
	require.NoError(t, err)

	ip := "192.0.2.0"
	userAgent := "Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion"

	passcode, err := dummy.GenOTPCode()
	require.NoError(t, err)

	c := &common.Cert{
		Ip:        ip,
		UserAgent: userAgent,
	}

	// 違うパスワードではできない
	dummyPasscode := "hogehoge"
	err = login.LoginOTP(ctx, id, dummyPasscode, c)
	require.Error(t, err)

	goretry.Retry(t, func() bool {
		err := login.LoginOTP(ctx, id, passcode, c)
		if err != nil {
			t.Log(err)
			return false
		}
		return len(c.SessionToken) != 0
	}, "ログインできる")

	// ---

	// 同じidで複数回はできない
	err = login.LoginOTP(ctx, id, passcode, c)
	require.Error(t, err)
}

func TestFailedID(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	id := "hogehoge"

	c := &common.Cert{}

	err := login.LoginOTP(ctx, id, "", c)
	require.Error(t, err)
}
