package src_test

import (
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

func TestPasskeyEndpointHandler(t *testing.T) {
	h := NewTestHandler(t)

	t.Run("取得可能", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()

		err = h.PasskeyEndpointHandler(c)
		require.NoError(t, err)

		response := src.PasskeyEndpointResponse{}
		require.NoError(t, m.Json(&response))

		snaps.MatchSnapshot(t, response)
	})
}

func TestChangePasswordHandler(t *testing.T) {
	h := NewTestHandler(t)

	t.Run("リダイレクト可能", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()

		err = h.ChangePasswordHandler(c)
		require.NoError(t, err)

		require.Equal(t, http.StatusFound, m.W.Code)
		require.Equal(t, m.W.Header().Get("Location"), "http://cateiru.test/forget_password")
	})
}
