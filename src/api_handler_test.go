package src_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/require"
)

func TestApiOpenidConfigurationHandler(t *testing.T) {
	h := NewTestHandler(t)

	t.Run("取得可能", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()

		err = h.ApiOpenidConfigurationHandler(c)
		require.NoError(t, err)

		response := src.OpenidConfiguration{}
		require.NoError(t, m.Json(&response))

		snaps.MatchSnapshot(t, response)
	})
}

func TestJwksJsonHandler(t *testing.T) {
	h := NewTestHandler(t)

	t.Run("取得可能", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()

		err = h.JwksJsonHandler(c)
		require.NoError(t, err)

		data := m.W.Body.Bytes()

		snaps.MatchSnapshot(t, string(data))
	})
}

func TestWebIdentityHandler(t *testing.T) {
	h := NewTestHandler(t)

	t.Run("取得可能", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()

		err = h.WebIdentityHandler(c)
		require.NoError(t, err)

		response := src.WebIdentityResponse{}
		require.NoError(t, m.Json(&response))

		snaps.MatchSnapshot(t, response.ProvidersUrl)
	})
}

func TestFedCMConfigHandler(t *testing.T) {
	h := NewTestHandler(t)

	t.Run("取得可能", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()

		err = h.FedCMConfigHandler(c)
		require.NoError(t, err)

		response := src.FedCMConfigResponse{}
		require.NoError(t, m.Json(&response))

		require.NotEmpty(t, response.AccountsEndpoint)
		require.NotEmpty(t, response.ClientMetadataEndpoint)
		require.NotEmpty(t, response.IdAssertionEndpoint)

		require.Equal(t, response.Branding.BackgroundColor, C.BrandBackgroundColor)
		require.Equal(t, response.Branding.Color, C.BrandColor)
		require.Equal(t, response.Branding.Name, C.BrandName)
	})
}

func TestFedCMAccountsHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)
	t.Run("ユーザー数0", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		c := m.Echo()

		err = h.FedCMAccountsHandler(c)
		require.NoError(t, err)

		response := src.FedCMAccountsResponse{}
		require.NoError(t, m.Json(&response))

		require.Len(t, response.Accounts, 0)
	})

	t.Run("ユーザー数1", func(t *testing.T) {
		email1 := RandomEmail(t)

		u1 := RegisterUser(t, ctx, email1)

		cookies := RegisterSession(t, ctx, &u1)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.FedCMAccountsHandler(c)
		require.NoError(t, err)

		response := src.FedCMAccountsResponse{}
		require.NoError(t, m.Json(&response))

		require.Len(t, response.Accounts, 1)

		require.Equal(t, response.Accounts[0].ID, u1.ID)
		require.Equal(t, response.Accounts[0].Name, u1.UserName)
		require.Equal(t, response.Accounts[0].Email, u1.Email)
		require.Equal(t, response.Accounts[0].GivenName, u1.GivenName.String)
		require.Equal(t, response.Accounts[0].Picture, u1.Avatar.String)
	})

	t.Run("ユーザー数2", func(t *testing.T) {
		email1 := RandomEmail(t)
		email2 := RandomEmail(t)

		u1 := RegisterUser(t, ctx, email1)
		u2 := RegisterUser(t, ctx, email2)

		cookies := RegisterSession(t, ctx, &u1, &u2)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookies)
		c := m.Echo()

		err = h.FedCMAccountsHandler(c)
		require.NoError(t, err)

		response := src.FedCMAccountsResponse{}
		require.NoError(t, m.Json(&response))

		require.Len(t, response.Accounts, 2)
	})
}
