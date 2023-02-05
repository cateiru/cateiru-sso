package login_test

import (
	"context"
	"net/http"
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

		FailedCount: 0,

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
		Ip: ip,
		Request: &http.Request{
			Header: http.Header{
				"User-Agent": {userAgent},
			},
		},
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

func TestTwoFailedOTP(t *testing.T) {
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
	id2 := utils.CreateID(0)

	buffer := &models.OnetimePasswordBuffer{
		Id:      id,
		IsLogin: true,

		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 10,
		},

		FailedCount: 3, // 0, 1, 2 :すでに計3回間違っている

		UserId: models.UserId{
			UserId: dummy.UserID,
		},
	}

	err = buffer.Add(ctx, db)
	require.NoError(t, err)

	buffer2 := &models.OnetimePasswordBuffer{
		Id:      id2,
		IsLogin: true,

		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 10,
		},

		FailedCount: 0,

		UserId: models.UserId{
			UserId: dummy.UserID,
		},
	}

	err = buffer2.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetOTPBufferByID(ctx, db, id)
		require.NoError(t, err)

		entity2, err := models.GetOTPBufferByID(ctx, db, id2)
		require.NoError(t, err)

		return entity != nil && entity2 != nil
	}, "")

	ip := "192.0.2.0"
	userAgent := "Mozilla/5.0 (platform; rv:geckoversion) Gecko/geckotrail Firefox/firefoxversion"

	c := &common.Cert{
		Ip: ip,
		Request: &http.Request{
			Header: http.Header{
				"User-Agent": {userAgent},
			},
		},
	}

	passcode, err := dummy.GenOTPCode()
	require.NoError(t, err)

	// case.1 4回以上失敗するとそれ以上は認証不可

	// 4回目失敗する
	err = login.LoginOTP(ctx, id, "123456", c)
	require.Error(t, err)

	// 4回以上するとそれ以降に成功してもBufferが消えているので認証が通らない (5回目)
	err = login.LoginOTP(ctx, id, passcode, c)
	require.Error(t, err)

	// case.2 失敗しても4回以内で成功すれば認証が通る

	// 1回失敗する
	err = login.LoginOTP(ctx, id2, "123456", c)
	require.Error(t, err)

	// 2回目で成功すれば認証は通る
	err = login.LoginOTP(ctx, id2, passcode, c)
	require.NoError(t, err)

}
