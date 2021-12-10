package models_test

import (
	"context"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	"github.com/stretchr/testify/require"
)

func TestCertification(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

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
			Password: "hoge",
		},
		UserId: models.UserId{
			UserId: userId,
		},
	}
	err = entry.Add(ctx, db)
	require.NoError(t, err)

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
