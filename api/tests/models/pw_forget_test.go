package models_test

import (
	"context"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/cateiru/cateiru-sso/api/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestPWForget(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	token := utils.CreateID(30)
	dummy := tools.NewDummyUser()

	entity := models.PWForget{
		ForgetToken: token,
		UserId: models.UserId{
			UserId: dummy.UserID,
		},
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

		return entity != nil && entity.UserId.UserId == dummy.UserID
	}, "要素が格納されている")
}
