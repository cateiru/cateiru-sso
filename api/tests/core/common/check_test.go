package common_test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/cateiru/cateiru-sso/api/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

// 既にメールアドレスが存在する
func TestExistMail(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	_, err = dummy.AddUserCert(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		isExist, err := common.CheckExistMail(ctx, db, dummy.Mail)
		require.NoError(t, err)

		return isExist
	}, "同じメールアドレスが存在する")
}

// メールアドレスは存在しない
func TestNotExistMail(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	_, err = dummy.AddUserCert(ctx, db)
	require.NoError(t, err)

	newMail := fmt.Sprintf("%s@example.com", utils.CreateID(4))

	goretry.Retry(t, func() bool {
		isExist, err := common.CheckExistMail(ctx, db, newMail)
		require.NoError(t, err)

		return !isExist
	}, "同じメールアドレスは存在しない")
}

// adminのメールアドレスか
func TestIsADMIN(t *testing.T) {
	mail := "example@example.com"
	t.Setenv("ADMIN_MAIL", mail)

	require.True(t, common.CheckAdminMail(mail))
	require.False(t, common.CheckAdminMail("hoge@example.com"))
}

func TestBlockList(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	// 実行毎にランダムなメールアドレスとIPを作成
	mail := fmt.Sprintf("%s@example.com", utils.CreateID(4))
	ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))

	// ブロックリストを追加
	mailBlock := &models.MailBlockList{
		Mail: mail,
	}
	ipBlock := &models.IPBlockList{
		IP: ip,
	}
	err = mailBlock.Add(ctx, db)
	require.NoError(t, err)
	err = ipBlock.Add(ctx, db)
	require.NoError(t, err)

	// 初回のみ。Retryが成功した場合DB内に格納済みであるためその下はリトライは適用しない
	goretry.Retry(t, func() bool {
		isBlocked, err := common.ChaeckBlock(ctx, db, ip, "example@example.com")
		require.NoError(t, err)

		return isBlocked
	}, "IPでブロック")

	isBlocked, err := common.ChaeckBlock(ctx, db, "256.256.256.256", mail)
	require.NoError(t, err)
	require.True(t, isBlocked, "メールアドレスでブロック")

	isBlocked, err = common.ChaeckBlock(ctx, db, ip, mail)
	require.NoError(t, err)
	require.True(t, isBlocked, "メールアドレスとでブロック")

	isBlocked, err = common.ChaeckBlock(ctx, db, "256.256.256.256", "example@example.com")
	require.NoError(t, err)
	require.False(t, isBlocked, "ブロックリストにない")

}

func TestCheckExpired(t *testing.T) {
	now := time.Now()
	period := &models.Period{
		// 1時間前の時間
		CreateDate:   now.Add(time.Duration(-1) * time.Hour),
		PeriodMinute: 1,
	}
	require.True(t, common.CheckExpired(period), "分で正しく有効期限切れになっている")

	periodSafe := &models.Period{
		CreateDate:   now,
		PeriodMinute: 10,
	}
	require.False(t, common.CheckExpired(periodSafe))

	periodHour := &models.Period{
		// 2時間前の時間
		CreateDate: now.Add(time.Duration(-2) * time.Hour),
		PeriodHour: 1,
	}
	require.True(t, common.CheckExpired(periodHour), "時間で正しく有効期限切れになっている")

	periodDay := &models.Period{
		// 72時間前の時間
		CreateDate: now.Add(time.Duration(-72) * time.Hour),
		PeriodDay:  1,
	}
	require.True(t, common.CheckExpired(periodDay), "時間で正しく有効期限切れになっている")
}

func TestCheckOTP(t *testing.T) {
	t.Setenv("ISSUER", "test_issuer")

	dummy, err := tools.NewDummyUser().NewOTP()
	require.NoError(t, err)
	passcode, err := dummy.GenOTPCode()
	require.NoError(t, err)
	secret := dummy.Otp.GetSecret()

	result, _ := common.CheckOTP(passcode, nil, &secret)
	require.True(t, result)

	failedPass := "239432"

	result2, _ := common.CheckOTP(failedPass, nil, &secret)
	require.False(t, result2)
}

func TestCheckOTPBackups(t *testing.T) {
	passcode := "hogehoge"

	cert := models.Certification{
		OnetimePasswordSecret:  "ghofaw",
		OnetimePasswordBackups: []string{passcode},
	}

	result, check := common.CheckOTP(passcode, &cert, nil)
	require.True(t, check)
	require.True(t, result)

	failedPass := "239432"

	result2, check := common.CheckOTP(failedPass, &cert, nil)
	require.False(t, check)
	require.False(t, result2)
}
