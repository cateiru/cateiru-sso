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

func TestService(t *testing.T) {
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

	s := oauth.Service{
		Scope:        []string{"openid"},
		ResponseType: "code",
		RedirectURL:  "https://example.com/login",
		ClientID:     clientId,
		State:        utils.CreateID(0),
		Prompt:       "consent",
		FromURL:      "https://example.com",
	}

	_, err = s.Required(ctx, db)
	require.NoError(t, err)

	// 失敗

	s = oauth.Service{
		Scope:        []string{"hoge"}, // openid がない
		ResponseType: "code",
		RedirectURL:  "https://example.com/login",
		ClientID:     clientId,
		State:        utils.CreateID(0),
		Prompt:       "consent",
		FromURL:      "https://example.com",
	}

	_, err = s.Required(ctx, db)
	require.Error(t, err)

	s = oauth.Service{
		Scope:        []string{"openid"},
		ResponseType: "code",
		RedirectURL:  "https://example.com", // redirect urlが定義されてない
		ClientID:     clientId,
		State:        utils.CreateID(0),
		Prompt:       "consent",
		FromURL:      "https://example.com",
	}

	_, err = s.Required(ctx, db)
	require.Error(t, err)

	s = oauth.Service{
		Scope:        []string{"openid"},
		ResponseType: "code",
		RedirectURL:  "https://example.com/login",
		ClientID:     clientId,
		State:        utils.CreateID(0),
		Prompt:       "consent",
		FromURL:      "https://example.com/hoge", //from urlが定義されてない
	}

	_, err = s.Required(ctx, db)
	require.Error(t, err)

	s = oauth.Service{
		Scope:        []string{}, //空
		ResponseType: "code",
		RedirectURL:  "https://example.com/login",
		ClientID:     clientId,
		State:        utils.CreateID(0),
		Prompt:       "consent",
		FromURL:      "https://example.com",
	}

	_, err = s.Required(ctx, db)
	require.Error(t, err)

	s = oauth.Service{
		Scope:        []string{"openid"},
		ResponseType: "", // 空
		RedirectURL:  "https://example.com/login",
		ClientID:     clientId,
		State:        utils.CreateID(0),
		Prompt:       "consent",
		FromURL:      "https://example.com",
	}

	_, err = s.Required(ctx, db)
	require.Error(t, err)

	s = oauth.Service{
		Scope:        []string{"openid"},
		ResponseType: "code",
		RedirectURL:  "", // 空
		ClientID:     clientId,
		State:        utils.CreateID(0),
		Prompt:       "consent",
		FromURL:      "https://example.com",
	}

	_, err = s.Required(ctx, db)
	require.Error(t, err)

	s = oauth.Service{
		Scope:        []string{"openid"},
		ResponseType: "code",
		RedirectURL:  "https://example.com/login",
		ClientID:     clientId,
		State:        utils.CreateID(0),
		Prompt:       "consent",
		FromURL:      "", // 空
	}

	_, err = s.Required(ctx, db)
	require.Error(t, err)

	s = oauth.Service{
		Scope:        []string{"openid"},
		ResponseType: "code",
		RedirectURL:  "https://example.com/login",
		ClientID:     "", // 空
		State:        utils.CreateID(0),
		Prompt:       "consent",
		FromURL:      "https://example.com", // 空
	}

	_, err = s.Required(ctx, db)
	require.Error(t, err)
}
