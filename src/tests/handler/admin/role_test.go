package admin_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru-sso/src/config"
	"github.com/cateiru/cateiru-sso/src/core/admin"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/handler"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/tests/tools"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func adminRoleServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/role", handler.AdminRoleHandler)

	return mux
}

func find(s []string, t string) bool {
	for _, v := range s {
		if t == v {
			return true
		}
	}

	return false
}

func TestAdminRole(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	roleDummy := tools.NewDummyUser()
	_, err = roleDummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	dummy := tools.NewDummyUser().AddRole("admin")
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetUserDataByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return entity != nil
	}, "")

	s := tools.NewTestServer(t, adminRoleServer(), true)
	defer s.Close()

	err = s.AddSession(ctx, db, dummy)
	require.NoError(t, err)

	form := admin.RoleRequest{
		Action: "enable",
		Role:   "test",
		UserId: roleDummy.UserID,
	}

	s.Post(t, "/role", form)

	goretry.Retry(t, func() bool {
		user, err := models.GetUserDataByUserID(ctx, db, roleDummy.UserID)
		require.NoError(t, err)

		role, err := models.GetRoleByUserID(ctx, db, roleDummy.UserID)
		require.NoError(t, err)

		return user != nil && role != nil && find(user.Role, "test") && find(role.Role, "test")
	}, "roleが追加されている")

	form = admin.RoleRequest{
		Action: "disable",
		Role:   "test",
		UserId: roleDummy.UserID,
	}

	s.Post(t, "/role", form)

	goretry.Retry(t, func() bool {
		user, err := models.GetUserDataByUserID(ctx, db, roleDummy.UserID)
		require.NoError(t, err)

		role, err := models.GetRoleByUserID(ctx, db, roleDummy.UserID)
		require.NoError(t, err)

		return user != nil && role != nil && !find(user.Role, "test") && !find(role.Role, "test")
	}, "roleが削除されている")
}

func TestRoleSelf(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser().AddRole("admin")
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetUserDataByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return entity != nil
	}, "")

	s := tools.NewTestServer(t, adminRoleServer(), true)
	defer s.Close()

	err = s.AddSession(ctx, db, dummy)
	require.NoError(t, err)

	form := admin.RoleRequest{
		Action: "enable",
		Role:   "test",
		UserId: dummy.UserID,
	}

	reqForm, err := json.Marshal(form)
	require.NoError(t, err)

	resp, err := s.Client.Post(s.Server.URL+"/role", "application/json", bytes.NewBuffer(reqForm))
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 400)

}
