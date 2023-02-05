package user_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru-sso/src/config"
	"github.com/cateiru/cateiru-sso/src/core/user/info"
	"github.com/cateiru/cateiru-sso/src/database"
	"github.com/cateiru/cateiru-sso/src/handler"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/cateiru-sso/src/tests/tools"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func infoServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/info", handler.UserInfoChangeHandler)

	return mux
}

func TestChangeInfo(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser()

	_, err = dummy.AddUserCert(ctx, db)
	require.NoError(t, err)
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetUserDataByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return entity != nil
	}, "")

	changed := info.Request{
		FirstName: "New",
		LastName:  "taro",
		UserName:  "aaaa",
		Theme:     "wwwww",
	}

	s := tools.NewTestServer(t, infoServer(), true)
	defer s.Close()
	err = s.AddSession(ctx, db, dummy)
	require.NoError(t, err)

	resp := s.Post(t, "/info", &changed)

	var respBody *info.Request

	err = json.Unmarshal(tools.ConvertByteResp(resp), &respBody)
	require.NoError(t, err)

	require.Equal(t, respBody.FirstName, changed.FirstName)

	goretry.Retry(t, func() bool {
		entity, err := models.GetUserDataByUserID(ctx, db, dummy.UserID)
		require.NoError(t, err)

		return entity != nil && entity.FirstName == changed.FirstName
	}, "")
}
