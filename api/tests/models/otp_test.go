package models_test

import (
	"context"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestOTPBuffer(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	userId := utils.CreateID(30)
	id := utils.CreateID(0)

	optBuffer := &models.OnetimePasswordBuffer{
		Id: id,

		PublicKey: "public",
		SecretKey: "secret",

		IsLogin: false,

		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 10,
		},
		UserId: models.UserId{
			UserId: userId,
		},
	}

	err = optBuffer.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetOTPBufferByID(ctx, db, id)
		require.NoError(t, err)

		return entity != nil && entity.UserId.UserId == userId
	}, "要素が取得できる")

	err = models.DeleteOTPBuffer(ctx, db, id)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetOTPBufferByID(ctx, db, id)
		require.NoError(t, err)

		return entity == nil
	}, "削除されている")
}
