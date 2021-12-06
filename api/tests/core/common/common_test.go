package common_test

import (
	"context"
	"fmt"
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
