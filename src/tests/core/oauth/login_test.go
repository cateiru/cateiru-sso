package oauth_test

import (
	"context"
	"testing"

	"github.com/cateiru/cateiru-sso/api/config"
	"github.com/cateiru/cateiru-sso/api/core/oauth"
	"github.com/cateiru/cateiru-sso/api/database"
	"github.com/cateiru/cateiru-sso/api/models"
	"github.com/cateiru/cateiru-sso/api/utils"
	goretry "github.com/cateiru/go-retry"
	"github.com/stretchr/testify/require"
)

func TestOAuthLogin(t *testing.T) {
	config.TestInit(t)

	ctx := context.Background()

	db, err := database.NewDatabase(ctx)
	require.NoError(t, err)
	defer db.Close()

	clientId := utils.CreateID(30)

	service := models.SSOService{
		ClientID: clientId,

		TokenSecret: utils.CreateID(0),

		Name:        "test",
		ServiceIcon: "",

		FromUrl: []string{"https://example.com"},
		ToUrl:   []string{"https://example.com/login"},

		UserId: models.UserId{
			UserId: "",
		},
	}
	err = service.Add(ctx, db)
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetSSOServiceByClientId(ctx, db, clientId)
		require.NoError(t, err)

		return entity != nil
	}, "")

	accessToken, err := oauth.LoginOAuth(ctx, db, clientId, "", "")
	require.NoError(t, err)

	goretry.Retry(t, func() bool {
		entity, err := models.GetAccessTokenByAccessToken(ctx, db, accessToken)
		require.NoError(t, err)

		return entity != nil && entity.ClientID == clientId
	}, "")
}
