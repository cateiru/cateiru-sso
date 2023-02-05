package admin_test

import (
	"context"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/src/config"
	"github.com/cateiru/cateiru-sso/src/core/admin"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestWorker(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	// サンプルとして1つ格納する
	userId := utils.CreateID(30)
	sessionToken := utils.CreateID(30)

	session := &models.SessionInfo{
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
		entry, err := models.GetSessionToken(ctx, db, sessionToken)
		require.NoError(t, err)

		return entry != nil
	}, "entryがある")

	err = admin.Worker(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entry, err := models.GetSessionToken(ctx, db, sessionToken)
		require.NoError(t, err)

		return entry == nil
	}, "entryがない")
}
