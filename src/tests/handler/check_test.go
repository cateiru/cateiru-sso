package handler_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/core/check"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/handler"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/stretchr/testify/require"
)

func checkServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/check", handler.CheckHandler)

	return mux
}

func TestCheckUserName(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()
	info, err := dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, checkServer(), true)

	resp := s.Get(t, fmt.Sprintf("/check?name=%s", info.UserName))

	var existUserName check.ResponseCheckUserName
	err = json.Unmarshal(tools.ConvertByteResp(resp), &existUserName)
	require.NoError(t, err)

	require.True(t, existUserName.Exist)

	resp = s.Get(t, "/check?name=hello")

	err = json.Unmarshal(tools.ConvertByteResp(resp), &existUserName)
	require.NoError(t, err)

	require.False(t, existUserName.Exist)
}
