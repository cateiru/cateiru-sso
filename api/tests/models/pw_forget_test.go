package models_test

import (
	"context"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/cateiru/cateiru-sso/api/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestPWForget(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	token := utils.CreateID(30)
	dummy := tools.NewDummyUser()

	entity := models.PWForget{
		ForgetToken: token,
		Mail:        dummy.Mail,
		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 30,
		},
	}

	err = entity.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetPWForgetByToken(ctx, db, token)
		require.NoError(t, err)

		return entity != nil && entity.Mail == dummy.Mail
	}, "要素が格納されている")

	entities, err := models.GetPWForgetByMail(ctx, db, dummy.Mail)
	require.NoError(t, err)
	require.Len(t, entities, 1)

	err = models.DeletePWForgetByToken(ctx, db, token)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetPWForgetByToken(ctx, db, token)
		require.NoError(t, err)

		return entity == nil
	}, "削除されている")
}
