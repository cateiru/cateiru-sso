package models_test

import (
	"context"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestRefreshToken(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	userId := utils.CreateID(30)
	sessionToken := utils.CreateID(30)
	refreshToken := utils.CreateID(30)

	session := &models.RefreshInfo{
		RefreshToken: refreshToken,
		SessionToken: sessionToken,

		Period: models.Period{
			CreateDate: time.Now(),
			PeriodHour: 6,
		},

		UserId: models.UserId{
			UserId: userId,
		},
	}

	err = session.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entry, err := models.GetRefreshToken(ctx, db, refreshToken)
		require.NoError(t, err)

		return entry != nil
	}, "entryがある")

	entry, err := models.GetRefreshToken(ctx, db, refreshToken)
	require.NoError(t, err)
	require.NotNil(t, entry)
	require.Equal(t, entry.SessionToken, sessionToken)

	entries, err := models.GetRefreshTokenByUserId(ctx, db, userId)
	require.NoError(t, err)
	require.Equal(t, len(entries), 1)

	require.Equal(t, entries[0].SessionToken, sessionToken)

	entry, err = models.GetRefreshTokenBySessionToken(ctx, db, sessionToken)
	require.NoError(t, err)
	require.NotNil(t, entry)
	require.Equal(t, entry.SessionToken, sessionToken)
}

func TestRefreshTokenTX(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	userId := utils.CreateID(30)
	sessionToken := utils.CreateID(30)
	refreshToken := utils.CreateID(30)

	session := &models.RefreshInfo{
		RefreshToken: refreshToken,
		SessionToken: sessionToken,

		Period: models.Period{
			CreateDate: time.Now(),
			PeriodHour: 6,
		},

		UserId: models.UserId{
			UserId: userId,
		},
	}

	err = session.Add(ctx, db)
	require.NoError(t, err)

	/////

	tx, err := database.NewTransaction(ctx, db)
	require.NoError(t, err)

	entity, err := models.GetRefreshTokenTX(tx, refreshToken)
	require.NoError(t, err)

	require.Equal(t, entity.RefreshToken, refreshToken)
	entity.SessionToken = utils.CreateID(30)

	err = entity.AddTX(tx)
	require.NoError(t, err)

	err = models.DeleteRefreshTokenTX(tx, refreshToken)
	require.NoError(t, err)

	err = tx.Commit()
	require.NoError(t, err)
	/////

	goretry.Retry(t, func() bool {
		entity, err = models.GetRefreshTokenBySessionToken(ctx, db, sessionToken)
		require.NoError(t, err)

		return entity == nil
	}, "削除されている")
}
