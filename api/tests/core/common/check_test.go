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
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/stretchr/testify/require"
)

// 既にメールアドレスが存在する
func TestExistMail(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	// 実行毎にランダムなメールアドレスを作成
	mail := fmt.Sprintf("%s@example.com", utils.CreateID(4))

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	certification := models.Certification{
		AccountCreateDate: time.Now(),

		OnetimePasswordSecret:  "test",
		OnetimePasswordBackups: []string{"test1", "test2"},

		UserMailPW: models.UserMailPW{
			Mail:     mail,
			Password: "test",
		},
		UserId: models.UserId{
			UserId: utils.CreateID(0),
		},
	}

	// 要素追加
	err = certification.Add(ctx, db)
	require.NoError(t, err)

	isExist, err := common.CheckExistMail(ctx, db, mail)
	require.NoError(t, err)

	require.True(t, isExist, "同じメールアドレスが存在する")
}

// メールアドレスは存在しない
func TestNotExistMail(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	// 実行毎にランダムなメールアドレスを作成
	mail := fmt.Sprintf("%s@example.com", utils.CreateID(4))

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	certification := models.Certification{
		AccountCreateDate: time.Now(),

		OnetimePasswordSecret:  "test",
		OnetimePasswordBackups: []string{"test1", "test2"},

		UserMailPW: models.UserMailPW{
			Mail:     mail,
			Password: "test",
		},
		UserId: models.UserId{
			UserId: utils.CreateID(0),
		},
	}

	// 要素追加
	err = certification.Add(ctx, db)
	require.NoError(t, err)

	newMail := "example@example.com"

	isExist, err := common.CheckExistMail(ctx, db, newMail)
	require.NoError(t, err)

	require.False(t, isExist, "同じメールアドレスは存在しない")
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

	isBlocked, err := common.ChaeckBlock(ctx, db, ip, "example@example.com")
	require.NoError(t, err)
	require.True(t, isBlocked, "IPでブロック")

	isBlocked, err = common.ChaeckBlock(ctx, db, "256.256.256.256", mail)
	require.NoError(t, err)
	require.True(t, isBlocked, "メールアドレスでブロック")

	isBlocked, err = common.ChaeckBlock(ctx, db, ip, mail)
	require.NoError(t, err)
	require.True(t, isBlocked, "メールアドレスとでブロック")

	isBlocked, err = common.ChaeckBlock(ctx, db, "256.256.256.256", "example@example.com")
	require.NoError(t, err)
	require.False(t, isBlocked, "ブロックリストにない")

}
