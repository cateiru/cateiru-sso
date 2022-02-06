package models_test

import (
	"context"
	"testing"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestSessionToken(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	userId := utils.CreateID(30)
	sessionToken := utils.CreateID(30)

	session := &models.SessionInfo{
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
		entry, err := models.GetSessionToken(ctx, db, sessionToken)
		require.NoError(t, err)

		return entry != nil
	}, "entryがある")

	entry, err := models.GetSessionToken(ctx, db, sessionToken)
	require.NoError(t, err)
	require.NotNil(t, entry)
	require.Equal(t, entry.SessionToken, sessionToken)

	entries, err := models.GetSessionTokenByUserId(ctx, db, userId)
	require.NoError(t, err)
	require.Equal(t, len(entries), 1)

	require.Equal(t, entries[0].SessionToken, sessionToken)

	err = models.DeleteSessionToken(ctx, db, sessionToken)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetSessionToken(ctx, db, sessionToken)
		require.NoError(t, err)

		return entity == nil
	}, "削除されている")
}

func TestSessionTX(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	userId := utils.CreateID(30)
	sessionToken := utils.CreateID(30)

	session := &models.SessionInfo{
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

	var entity *models.SessionInfo
	for i := 0; 3 > i; i++ {
		tx, err := database.NewTransaction(ctx, db)
		require.NoError(t, err)

		entity, err = models.GetSessionTokenTX(tx, sessionToken)
		require.NoError(t, err)

		require.Equal(t, entity.SessionToken, sessionToken)

		err = entity.AddTX(tx)
		require.NoError(t, err)

		err = models.DeleteSessionTokenTX(tx, sessionToken)
		require.NoError(t, err)

		err = tx.Commit()
		require.NoError(t, err)
		if err != nil && err != datastore.ErrConcurrentTransaction {
			t.Fatal()
		}
		if err == nil {
			return
		}

	}
	////

	entity, err = models.GetSessionToken(ctx, db, sessionToken)
	require.NoError(t, err)
	require.Nil(t, entity)
}

func TestDeleteSession(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	userId := utils.CreateID(30)
	sessionToken := utils.CreateID(30)

	session := &models.SessionInfo{
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

	// ---

	err = models.DeleteSessionByUserId(ctx, db, userId)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetSessionToken(ctx, db, sessionToken)
		require.NoError(t, err)

		return entity == nil
	}, "")
}
