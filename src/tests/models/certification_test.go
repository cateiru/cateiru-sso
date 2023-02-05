package models_test

import (
	"context"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/src/config"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/tests/tools"
	"github.com/cateiru/cateiru-sso/src/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestCertification(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	mail := "example@example.com"
	userId := utils.CreateID(30)

	entry := &models.Certification{
		AccountCreateDate:      time.Now(),
		OnetimePasswordSecret:  "",
		OnetimePasswordBackups: nil,
		UserMailPW: models.UserMailPW{
			Mail:     mail,
			Password: []byte("password"),
			Salt:     []byte(""),
		},
		UserId: models.UserId{
			UserId: userId,
		},
	}
	err = entry.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		result, err := models.GetCertificationByMail(ctx, db, mail)
		require.NoError(t, err)

		return result != nil
	}, "entryがある")

	// メールアドレスで探索
	result, err := models.GetCertificationByMail(ctx, db, mail)
	require.NoError(t, err)
	require.NotNil(t, result, "メールアドレスで探して要素がある")
	require.Equal(t, result.Password, entry.Password, "パスワードが同じ")

	result, err = models.GetCertificationByMail(ctx, db, "hoge@example.com")
	require.NoError(t, err)
	require.Nil(t, result, "メールアドレスで探したけど要素がなかった")

	// user idで探索
	result, err = models.GetCertificationByUserID(ctx, db, userId)
	require.NoError(t, err)
	require.NotNil(t, result, "user idで探して要素がある")
	require.Equal(t, result.Password, entry.Password, "パスワードが同じ")

	result, err = models.GetCertificationByUserID(ctx, db, "foo")
	require.NoError(t, err)
	require.Nil(t, result, "user idで探したけど要素がなかった")

}

func TestDeleteCert(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	_, err = dummy.AddUserCert(ctx, db)
	require.NoError(t, err)

	// 実際に格納されているか確認する
	goretry.Retry(t, func() bool {
		entity, err := models.GetCertificationByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return entity != nil
	}, "格納された")

	// 削除する
	err = models.DeleteCertificationByUserId(ctx, db, dummy.UserID)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetCertificationByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return entity == nil
	}, "削除された")
}
