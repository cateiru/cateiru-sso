package src_test

import (
	"context"
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestNewAuthenticationRequest(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	t.Run("成功: リファラー未設定", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		clientId, _ := RegisterClient(t, ctx, &u, "openid", "profile")

		r := models.ClientRedirect{
			ClientID: clientId,
			URL:      "https://example.test",
			Host:     "example.test",
		}
		err := r.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := easy.NewMultipart()

		form.Insert("scope", "openid profile email")
		form.Insert("response_type", "code")
		form.Insert("client_id", clientId)
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

		// リファラーを設定する
		m.R.Header.Set("Referer", "https://example.test")

		c := m.Echo()

		authenticationRequest, err := h.NewAuthenticationRequest(ctx, c)
		require.NoError(t, err)

		require.Equal(t, authenticationRequest.Scopes, []string{"openid", "profile"}, "クライアントが有効かつ、指定されたものになる")
		require.Equal(t, authenticationRequest.ResponseType, lib.ResponseTypeAuthorizationCode)
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

		require.Equal(t, authenticationRequest.Client.ClientID, clientId)
		require.False(t, authenticationRequest.Client.IsAllow)
		require.Equal(t, authenticationRequest.AllowRules, []*models.ClientAllowRule{})
		require.Equal(t, authenticationRequest.RefererHost, "", "リファラーチェックはしていないので空")
	})

	t.Run("成功: AllowRuleが設定されている", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		clientId, _ := RegisterClient(t, ctx, &u, "openid", "profile")

		r := models.ClientRedirect{
			ClientID: clientId,
			URL:      "https://example.test",
			Host:     "example.test",
		}
		err := r.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		// AllowRuleを設定する
		RegisterAllowRules(t, ctx, clientId, false, "example.test")

		// クライアントを更新する
		client, err := models.Clients(models.ClientWhere.ClientID.EQ(clientId)).One(ctx, DB)
		require.NoError(t, err)

		client.IsAllow = true

		_, err = client.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := easy.NewMultipart()

		form.Insert("scope", "openid profile email")
		form.Insert("response_type", "code")
		form.Insert("client_id", clientId)
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

		// リファラーを設定する
		m.R.Header.Set("Referer", "https://example.test")

		c := m.Echo()

		authenticationRequest, err := h.NewAuthenticationRequest(ctx, c)
		require.NoError(t, err)

		require.Equal(t, authenticationRequest.Scopes, []string{"openid", "profile"}, "クライアントが有効かつ、指定されたものになる")
		require.Equal(t, authenticationRequest.ResponseType, lib.ResponseTypeAuthorizationCode)
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

		require.Equal(t, authenticationRequest.Client.ClientID, clientId)
		require.True(t, authenticationRequest.Client.IsAllow)
		require.Len(t, authenticationRequest.AllowRules, 1)
		require.Equal(t, authenticationRequest.AllowRules[0].EmailDomain, null.NewString("example.test", true))
		require.Equal(t, authenticationRequest.RefererHost, "", "リファラーチェックはしていないので空")
	})

	t.Run("成功: リファラー設定済み", func(t *testing.T) {
		t.Run("originの場合", func(t *testing.T) {
			email := RandomEmail(t)
			u := RegisterUser(t, ctx, email)
			clientId, _ := RegisterClient(t, ctx, &u, "openid", "profile")

			r := models.ClientRedirect{
				ClientID: clientId,
				URL:      "https://example.test",
				Host:     "example.test",
			}
			err := r.Insert(ctx, DB, boil.Infer())
			require.NoError(t, err)

			referrer := models.ClientReferrer{
				ClientID: clientId,
				Host:     "example.test",
				URL:      "https://example.test",
			}
			err = referrer.Insert(ctx, DB, boil.Infer())
			require.NoError(t, err)

			form := easy.NewMultipart()

			form.Insert("scope", "openid profile email")
			form.Insert("response_type", "code")
			form.Insert("client_id", clientId)
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

			// リファラーを設定する
			// originの場合はパス、クエリはリファラーに含まれない
			m.R.Header.Set("Referer", "https://example.test")

			c := m.Echo()

			authenticationRequest, err := h.NewAuthenticationRequest(ctx, c)
			require.NoError(t, err)

			require.Equal(t, authenticationRequest.RefererHost, "example.test")
		})

		t.Run("unsafe-url", func(t *testing.T) {
			email := RandomEmail(t)
			u := RegisterUser(t, ctx, email)
			clientId, _ := RegisterClient(t, ctx, &u, "openid", "profile")

			r := models.ClientRedirect{
				ClientID: clientId,
				URL:      "https://example.test",
				Host:     "example.test",
			}
			err := r.Insert(ctx, DB, boil.Infer())
			require.NoError(t, err)

			referrer := models.ClientReferrer{
				ClientID: clientId,
				Host:     "example.test",
				URL:      "https://example.test",
			}
			err = referrer.Insert(ctx, DB, boil.Infer())
			require.NoError(t, err)

			form := easy.NewMultipart()

			form.Insert("scope", "openid profile email")
			form.Insert("response_type", "code")
			form.Insert("client_id", clientId)
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

			// リファラーを設定する
			m.R.Header.Set("Referer", "https://example.test/hoge/huga?aa=test")

			c := m.Echo()

			authenticationRequest, err := h.NewAuthenticationRequest(ctx, c)
			require.NoError(t, err)

			require.Equal(t, authenticationRequest.RefererHost, "example.test")
		})
	})

	t.Run("失敗: scopeが存在しない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		clientId, _ := RegisterClient(t, ctx, &u, "openid", "profile")

		r := models.ClientRedirect{
			ClientID: clientId,
			URL:      "https://example.test",
			Host:     "example.test",
		}
		err := r.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := easy.NewMultipart()

		form.Insert("response_type", "code")
		form.Insert("client_id", clientId)
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

		_, err = h.NewAuthenticationRequest(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request_uri, message=scope is required")
	})

	t.Run("失敗: scopeにopenidが存在しない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		clientId, _ := RegisterClient(t, ctx, &u, "openid", "profile")

		r := models.ClientRedirect{
			ClientID: clientId,
			URL:      "https://example.test",
			Host:     "example.test",
		}
		err := r.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := easy.NewMultipart()

		form.Insert("scope", "profile email")
		form.Insert("response_type", "code")
		form.Insert("client_id", clientId)
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

		_, err = h.NewAuthenticationRequest(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request_uri, message=scope is invalid")
	})

	t.Run("失敗: request_typeが存在しない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		clientId, _ := RegisterClient(t, ctx, &u, "openid", "profile")

		r := models.ClientRedirect{
			ClientID: clientId,
			URL:      "https://example.test",
			Host:     "example.test",
		}
		err := r.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := easy.NewMultipart()

		form.Insert("scope", "openid profile email")
		form.Insert("client_id", clientId)
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

		_, err = h.NewAuthenticationRequest(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request_uri, message=request_type is invalid")
	})

	t.Run("失敗: request_typeの値が不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		clientId, _ := RegisterClient(t, ctx, &u, "openid", "profile")

		r := models.ClientRedirect{
			ClientID: clientId,
			URL:      "https://example.test",
			Host:     "example.test",
		}
		err := r.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := easy.NewMultipart()

		form.Insert("scope", "openid profile email")
		form.Insert("response_type", "hogehoge")
		form.Insert("client_id", clientId)
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

		_, err = h.NewAuthenticationRequest(ctx, c)
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

		_, err = h.NewAuthenticationRequest(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request_uri, message=client_id is required")
	})

	t.Run("失敗: clientが存在しない", func(t *testing.T) {
		form := easy.NewMultipart()

		form.Insert("scope", "openid profile email")
		form.Insert("response_type", "code")
		form.Insert("client_id", "clienthogehoge")
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

		_, err = h.NewAuthenticationRequest(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request_uri, message=client_id is invalid")
	})

	t.Run("失敗: redirect_uriが存在しない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		clientId, _ := RegisterClient(t, ctx, &u, "openid", "profile")

		r := models.ClientRedirect{
			ClientID: clientId,
			URL:      "https://example.test",
			Host:     "example.test",
		}
		err := r.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := easy.NewMultipart()

		form.Insert("scope", "openid profile email")
		form.Insert("response_type", "code")
		form.Insert("client_id", clientId)
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

		_, err = h.NewAuthenticationRequest(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request_uri, message=redirect_uri is required")
	})

	t.Run("失敗: redirect_uriの値が不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		clientId, _ := RegisterClient(t, ctx, &u, "openid", "profile")

		r := models.ClientRedirect{
			ClientID: clientId,
			URL:      "https://example.test",
			Host:     "example.test",
		}
		err := r.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := easy.NewMultipart()

		form.Insert("scope", "openid profile email")
		form.Insert("response_type", "code")
		form.Insert("client_id", clientId)
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

		_, err = h.NewAuthenticationRequest(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request_uri, message=redirect_uri is invalid")
	})

	t.Run("失敗: クライアントに登録しているRedirectURIがない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		clientId, _ := RegisterClient(t, ctx, &u, "openid", "profile")

		r := models.ClientRedirect{
			ClientID: clientId,
			URL:      "https://example.test",
			Host:     "example.test",
		}
		err := r.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := easy.NewMultipart()

		form.Insert("scope", "openid profile email")
		form.Insert("response_type", "code")
		form.Insert("client_id", clientId)
		form.Insert("redirect_uri", "https://example2.test")
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

		_, err = h.NewAuthenticationRequest(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request_uri, message=redirect_uri is invalid")
	})

	t.Run("失敗: 設定されているリファラーに存在しない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		clientId, _ := RegisterClient(t, ctx, &u, "openid", "profile")

		r := models.ClientRedirect{
			ClientID: clientId,
			URL:      "https://example.test",
			Host:     "example.test",
		}
		err := r.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		referrer := models.ClientReferrer{
			ClientID: clientId,
			Host:     "example.test",
			URL:      "https://example.test",
		}
		err = referrer.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := easy.NewMultipart()

		form.Insert("scope", "openid profile email")
		form.Insert("response_type", "code")
		form.Insert("client_id", clientId)
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

		m.R.Header.Set("Referer", "https://hoge.test")

		c := m.Echo()

		_, err = h.NewAuthenticationRequest(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request_uri, message=referer is invalid")
	})

	t.Run("失敗: リファラーが設定されているけど、リファラーが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		clientId, _ := RegisterClient(t, ctx, &u, "openid", "profile")

		r := models.ClientRedirect{
			ClientID: clientId,
			URL:      "https://example.test",
			Host:     "example.test",
		}
		err := r.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		referrer := models.ClientReferrer{
			ClientID: clientId,
			Host:     "example.test",
			URL:      "https://example.test",
		}
		err = referrer.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := easy.NewMultipart()

		form.Insert("scope", "openid profile email")
		form.Insert("response_type", "code")
		form.Insert("client_id", clientId)
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

		_, err = h.NewAuthenticationRequest(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request_uri, message=referer is invalid")
	})
}

func TestGetPreviewResponse(t *testing.T) {
	ctx := context.Background()

	t.Run("通常のクライアント", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		a := src.AuthenticationRequest{
			Scopes: []string{
				"openid",
				"profile",
			},
			ResponseType: lib.ResponseTypeAuthorizationCode,
			RedirectUri: &url.URL{
				Scheme: "https",
				Host:   "example.test",
				Path:   "/",
			},
			State:        null.NewString("state_test", true),
			ResponseMode: lib.ResponseModeQuery,
			Nonce:        null.NewString("nonce_test", true),
			Display:      lib.DisplayPage,
			Prompts: []lib.Prompt{
				lib.PromptConsent,
			},
			MaxAge: uint64(0),
			UiLocales: []string{
				"ja_JP",
			},
			IdTokenHint: null.NewString("id_token_hint_test", true),
			LoginHint:   null.NewString("login_hint_test", true),
			AcrValues:   null.NewString("acr_values_test", true),

			Client: &models.Client{
				ClientID:    "client_id_test",
				Name:        "client_name_test",
				Description: null.NewString("client_description_test", true),
				Image:       null.NewString("client_image_test", true),

				OrgID:         null.NewString("", false),
				OrgMemberOnly: false,

				OwnerUserID: u.ID,
			},

			AllowRules: []*models.ClientAllowRule{
				{
					EmailDomain: null.NewString("example.test", true),
				},
				{
					UserID: null.NewString("user_id", true),
				},
			},
		}

		response, err := a.GetPreviewResponse(ctx, C.OauthLoginSessionPeriod, DB, "")
		require.NoError(t, err)

		require.Equal(t, *response, src.PublicAuthenticationRequest{
			ClientId:          "client_id_test",
			ClientName:        "client_name_test",
			ClientDescription: null.NewString("client_description_test", true),
			Image:             null.NewString("client_image_test", true),

			OrgName:       null.NewString("", false),
			OrgImage:      null.NewString("", false),
			OrgMemberOnly: false,

			Scopes: []string{
				"openid",
				"profile",
			},
			RedirectUri:  "https://example.test/",
			ResponseType: "code",

			RegisterUserName:  u.UserName,
			RegisterUserImage: u.Avatar,

			Prompts: []lib.Prompt{
				lib.PromptConsent,
			},
		})
	})

	t.Run("組織のクライアント", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		org := RegisterOrg(t, ctx)

		a := src.AuthenticationRequest{
			Scopes: []string{
				"openid",
				"profile",
			},
			ResponseType: lib.ResponseTypeAuthorizationCode,
			RedirectUri: &url.URL{
				Scheme: "https",
				Host:   "example.test",
				Path:   "/",
			},
			State:        null.NewString("state_test", true),
			ResponseMode: lib.ResponseModeQuery,
			Nonce:        null.NewString("nonce_test", true),
			Display:      lib.DisplayPage,
			Prompts: []lib.Prompt{
				lib.PromptConsent,
			},
			MaxAge: uint64(0),
			UiLocales: []string{
				"ja_JP",
			},
			IdTokenHint: null.NewString("id_token_hint_test", true),
			LoginHint:   null.NewString("login_hint_test", true),
			AcrValues:   null.NewString("acr_values_test", true),

			Client: &models.Client{
				ClientID:    "client_id_test",
				Name:        "client_name_test",
				Description: null.NewString("client_description_test", true),
				Image:       null.NewString("client_image_test", true),

				OrgID:         null.NewString(org, true),
				OwnerUserID:   u.ID,
				OrgMemberOnly: true,
			},

			AllowRules: []*models.ClientAllowRule{
				{
					EmailDomain: null.NewString("example.test", true),
				},
				{
					UserID: null.NewString("user_id", true),
				},
			},
		}

		response, err := a.GetPreviewResponse(ctx, C.OauthLoginSessionPeriod, DB, "")
		require.NoError(t, err)

		require.Equal(t, *response, src.PublicAuthenticationRequest{
			ClientId:          "client_id_test",
			ClientName:        "client_name_test",
			ClientDescription: null.NewString("client_description_test", true),
			Image:             null.NewString("client_image_test", true),

			OrgName:       null.NewString("test", true),
			OrgImage:      null.NewString("", false),
			OrgMemberOnly: true,

			Scopes: []string{
				"openid",
				"profile",
			},
			RedirectUri:  "https://example.test/",
			ResponseType: "code",

			RegisterUserName:  u.UserName,
			RegisterUserImage: u.Avatar,

			Prompts: []lib.Prompt{
				lib.PromptConsent,
			},
		})
	})

	t.Run("成功: promptにloginがある場合はトークンも作成する", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u, "openid", "profile")
		client, err := models.Clients(
			models.ClientWhere.ClientID.EQ(clientId),
		).One(ctx, DB)
		require.NoError(t, err)

		a := src.AuthenticationRequest{
			Scopes: []string{
				"openid",
				"profile",
			},
			ResponseType: lib.ResponseTypeAuthorizationCode,
			RedirectUri: &url.URL{
				Scheme: "https",
				Host:   "example.test",
				Path:   "/",
			},
			State:        null.NewString("state_test", true),
			ResponseMode: lib.ResponseModeQuery,
			Nonce:        null.NewString("nonce_test", true),
			Display:      lib.DisplayPage,
			Prompts: []lib.Prompt{
				lib.PromptConsent,
				lib.PromptLogin,
			},
			MaxAge: uint64(0),
			UiLocales: []string{
				"ja_JP",
			},
			IdTokenHint: null.NewString("id_token_hint_test", true),
			LoginHint:   null.NewString("login_hint_test", true),
			AcrValues:   null.NewString("acr_values_test", true),

			Client: client,

			AllowRules: []*models.ClientAllowRule{
				{
					EmailDomain: null.NewString("example.test", true),
				},
				{
					UserID: null.NewString("user_id", true),
				},
			},
		}

		response, err := a.GetPreviewResponse(ctx, C.OauthLoginSessionPeriod, DB, "")
		require.NoError(t, err)

		require.NotNil(t, response.LoginSession)

		s, err := models.OauthLoginSessions(
			models.OauthLoginSessionWhere.Token.EQ(response.LoginSession.LoginSessionToken),
		).One(ctx, DB)
		require.NoError(t, err)
		require.Equal(t, s.ClientID, client.ClientID)
	})

	t.Run("成功: promptにloginがあり引数のトークンがすでにログイン済みの場合は何もしない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u, "openid", "profile")
		client, err := models.Clients(
			models.ClientWhere.ClientID.EQ(clientId),
		).One(ctx, DB)
		require.NoError(t, err)

		token, err := lib.RandomStr(31)
		require.NoError(t, err)
		session := models.OauthLoginSession{
			ClientID:     client.ClientID,
			Token:        token,
			ReferrerHost: null.NewString("", false),
			Period:       time.Now().Add(10 * time.Hour),
			LoginOk:      true, // ログイン済み
		}
		require.NoError(t, session.Insert(ctx, DB, boil.Infer()))

		a := src.AuthenticationRequest{
			Scopes: []string{
				"openid",
				"profile",
			},
			ResponseType: lib.ResponseTypeAuthorizationCode,
			RedirectUri: &url.URL{
				Scheme: "https",
				Host:   "example.test",
				Path:   "/",
			},
			State:        null.NewString("state_test", true),
			ResponseMode: lib.ResponseModeQuery,
			Nonce:        null.NewString("nonce_test", true),
			Display:      lib.DisplayPage,
			Prompts: []lib.Prompt{
				lib.PromptConsent,
				lib.PromptLogin,
			},
			MaxAge: uint64(0),
			UiLocales: []string{
				"ja_JP",
			},
			IdTokenHint: null.NewString("id_token_hint_test", true),
			LoginHint:   null.NewString("login_hint_test", true),
			AcrValues:   null.NewString("acr_values_test", true),

			Client: client,

			AllowRules: []*models.ClientAllowRule{
				{
					EmailDomain: null.NewString("example.test", true),
				},
				{
					UserID: null.NewString("user_id", true),
				},
			},
		}

		response, err := a.GetPreviewResponse(ctx, C.OauthLoginSessionPeriod, DB, token)
		require.NoError(t, err)

		require.Nil(t, response.LoginSession)
	})

	t.Run("成功: promptにloginがあり引数のトークンが不正の場合はログインセッションが作られる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u, "openid", "profile")
		client, err := models.Clients(
			models.ClientWhere.ClientID.EQ(clientId),
		).One(ctx, DB)
		require.NoError(t, err)

		token, err := lib.RandomStr(31)
		require.NoError(t, err)

		a := src.AuthenticationRequest{
			Scopes: []string{
				"openid",
				"profile",
			},
			ResponseType: lib.ResponseTypeAuthorizationCode,
			RedirectUri: &url.URL{
				Scheme: "https",
				Host:   "example.test",
				Path:   "/",
			},
			State:        null.NewString("state_test", true),
			ResponseMode: lib.ResponseModeQuery,
			Nonce:        null.NewString("nonce_test", true),
			Display:      lib.DisplayPage,
			Prompts: []lib.Prompt{
				lib.PromptConsent,
				lib.PromptLogin,
			},
			MaxAge: uint64(0),
			UiLocales: []string{
				"ja_JP",
			},
			IdTokenHint: null.NewString("id_token_hint_test", true),
			LoginHint:   null.NewString("login_hint_test", true),
			AcrValues:   null.NewString("acr_values_test", true),

			Client: client,

			AllowRules: []*models.ClientAllowRule{
				{
					EmailDomain: null.NewString("example.test", true),
				},
				{
					UserID: null.NewString("user_id", true),
				},
			},
		}

		response, err := a.GetPreviewResponse(ctx, C.OauthLoginSessionPeriod, DB, token)
		require.NoError(t, err)

		require.NotNil(t, response.LoginSession)

		s, err := models.OauthLoginSessions(
			models.OauthLoginSessionWhere.Token.EQ(response.LoginSession.LoginSessionToken),
		).One(ctx, DB)
		require.NoError(t, err)
		require.Equal(t, s.ClientID, client.ClientID)
	})

	t.Run("成功: promptにloginがあり引数のトークンが有効期限切れの場合は新たにログインセッションが作られる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u, "openid", "profile")
		client, err := models.Clients(
			models.ClientWhere.ClientID.EQ(clientId),
		).One(ctx, DB)
		require.NoError(t, err)

		token, err := lib.RandomStr(31)
		require.NoError(t, err)
		session := models.OauthLoginSession{
			ClientID:     client.ClientID,
			Token:        token,
			ReferrerHost: null.NewString("", false),
			Period:       time.Now().Add(-10 * time.Hour), // 有効期限切れ
			LoginOk:      true,                            // ログイン済み
		}
		require.NoError(t, session.Insert(ctx, DB, boil.Infer()))

		a := src.AuthenticationRequest{
			Scopes: []string{
				"openid",
				"profile",
			},
			ResponseType: lib.ResponseTypeAuthorizationCode,
			RedirectUri: &url.URL{
				Scheme: "https",
				Host:   "example.test",
				Path:   "/",
			},
			State:        null.NewString("state_test", true),
			ResponseMode: lib.ResponseModeQuery,
			Nonce:        null.NewString("nonce_test", true),
			Display:      lib.DisplayPage,
			Prompts: []lib.Prompt{
				lib.PromptConsent,
				lib.PromptLogin,
			},
			MaxAge: uint64(0),
			UiLocales: []string{
				"ja_JP",
			},
			IdTokenHint: null.NewString("id_token_hint_test", true),
			LoginHint:   null.NewString("login_hint_test", true),
			AcrValues:   null.NewString("acr_values_test", true),

			Client: client,

			AllowRules: []*models.ClientAllowRule{
				{
					EmailDomain: null.NewString("example.test", true),
				},
				{
					UserID: null.NewString("user_id", true),
				},
			},
		}

		response, err := a.GetPreviewResponse(ctx, C.OauthLoginSessionPeriod, DB, token)
		require.NoError(t, err)

		require.NotNil(t, response.LoginSession)

		s, err := models.OauthLoginSessions(
			models.OauthLoginSessionWhere.Token.EQ(response.LoginSession.LoginSessionToken),
		).One(ctx, DB)
		require.NoError(t, err)
		require.Equal(t, s.ClientID, client.ClientID)
	})

	t.Run("失敗: promptにloginがあり引数のトークンはあるがログインしていない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u, "openid", "profile")
		client, err := models.Clients(
			models.ClientWhere.ClientID.EQ(clientId),
		).One(ctx, DB)
		require.NoError(t, err)

		token, err := lib.RandomStr(31)
		require.NoError(t, err)
		session := models.OauthLoginSession{
			ClientID:     client.ClientID,
			Token:        token,
			ReferrerHost: null.NewString("", false),
			Period:       time.Now().Add(10 * time.Hour),
			LoginOk:      false, // 未ログイン
		}
		require.NoError(t, session.Insert(ctx, DB, boil.Infer()))

		a := src.AuthenticationRequest{
			Scopes: []string{
				"openid",
				"profile",
			},
			ResponseType: lib.ResponseTypeAuthorizationCode,
			RedirectUri: &url.URL{
				Scheme: "https",
				Host:   "example.test",
				Path:   "/",
			},
			State:        null.NewString("state_test", true),
			ResponseMode: lib.ResponseModeQuery,
			Nonce:        null.NewString("nonce_test", true),
			Display:      lib.DisplayPage,
			Prompts: []lib.Prompt{
				lib.PromptConsent,
				lib.PromptLogin,
			},
			MaxAge: uint64(0),
			UiLocales: []string{
				"ja_JP",
			},
			IdTokenHint: null.NewString("id_token_hint_test", true),
			LoginHint:   null.NewString("login_hint_test", true),
			AcrValues:   null.NewString("acr_values_test", true),

			Client: client,

			AllowRules: []*models.ClientAllowRule{
				{
					EmailDomain: null.NewString("example.test", true),
				},
				{
					UserID: null.NewString("user_id", true),
				},
			},
		}

		_, err = a.GetPreviewResponse(ctx, C.OauthLoginSessionPeriod, DB, token)
		require.EqualError(t, err, "code=400, error=invalid_request_uri, message=no login")
	})
}

func TestGetLoginSession(t *testing.T) {
	ctx := context.Background()

	email := RandomEmail(t)
	u := RegisterUser(t, ctx, email)

	clientId, _ := RegisterClient(t, ctx, &u, "openid", "profile")
	client, err := models.Clients(models.ClientWhere.ClientID.EQ(clientId)).One(ctx, DB)
	require.NoError(t, err)

	t.Run("成功", func(t *testing.T) {
		a := src.AuthenticationRequest{
			Scopes: []string{
				"openid",
				"profile",
			},
			ResponseType: lib.ResponseTypeAuthorizationCode,
			RedirectUri: &url.URL{
				Scheme: "https",
				Host:   "example.test",
				Path:   "/",
			},
			State:        null.NewString("state_test", true),
			ResponseMode: lib.ResponseModeQuery,
			Nonce:        null.NewString("nonce_test", true),
			Display:      lib.DisplayPage,
			Prompts: []lib.Prompt{
				lib.PromptConsent,
			},
			MaxAge: uint64(0),
			UiLocales: []string{
				"ja_JP",
			},
			IdTokenHint: null.NewString("id_token_hint_test", true),
			LoginHint:   null.NewString("login_hint_test", true),
			AcrValues:   null.NewString("acr_values_test", true),

			Client: client,

			AllowRules: []*models.ClientAllowRule{
				{
					EmailDomain: null.NewString("example.test", true),
				},
				{
					UserID: null.NewString("user_id", true),
				},
			},
			RefererHost: "example.test",
		}

		request, err := a.GetLoginSession(ctx, C.OauthLoginSessionPeriod, DB)
		require.NoError(t, err)

		token, err := models.OauthLoginSessions(
			models.OauthLoginSessionWhere.Token.EQ(request.LoginSessionToken),
		).One(ctx, DB)
		require.NoError(t, err)

		require.Equal(t, token.ClientID, client.ClientID)
		require.Equal(t, token.ReferrerHost.String, a.RefererHost)
		require.Equal(t, token.Period.Minute(), request.LimitDate.Minute())

	})
}

func TestCheckUserAuthenticationPossible(t *testing.T) {
	ctx := context.Background()

	t.Run("ルールがない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		a := src.AuthenticationRequest{
			AllowRules: []*models.ClientAllowRule{},
			Client: &models.Client{
				OrgID:         null.NewString("", false),
				OrgMemberOnly: false,
			},
		}

		ok, err := a.CheckUserAuthenticationPossible(ctx, DB, &u)
		require.NoError(t, err)
		require.True(t, ok)
	})

	t.Run("ユーザーIDが一致", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		a := src.AuthenticationRequest{
			AllowRules: []*models.ClientAllowRule{
				{
					UserID: null.NewString(u.ID, true),
				},
			},
			Client: &models.Client{
				OrgID:         null.NewString("", false),
				OrgMemberOnly: false,
			},
		}

		ok, err := a.CheckUserAuthenticationPossible(ctx, DB, &u)
		require.NoError(t, err)
		require.True(t, ok)
	})

	t.Run("ユーザーのメールが一致", func(t *testing.T) {
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		email := fmt.Sprintf("%s@example.test", r)
		u := RegisterUser(t, ctx, email)

		a := src.AuthenticationRequest{
			AllowRules: []*models.ClientAllowRule{
				{
					EmailDomain: null.NewString("example.test", true),
				},
			},
			Client: &models.Client{
				OrgID:         null.NewString("", false),
				OrgMemberOnly: false,
			},
		}

		ok, err := a.CheckUserAuthenticationPossible(ctx, DB, &u)
		require.NoError(t, err)
		require.True(t, ok)
	})

	t.Run("OrgMemberOnlyがtrueかつユーザーがorgに所属している", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)

		a := src.AuthenticationRequest{
			AllowRules: []*models.ClientAllowRule{},
			Client: &models.Client{
				OrgID:         null.NewString(orgId, true),
				OrgMemberOnly: true,
			},
		}

		ok, err := a.CheckUserAuthenticationPossible(ctx, DB, &u)
		require.NoError(t, err)
		require.True(t, ok)
	})

	t.Run("失敗: OrgMemberOnlyがtrueかつユーザーがorgに所属していない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx)

		a := src.AuthenticationRequest{
			AllowRules: []*models.ClientAllowRule{},
			Client: &models.Client{
				OrgID:         null.NewString(orgId, true),
				OrgMemberOnly: true,
			},
		}

		ok, err := a.CheckUserAuthenticationPossible(ctx, DB, &u)
		require.NoError(t, err)
		require.False(t, ok)
	})

	t.Run("失敗: ルールに一致しているけどOrgMemberOnlyがtrueかつユーザーがorgに所属していない", func(t *testing.T) {
		r, err := lib.RandomStr(10)
		require.NoError(t, err)
		email := fmt.Sprintf("%s@example.test", r)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx)

		a := src.AuthenticationRequest{
			AllowRules: []*models.ClientAllowRule{
				{
					EmailDomain: null.NewString("example.test", true),
				},
			},
			Client: &models.Client{
				OrgID:         null.NewString(orgId, true),
				OrgMemberOnly: true,
			},
		}

		ok, err := a.CheckUserAuthenticationPossible(ctx, DB, &u)
		require.NoError(t, err)
		require.False(t, ok)
	})

	t.Run("失敗: ルールが設定されているけど、そのルールに一致しない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		a := src.AuthenticationRequest{
			AllowRules: []*models.ClientAllowRule{
				{
					UserID: null.NewString("123", true),
				},
			},
			Client: &models.Client{
				OrgID:         null.NewString("", false),
				OrgMemberOnly: false,
			},
		}

		ok, err := a.CheckUserAuthenticationPossible(ctx, DB, &u)
		require.NoError(t, err)
		require.False(t, ok)
	})
}

func TestSetLoggedInOauthLoginSession(t *testing.T) {
	ctx := context.Background()

	t.Run("LoginOkがtrueになっている", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		clientId, _ := RegisterClient(t, ctx, &u, "openid")

		token, err := lib.RandomStr(31)
		require.NoError(t, err)
		session := models.OauthLoginSession{
			Token:        token,
			ClientID:     clientId,
			ReferrerHost: null.NewString("", false),
			Period:       time.Now().Add(1 * time.Hour),
		}
		require.NoError(t, session.Insert(ctx, DB, boil.Infer()))

		err = src.SetLoggedInOauthLoginSession(ctx, DB, token)
		require.NoError(t, err)

		s, err := models.OauthLoginSessions(
			models.OauthLoginSessionWhere.Token.EQ(token),
		).One(ctx, DB)
		require.NoError(t, err)

		require.True(t, s.LoginOk)
	})

	t.Run("トークンが存在しない", func(t *testing.T) {
		token, err := lib.RandomStr(31)
		require.NoError(t, err)

		err = src.SetLoggedInOauthLoginSession(ctx, DB, token)
		require.NoError(t, err, "エラーにはならない")
	})

	t.Run("トークンが有効期限切れ", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		clientId, _ := RegisterClient(t, ctx, &u, "openid")

		token, err := lib.RandomStr(31)
		require.NoError(t, err)
		session := models.OauthLoginSession{
			Token:        token,
			ClientID:     clientId,
			ReferrerHost: null.NewString("", false),
			Period:       time.Now().Add(-1 * time.Hour), // 過去
		}
		require.NoError(t, session.Insert(ctx, DB, boil.Infer()))

		err = src.SetLoggedInOauthLoginSession(ctx, DB, token)
		require.NoError(t, err)

		s, err := models.OauthLoginSessions(
			models.OauthLoginSessionWhere.Token.EQ(token),
		).One(ctx, DB)
		require.NoError(t, err)

		require.False(t, s.LoginOk, "有効期限が切れているので更新されていない")
	})
}
