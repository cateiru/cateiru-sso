package models_test

import (
	"context"
	"testing"

	"github.com/cateiru/cateiru-sso/src/config"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/tests/tools"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestRole(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	role := &models.Role{
		Role: []string{"user"},

		UserId: models.UserId{
			UserId: dummy.UserID,
		},
	}

	err = role.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetRoleByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return entity != nil && entity.Role[0] == "user"
	}, "roleが格納されて取得できる")

	err = models.DeleteRoleByUserID(ctx, db, dummy.UserID)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetRoleByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return entity == nil
	}, "削除できている")
}
