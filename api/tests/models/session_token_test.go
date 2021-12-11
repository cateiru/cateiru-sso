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

func TestSessionToken(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	userId := utils.CreateID(30)
	sessionToken := utils.CreateID(30)

	session := &models.SessionInfo{
		SessionToken: sessionToken,

		TokenInfo: models.TokenInfo{
			CreateDate: time.Now(),
			PeriodHour: 6,

			UserId: models.UserId{
				UserId: userId,
			},
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
}
