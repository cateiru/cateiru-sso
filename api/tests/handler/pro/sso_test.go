package pro_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru-sso/api/core/pro"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/handler"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/stretchr/testify/require"
)

func ssoServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.ProSSOHandler)

	return mux
}

func TestSSO(t *testing.T) {
	t.Setenv("DATASTORE_EMULATOR_HOST", "localhost:18001")
	t.Setenv("DATASTORE_PROJECT_ID", "project-test")

	t.Setenv("ISSUER", "TestIssuer")

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	dummy := tools.NewDummyUser().AddRole("pro")
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, ssoServer(), true)
	s.AddSession(ctx, db, dummy)

	// --- SSOを追加する

	form := pro.AddRequestForm{
		Name:      "Test",
		FromURL:   []string{"https://example.com/login"},
		ToURL:     []string{"https://example.com/login/redirect"},
		LoginOnly: true,

		SessionTokenPeriod: 10,
		RefreshTokenPeriod: 30,
	}

	resp := s.Post(t, "/", form)

	var keys pro.AddResponse
	err = json.Unmarshal(tools.ConvertByteResp(resp), &keys)
	require.NoError(t, err)

	require.NotEmpty(t, keys.PrivateKey)
	require.NotEmpty(t, keys.PublicKey)
	require.NotEmpty(t, keys.SecretKey)

	// --- SSO一覧を取得する

	resp = s.Get(t, "/")

	var sso []models.SSOService
	err = json.Unmarshal(tools.ConvertByteResp(resp), &sso)
	require.NoError(t, err)

	require.Len(t, sso, 1)
	require.Equal(t, sso[0].UserId.UserId, dummy.UserID)

	// --- SSOを削除する

	s.Delete(t, fmt.Sprintf("/?key=%s", keys.PublicKey))

	// --- もう一度一覧を取得する（削除されたか確認する）

	resp, err = s.Client.Get(s.Server.URL + "/")
	require.NoError(t, err)
	require.Equal(t, resp.StatusCode, 400)
}
