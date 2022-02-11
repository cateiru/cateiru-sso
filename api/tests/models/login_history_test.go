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

func TestLoginHistory(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	// 10個入れる
	for i := 0; 10 > i; i++ {
		history := &models.LoginHistory{
			AccessId:  utils.CreateID(20),
			Date:      time.Now(),
			IpAddress: "192.168.0.1",
			UserAgent: "",

			UserId: models.UserId{
				UserId: dummy.UserID,
			},
		}

		err = history.Add(ctx, db)
		require.NoError(t, err)
	}

	goretry.Retry(t, func() bool {
		histores, err := models.GetAllLoginHistory(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return len(histores) == 10 && histores[0].IpAddress == "192.168.0.1"
	}, "Entityが10個ある")

	// ---- limitを指定して取得する

	goretry.Retry(t, func() bool {
		histores, err := models.GetAllLoginHistory(ctx, db, dummy.UserID, 3)
		require.NoError(t, err)

		return len(histores) == 3 && histores[0].IpAddress == "192.168.0.1"
	}, "Entityが3つだけ取得できる")

	// ---- 削除する

	err = models.DeleteAllLoginHistories(ctx, db, dummy.UserID)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		histores, err := models.GetAllLoginHistory(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return len(histores) == 0
	}, "Entityが全部削除されている")
}
