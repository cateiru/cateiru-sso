package handler_test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/core/oauth"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/handler"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/tests/tools"
	"github.com/cateiru/cateiru-sso/api/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func oauthServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/perview", handler.OAuthPreview)
	mux.HandleFunc("/login", handler.OAuthLogin)

	return mux
}

func TestPerview(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	clientId := utils.CreateID(30)

	service := models.SSOService{
		ClientID:    clientId,
		TokenSecret: utils.CreateID(0),

		Name:        "test",
		ServiceIcon: "image",

		FromUrl: []string{"https://example.com"},
		ToUrl:   []string{"https://example.com/login"},

		UserId: models.UserId{
			UserId: utils.CreateID(0),
		},
	}
	err = service.Add(ctx, db)
	require.NoError(t, err)

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, oauthServer(), true)
	s.AddSession(ctx, db, dummy)

	req := oauth.Service{
		Scope:        []string{"openid"},
		ResponseType: "code",
		RedirectURL:  "https://example.com/login",
		ClientID:     clientId,
		State:        utils.CreateID(0),
		Prompt:       "consent",
		FromURL:      "https://example.com",
	}

	resp := s.Post(t, "/perview", req)

	var respBody oauth.ResponsePerview
	err = json.Unmarshal(tools.ConvertByteResp(resp), &respBody)
	require.NoError(t, err)

	require.Equal(t, respBody.Name, "test")
	require.Equal(t, respBody.ServiceIcon, "image")
}

func TestOauthLogin(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	clientId := utils.CreateID(30)

	service := models.SSOService{
		ClientID:    clientId,
		TokenSecret: utils.CreateID(0),

		Name:        "test",
		ServiceIcon: "image",

		FromUrl: []string{"https://example.com"},
		ToUrl:   []string{"https://example.com/login"},

		UserId: models.UserId{
			UserId: utils.CreateID(0),
		},
	}
	err = service.Add(ctx, db)
	require.NoError(t, err)

	dummy := tools.NewDummyUser()
	_, err = dummy.AddUserInfo(ctx, db)
	require.NoError(t, err)

	s := tools.NewTestServer(t, oauthServer(), true)
	s.AddSession(ctx, db, dummy)

	req := oauth.Service{
		Scope:        []string{"openid"},
		ResponseType: "code",
		RedirectURL:  "https://example.com/login",
		ClientID:     clientId,
		State:        utils.CreateID(0),
		Prompt:       "consent",
		FromURL:      "https://example.com",
	}

	resp := s.Post(t, "/login", req)

	var respBody oauth.LoginResponse
	err = json.Unmarshal(tools.ConvertByteResp(resp), &respBody)
	require.NoError(t, err)

	require.NotEmpty(t, respBody.AccessToken)

	goretry.Retry(t, func() bool {
		entity, err := models.GetAccessTokenByAccessToken(ctx, db, respBody.AccessToken)
		require.NoError(t, err)

		return entity != nil && entity.ClientID == clientId && entity.UserId.UserId == dummy.UserID
	}, "")
}
