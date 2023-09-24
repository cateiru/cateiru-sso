package src_test

import (
	"testing"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/stretchr/testify/require"
)

func TestGetAuthenticationRequest(t *testing.T) {
	h := NewTestHandler(t)

	t.Run("成功", func(t *testing.T) {
		form := easy.NewMultipart()

		form.Insert("scope", "openid profile email")
		form.Insert("response_type", "code")
		form.Insert("client_id", "test")
		form.Insert("redirect_uri", "https://example.test")
		form.Insert("state", "state_test")
		form.Insert("response_mode", "query")
		form.Insert("nonce", "nonce_test")
		form.Insert("display", "page")
		form.Insert("prompt", "login consent")
		form.Insert("max_age", "3600")
		form.Insert("ui_locales", "ja_JP")
		form.Insert("id_token_hint", "id_token_hint_test")
		form.Insert("login_hint", "login_hint_test")
		form.Insert("acr_values", "acr_values_test")

		m, err := easy.NewFormData("/", "POST", form)
		require.NoError(t, err)

		c := m.Echo()

		authenticationRequest, err := h.GetAuthenticationRequest(c)
		require.NoError(t, err)

		require.Equal(t, authenticationRequest.Scopes, []string{"openid", "profile", "email"})
		require.Equal(t, authenticationRequest.ResponseType, lib.ResponseTypeAuthorizationCode)
		require.Equal(t, authenticationRequest.ClientId, "test")
		require.Equal(t, authenticationRequest.RedirectUri.String(), "https://example.test")
		require.True(t, authenticationRequest.State.Valid)
		require.Equal(t, authenticationRequest.State.String, "state_test")
		require.Equal(t, authenticationRequest.ResponseMode, lib.ResponseModeQuery)
		require.True(t, authenticationRequest.Nonce.Valid)
		require.Equal(t, authenticationRequest.Nonce.String, "nonce_test")
		require.Equal(t, authenticationRequest.Display, lib.DisplayPage)
		require.Equal(t, authenticationRequest.Prompts, []lib.Prompt{
			lib.PromptLogin,
			lib.PromptConsent,
		})
		require.Equal(t, authenticationRequest.MaxAge, uint64(3600))
		require.Equal(t, authenticationRequest.UiLocales, []string{"ja_JP"})
		require.True(t, authenticationRequest.IdTokenHint.Valid)
		require.Equal(t, authenticationRequest.IdTokenHint.String, "id_token_hint_test")
		require.True(t, authenticationRequest.LoginHint.Valid)
		require.Equal(t, authenticationRequest.LoginHint.String, "login_hint_test")
		require.True(t, authenticationRequest.AcrValues.Valid)
		require.Equal(t, authenticationRequest.AcrValues.String, "acr_values_test")
	})

	t.Run("失敗: scopeが存在しない", func(t *testing.T) {
		form := easy.NewMultipart()

		form.Insert("response_type", "code")
		form.Insert("client_id", "test")
		form.Insert("redirect_uri", "https://example.test")
		form.Insert("state", "state_test")
		form.Insert("response_mode", "query")
		form.Insert("nonce", "nonce_test")
		form.Insert("display", "page")
		form.Insert("prompt", "login consent")
		form.Insert("max_age", "3600")
		form.Insert("ui_locales", "ja_JP")
		form.Insert("id_token_hint", "id_token_hint_test")
		form.Insert("login_hint", "login_hint_test")
		form.Insert("acr_values", "acr_values_test")

		m, err := easy.NewFormData("/", "POST", form)
		require.NoError(t, err)

		c := m.Echo()

		_, err = h.GetAuthenticationRequest(c)
		require.EqualError(t, err, "code=400, error=invalid_request_uri, message=scope is required")
	})

	t.Run("失敗: scopeにopenidが存在しない", func(t *testing.T) {
		form := easy.NewMultipart()

		form.Insert("scope", "profile email")
		form.Insert("response_type", "code")
		form.Insert("client_id", "test")
		form.Insert("redirect_uri", "https://example.test")
		form.Insert("state", "state_test")
		form.Insert("response_mode", "query")
		form.Insert("nonce", "nonce_test")
		form.Insert("display", "page")
		form.Insert("prompt", "login consent")
		form.Insert("max_age", "3600")
		form.Insert("ui_locales", "ja_JP")
		form.Insert("id_token_hint", "id_token_hint_test")
		form.Insert("login_hint", "login_hint_test")
		form.Insert("acr_values", "acr_values_test")

		m, err := easy.NewFormData("/", "POST", form)
		require.NoError(t, err)

		c := m.Echo()

		_, err = h.GetAuthenticationRequest(c)
		require.EqualError(t, err, "code=400, error=invalid_request_uri, message=scope is invalid")
	})

	t.Run("失敗: request_typeが存在しない", func(t *testing.T) {
		form := easy.NewMultipart()

		form.Insert("scope", "openid profile email")
		form.Insert("client_id", "test")
		form.Insert("redirect_uri", "https://example.test")
		form.Insert("state", "state_test")
		form.Insert("response_mode", "query")
		form.Insert("nonce", "nonce_test")
		form.Insert("display", "page")
		form.Insert("prompt", "login consent")
		form.Insert("max_age", "3600")
		form.Insert("ui_locales", "ja_JP")
		form.Insert("id_token_hint", "id_token_hint_test")
		form.Insert("login_hint", "login_hint_test")
		form.Insert("acr_values", "acr_values_test")

		m, err := easy.NewFormData("/", "POST", form)
		require.NoError(t, err)

		c := m.Echo()

		_, err = h.GetAuthenticationRequest(c)
		require.EqualError(t, err, "code=400, error=invalid_request_uri, message=request_type is invalid")
	})

	t.Run("失敗: request_typeの値が不正", func(t *testing.T) {
		form := easy.NewMultipart()

		form.Insert("scope", "openid profile email")
		form.Insert("response_type", "hogehoge")
		form.Insert("client_id", "test")
		form.Insert("redirect_uri", "https://example.test")
		form.Insert("state", "state_test")
		form.Insert("response_mode", "query")
		form.Insert("nonce", "nonce_test")
		form.Insert("display", "page")
		form.Insert("prompt", "login consent")
		form.Insert("max_age", "3600")
		form.Insert("ui_locales", "ja_JP")
		form.Insert("id_token_hint", "id_token_hint_test")
		form.Insert("login_hint", "login_hint_test")
		form.Insert("acr_values", "acr_values_test")

		m, err := easy.NewFormData("/", "POST", form)
		require.NoError(t, err)

		c := m.Echo()

		_, err = h.GetAuthenticationRequest(c)
		require.EqualError(t, err, "code=400, error=invalid_request_uri, message=request_type is invalid")
	})

	t.Run("失敗: client_idが存在しない", func(t *testing.T) {
		form := easy.NewMultipart()

		form.Insert("scope", "openid profile email")
		form.Insert("response_type", "code")
		form.Insert("redirect_uri", "https://example.test")
		form.Insert("state", "state_test")
		form.Insert("response_mode", "query")
		form.Insert("nonce", "nonce_test")
		form.Insert("display", "page")
		form.Insert("prompt", "login consent")
		form.Insert("max_age", "3600")
		form.Insert("ui_locales", "ja_JP")
		form.Insert("id_token_hint", "id_token_hint_test")
		form.Insert("login_hint", "login_hint_test")
		form.Insert("acr_values", "acr_values_test")

		m, err := easy.NewFormData("/", "POST", form)
		require.NoError(t, err)

		c := m.Echo()

		_, err = h.GetAuthenticationRequest(c)
		require.EqualError(t, err, "code=400, error=invalid_request_uri, message=client_id is required")
	})

	t.Run("失敗: redirect_uriが存在しない", func(t *testing.T) {
		form := easy.NewMultipart()

		form.Insert("scope", "openid profile email")
		form.Insert("response_type", "code")
		form.Insert("client_id", "test")
		form.Insert("state", "state_test")
		form.Insert("response_mode", "query")
		form.Insert("nonce", "nonce_test")
		form.Insert("display", "page")
		form.Insert("prompt", "login consent")
		form.Insert("max_age", "3600")
		form.Insert("ui_locales", "ja_JP")
		form.Insert("id_token_hint", "id_token_hint_test")
		form.Insert("login_hint", "login_hint_test")
		form.Insert("acr_values", "acr_values_test")

		m, err := easy.NewFormData("/", "POST", form)
		require.NoError(t, err)

		c := m.Echo()

		_, err = h.GetAuthenticationRequest(c)
		require.EqualError(t, err, "code=400, error=invalid_request_uri, message=redirect_uri is required")
	})

	t.Run("失敗: redirect_uriの値が不正", func(t *testing.T) {
		form := easy.NewMultipart()

		form.Insert("scope", "openid profile email")
		form.Insert("response_type", "code")
		form.Insert("client_id", "test")
		form.Insert("redirect_uri", "hogehoge")
		form.Insert("state", "state_test")
		form.Insert("response_mode", "query")
		form.Insert("nonce", "nonce_test")
		form.Insert("display", "page")
		form.Insert("prompt", "login consent")
		form.Insert("max_age", "3600")
		form.Insert("ui_locales", "ja_JP")
		form.Insert("id_token_hint", "id_token_hint_test")
		form.Insert("login_hint", "login_hint_test")
		form.Insert("acr_values", "acr_values_test")

		m, err := easy.NewFormData("/", "POST", form)
		require.NoError(t, err)

		c := m.Echo()

		_, err = h.GetAuthenticationRequest(c)
		require.EqualError(t, err, "code=400, error=invalid_request_uri, message=redirect_uri is invalid")
	})
}
