package common_test

import (
	"context"
	"testing"

	"github.com/cateiru/cateiru-sso/api/core/common"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestProMore(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	// Proユーザにする
	dummy := tools.NewDummyUser().AddRole("pro")

	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetUserDataByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return entity != nil
	}, "entityが格納された")

	err = common.ProMoreOnly(ctx, db, dummy.UserID)
	require.NoError(t, err)
}

func TestNoPro(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetUserDataByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return entity != nil
	}, "entityが格納された")

	err = common.ProMoreOnly(ctx, db, dummy.UserID)
	require.Error(t, err)
}
