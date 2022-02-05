package handler_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/handler"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func adminserver() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/user", handler.AdminUserHandler)

	return mux
}

func TestGetAllUser(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	for i := 0; i < 10; i++ {
		dummy := tools.NewDummyUser()
		_, err = dummy.AddUserInfo(ctx, db)
		require.NoError(t, err)
	}

	dummy := tools.NewDummyUser().AddRole("admin")
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetUserDataByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return entity != nil
	}, "")

	s := tools.NewTestServer(t, adminserver(), true)
	defer s.Close()

	err = s.AddSession(ctx, db, dummy)
	require.NoError(t, err)

	resp := s.Get(t, "/user")

	var respBody []models.User
	err = json.Unmarshal(tools.ConvertByteResp(resp), &respBody)
	require.NoError(t, err)

	require.True(t, len(respBody) != 0)
}
