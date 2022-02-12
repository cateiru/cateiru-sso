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

func TestAccessToken(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	entity := models.SSOAccessToken{
		SSOAccessToken: utils.CreateID(0),
		ClientID:       utils.CreateID(0),

		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 5,
		},
		UserId: models.UserId{
			UserId: dummy.UserID,
		},
	}

	err = entity.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		getE, err := models.GetAccessTokenByAccessToken(ctx, db, entity.SSOAccessToken)
		require.NoError(t, err)

		return getE != nil && getE.UserId.UserId == dummy.UserID
	}, "")

	// 削除
	err = models.DeleteAccessTokenByClientID(ctx, db, entity.ClientID)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		getE, err := models.GetAccessTokenByAccessToken(ctx, db, entity.SSOAccessToken)
		require.NoError(t, err)

		return getE == nil
	}, "")
}

func TestDeleteSSOAccessTokenByUserId(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	entity := models.SSOAccessToken{
		SSOAccessToken: utils.CreateID(0),
		ClientID:       utils.CreateID(0),

		RedirectURI: "https://example.com",

		Period: models.Period{
			CreateDate:   time.Now(),
			PeriodMinute: 5,
		},
		UserId: models.UserId{
			UserId: dummy.UserID,
		},
	}

	err = entity.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		getE, err := models.GetAccessTokenByAccessToken(ctx, db, entity.SSOAccessToken)
		require.NoError(t, err)

		return getE != nil
	}, "")

	err = models.DeleteAccessTokenByUserId(ctx, db, dummy.UserID)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		getE, err := models.GetAccessTokenByAccessToken(ctx, db, entity.SSOAccessToken)
		require.NoError(t, err)

		return getE == nil
	}, "削除されている")
}
