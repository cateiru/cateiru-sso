package handler_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/handler"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/stretchr/testify/require"
)

func meServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.MeHandler)

	return mux
}

func TestMe(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, meServer(), true)
	s.AddSession(ctx, db, dummy)

	resp := s.Get(t, "/")

	var userInfo models.User
	err = json.Unmarshal(tools.ConvertByteResp(resp), &userInfo)
	require.NoError(t, err)

	require.Equal(t, userInfo.FirstName, "TestFirstName")
	require.Equal(t, userInfo.LastName, "TestLastName")
	require.Equal(t, userInfo.UserName, "TestUserName")
	require.Equal(t, userInfo.UserId.UserId, dummy.UserID)
	require.Equal(t, userInfo.Mail, dummy.Mail)
}

func TestMeNotVerify(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	s := tools.NewTestServer(t, meServer(), true)

	resp, err := s.Client.Get(s.Server.URL + "/")
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 403, "認証情報がないので何も返さない")
}
