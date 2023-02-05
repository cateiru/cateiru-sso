package models_test

import (
	"context"
	"testing"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/src/config"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestRefreshToken(t *testing.T) {
	config.TestInit(t)

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

	err = models.DeleteRefreshToken(ctx, db, refreshToken)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetRefreshToken(ctx, db, refreshToken)
		require.NoError(t, err)

		return entity == nil
	}, "削除されている")
}

func TestRefreshTokenTX(t *testing.T) {
	config.TestInit(t)

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

	var entity *models.RefreshInfo

	/////
	for i := 0; 3 > i; i++ {
		tx, err := database.NewTransaction(ctx, db)
		require.NoError(t, err)

		entity, err = models.GetRefreshTokenTX(tx, refreshToken)
		require.NoError(t, err)

		require.Equal(t, entity.RefreshToken, refreshToken)
		entity.SessionToken = utils.CreateID(30)

		err = entity.AddTX(tx)
		require.NoError(t, err)

		err = models.DeleteRefreshTokenTX(tx, refreshToken)
		require.NoError(t, err)

		err = tx.Commit()
		if err != nil && err != datastore.ErrConcurrentTransaction {
			t.Fatal()
		}
		if err == nil {
			return
		}
	}
	/////

	goretry.Retry(t, func() bool {
		entity, err = models.GetRefreshTokenBySessionToken(ctx, db, sessionToken)
		require.NoError(t, err)

		return entity == nil
	}, "削除されている")
}

func TestDeleteRefresh(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	userId := utils.CreateID(30)
	sessionToken := utils.CreateID(0)
	refreshToken := utils.CreateID(0)

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
		entity, err := models.GetRefreshToken(ctx, db, refreshToken)
		require.NoError(t, err)

		return entity != nil
	}, "")

	// ---

	err = models.DeleteRefreshByUserId(ctx, db, userId)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetRefreshToken(ctx, db, refreshToken)
		require.NoError(t, err)

		return entity == nil
	}, "")
}

func TestDeleteRefreshPeriod(t *testing.T) {
	config.TestInit(t)

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
			CreateDate: time.Now().Add(time.Duration(-10) * time.Hour),
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

	err = models.DeleteRefreshTokenPeriod(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entry, err := models.GetRefreshToken(ctx, db, refreshToken)
		require.NoError(t, err)

		return entry == nil
	}, "entryがない")
}
