package models_test

import (
	"context"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/src/config"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/tests/tools"
	"github.com/cateiru/cateiru-sso/src/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestSSORefresh(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	entity := models.SSORefreshToken{
		SSOAccessToken:  utils.CreateID(0),
		SSORefreshToken: utils.CreateID(0),

		ClientID: utils.CreateID(0),

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
		getE, err := models.GetSSORefreshTokenByRefreshToken(ctx, db, entity.SSORefreshToken)
		require.NoError(t, err)

		return getE != nil && getE.UserId.UserId == dummy.UserID
	}, "")

	// 削除
	err = models.DeleteSSORefreshTokenByClientId(ctx, db, entity.ClientID)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		getE, err := models.GetSSORefreshTokenByRefreshToken(ctx, db, entity.SSOAccessToken)
		require.NoError(t, err)

		return getE == nil
	}, "")
}

func TestDeleteSSORefreshByUserId(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	entity := models.SSORefreshToken{
		SSOAccessToken:  utils.CreateID(0),
		SSORefreshToken: utils.CreateID(0),

		ClientID: utils.CreateID(0),

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
		getE, err := models.GetSSORefreshTokenByRefreshToken(ctx, db, entity.SSORefreshToken)
		require.NoError(t, err)

		return getE != nil
	}, "")

	err = models.DeleteSSORefreshTokenByUserId(ctx, db, dummy.UserID)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		getE, err := models.GetSSORefreshTokenByRefreshToken(ctx, db, entity.SSORefreshToken)
		require.NoError(t, err)

		return getE == nil
	}, "削除されている")
}

func TestDeleteSSORefreshByRefreshToken(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	entity := models.SSORefreshToken{
		SSOAccessToken:  utils.CreateID(0),
		SSORefreshToken: utils.CreateID(0),

		ClientID: utils.CreateID(0),

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
		getE, err := models.GetSSORefreshTokenByRefreshToken(ctx, db, entity.SSORefreshToken)
		require.NoError(t, err)

		return getE != nil
	}, "")

	err = models.DeleteSSORefreshTokenByRefreshToken(ctx, db, entity.SSORefreshToken)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		getE, err := models.GetSSORefreshTokenByRefreshToken(ctx, db, entity.SSORefreshToken)
		require.NoError(t, err)

		return getE == nil
	}, "削除されている")
}

func TestDeleteSSORefreshTokenByClientAndUserID(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	entity := models.SSORefreshToken{
		SSOAccessToken:  utils.CreateID(0),
		SSORefreshToken: utils.CreateID(0),

		ClientID: utils.CreateID(0),

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
		getE, err := models.GetSSORefreshTokenByRefreshToken(ctx, db, entity.SSORefreshToken)
		require.NoError(t, err)

		return getE != nil
	}, "")

	err = models.DeleteSSORefreshTokenByUserIdAndClientID(ctx, db, dummy.UserID, entity.ClientID)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		getE, err := models.GetSSORefreshTokenByRefreshToken(ctx, db, entity.SSORefreshToken)
		require.NoError(t, err)

		return getE == nil
	}, "削除されている")
}

func TestDeleteSSORefreshTokenPeriod(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	entity := models.SSORefreshToken{
		SSOAccessToken:  utils.CreateID(0),
		SSORefreshToken: utils.CreateID(0),

		ClientID: utils.CreateID(0),

		Period: models.Period{
			CreateDate:   time.Now().Add(time.Duration(-1) * time.Hour),
			PeriodMinute: 5,
		},
		UserId: models.UserId{
			UserId: dummy.UserID,
		},
	}

	err = entity.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		getE, err := models.GetSSORefreshTokenByRefreshToken(ctx, db, entity.SSORefreshToken)
		require.NoError(t, err)

		return getE != nil && getE.UserId.UserId == dummy.UserID
	}, "")

	err = models.DeleteSSORefreshTokenPeriod(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		getE, err := models.GetSSORefreshTokenByRefreshToken(ctx, db, entity.SSORefreshToken)
		require.NoError(t, err)

		return getE == nil
	}, "")
}
