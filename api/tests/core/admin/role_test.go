package admin_test

import (
	"context"
	"testing"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/core/admin"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func find(s []string, t string) bool {
	for _, v := range s {
		if t == v {
			return true
		}
	}

	return false
}

func TestRole(t *testing.T) {
	config.TestInit(t)

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
	}, "roleが追加されている")

	err = admin.AddRole(ctx, db, "test", dummy.UserID)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		user, err := models.GetUserDataByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		role, err := models.GetRoleByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return user != nil && role != nil && find(user.Role, "test") && find(role.Role, "test")
	}, "")

	err = admin.DeleteRole(ctx, db, "test", dummy.UserID)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		user, err := models.GetUserDataByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		role, err := models.GetRoleByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return user != nil && role != nil && !find(user.Role, "test") && !find(role.Role, "test")
	}, "roleが削除されている")
}
