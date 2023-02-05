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

func TestSSOServiceLog(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	clientId := utils.CreateID(0)
	userIds := []string{
		utils.CreateID(0),
		utils.CreateID(0),
	}

	for _, userid := range userIds {
		entity := models.SSOServiceLog{
			LogId:      utils.CreateID(0),
			AcceptDate: time.Now(),
			ClientID:   clientId,

			UserId: models.UserId{
				UserId: userid,
			},
		}

		err = entity.Add(ctx, db)
		require.NoError(t, err)
	}

	goretry.Retry(t, func() bool {
		logs, err := models.GetSSOServiceLogsByClientId(ctx, db, clientId)
		require.NoError(t, err)

		return len(logs) == 2
	}, "")

	count, err := models.CountSSOServiceLogByClientId(ctx, db, clientId)
	require.NoError(t, err)

	require.Equal(t, count, 2)

	logs, err := models.GetSSOServiceLogsByUserId(ctx, db, userIds[0])
	require.NoError(t, err)

	require.Len(t, logs, 1)

	err = models.DeleteSSOServiceLogByClientId(ctx, db, clientId)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		logs, err := models.GetSSOServiceLogsByClientId(ctx, db, clientId)
		require.NoError(t, err)

		return len(logs) == 0
	}, "")
}

func TestDeleteServiceLogByUserId(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	clientId := utils.CreateID(0)
	dummy := tools.NewDummyUser()

	entity := models.SSOServiceLog{
		LogId:      utils.CreateID(0),
		AcceptDate: time.Now(),
		ClientID:   clientId,

		UserId: models.UserId{
			UserId: dummy.UserID,
		},
	}

	err = entity.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		logs, err := models.GetSSOServiceLogsByClientId(ctx, db, clientId)
		require.NoError(t, err)

		return len(logs) == 1
	}, "")

	err = models.DeleteSSOServiceLogByUserId(ctx, db, dummy.UserID)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		logs, err := models.GetSSOServiceLogsByClientId(ctx, db, clientId)
		require.NoError(t, err)

		return len(logs) == 0
	}, "削除されている")
}

func TestDeleteServiceLogByUserIDAndClientId(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	clientId := utils.CreateID(0)
	dummy := tools.NewDummyUser()

	entity := models.SSOServiceLog{
		LogId:      utils.CreateID(0),
		AcceptDate: time.Now(),
		ClientID:   clientId,

		UserId: models.UserId{
			UserId: dummy.UserID,
		},
	}

	err = entity.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		logs, err := models.GetSSOServiceLogsByClientId(ctx, db, clientId)
		require.NoError(t, err)

		return len(logs) == 1
	}, "")

	err = models.DeleteSSOServiceLogByUserIDAndClientId(ctx, db, dummy.UserID, clientId)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		logs, err := models.GetSSOServiceLogsByClientId(ctx, db, clientId)
		require.NoError(t, err)

		return len(logs) == 0
	}, "削除されている")
}
