package history_test

import (
	"context"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/api/core/user/history"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/cateiru/cateiru-sso/api/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestLoginHistores(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	_, err = dummy.AddUserCert(ctx, db)
	require.NoError(t, err)

	// 10個入れる
	for i := 0; 10 > i; i++ {
		history := &models.LoginHistory{
			AccessId:     utils.CreateID(20),
			Date:         time.Now(),
			IpAddress:    "192.168.0.1",
			IsSSO:        false,
			SSOPublicKey: "",
			UserAgent:    "",

			UserId: models.UserId{
				UserId: dummy.UserID,
			},
		}

		err = history.Add(ctx, db)
		require.NoError(t, err)
	}

	goretry.Retry(t, func() bool {
		histories, err := history.LoginHistories(ctx, db, dummy.UserID, -1)
		if err != nil {
			t.Log(err)
			return false
		}

		return len(histories) == 10
	}, "")
}
