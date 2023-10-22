package src_test

import (
	"context"
	"testing"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestNewOidcRequestAuthenticationCodeFlow(t *testing.T) {
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

		request, err := h.NewOidcRequest(ctx, c)
		require.NoError(t, err)

		authenticationRequest := request.(*src.AuthenticationRequest)

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

		request, err := h.NewOidcRequest(ctx, c)
		require.NoError(t, err)

		authenticationRequest := request.(*src.AuthenticationRequest)

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

			request, err := h.NewOidcRequest(ctx, c)
			require.NoError(t, err)

			authenticationRequest := request.(*src.AuthenticationRequest)

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

			request, err := h.NewOidcRequest(ctx, c)
			require.NoError(t, err)

			authenticationRequest := request.(*src.AuthenticationRequest)

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

		_, err = h.NewOidcRequest(ctx, c)
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

		_, err = h.NewOidcRequest(ctx, c)
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

		_, err = h.NewOidcRequest(ctx, c)
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

		_, err = h.NewOidcRequest(ctx, c)
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

		_, err = h.NewOidcRequest(ctx, c)
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

		_, err = h.NewOidcRequest(ctx, c)
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

		_, err = h.NewOidcRequest(ctx, c)
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

		_, err = h.NewOidcRequest(ctx, c)
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

		_, err = h.NewOidcRequest(ctx, c)
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

		_, err = h.NewOidcRequest(ctx, c)
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

		_, err = h.NewOidcRequest(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request_uri, message=referer is invalid")
	})
}

func TestNewOidcRequestImplicitFlow(t *testing.T) {
	// TODO: 実装したら追加する
}

func TestNewOidcRequestHybridFlow(t *testing.T) {
	// TODO: 実装したら追加する
}
