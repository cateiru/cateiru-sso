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

func TestLoginHistory(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	userId := utils.CreateID(10)

	history := &models.LoginHistory{
		AccessId:     utils.CreateID(20),
		Date:         time.Now(),
		IpAddress:    "192.168.0.1",
		IsSSO:        false,
		SSOPublicKey: "",

		UserId: models.UserId{
			UserId: userId,
		},
	}

	err = history.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		histores, err := models.GetAllLoginHistory(ctx, db, userId)
		require.NoError(t, err)

		return histores[0].IpAddress == "192.168.0.1"
	}, "Entryがある")
}
