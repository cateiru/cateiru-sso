package src_test

import (
	"context"
	"net/url"
	"testing"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestOIDCRequireHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	t.Run("成功: プレビューを取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		client := RegisterClient(t, ctx, &u, "openid", "profile")

		r := models.ClientRedirect{
			ClientID: client.ClientID,
			URL:      "https://example.test",
			Host:     "example.test",
		}
		err := r.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()

		form.Insert("scope", "openid profile email")
		form.Insert("response_type", "code")
		form.Insert("client_id", client.ClientID)
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
		m.Cookie(cookies)

		c := m.Echo()

		err = h.OIDCRequireHandler(c)
		require.NoError(t, err)

		response := src.PublicAuthenticationRequest{}
		require.NoError(t, m.Json(&response))

		// GetPreviewResponse のテストて見ているのでここでは必要最低限で見る
		require.Equal(t, response.ClientId, client.ClientID)
	})

	t.Run("ログインしていないと、トークンを返す", func(t *testing.T) {
		client := RegisterClient(t, ctx, nil, "openid", "profile")

		r := models.ClientRedirect{
			ClientID: client.ClientID,
			URL:      "https://example.test",
			Host:     "example.test",
		}
		err := r.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := easy.NewMultipart()

		form.Insert("scope", "openid profile email")
		form.Insert("response_type", "code")
		form.Insert("client_id", client.ClientID)
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

		err = h.OIDCRequireHandler(c)
		require.NoError(t, err)

		response := src.PublicAuthenticationLoginSession{}
		require.NoError(t, m.Json(&response))

		require.NotEmpty(t, response.LoginSessionToken)
	})

	t.Run("失敗: ユーザーチェック不可", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		client := RegisterClient(t, ctx, &u, "openid", "profile")

		// Allow Ruleを設定してユーザーを弾く
		client, err := models.Clients(
			models.ClientWhere.ClientID.EQ(client.ClientID),
		).One(ctx, DB)
		require.NoError(t, err)

		client.IsAllow = true

		_, err = client.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		r := models.ClientRedirect{
			ClientID: client.ClientID,
			URL:      "https://example.test",
			Host:     "example.test",
		}
		err = r.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		// ルール作成
		allowRule := models.ClientAllowRule{
			ClientID:    client.ClientID,
			EmailDomain: null.NewString("nya.test", true),
		}
		require.NoError(t, allowRule.Insert(ctx, DB, boil.Infer()))

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()

		form.Insert("scope", "openid profile email")
		form.Insert("response_type", "code")
		form.Insert("client_id", client.ClientID)
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
		m.Cookie(cookies)

		c := m.Echo()

		err = h.OIDCRequireHandler(c)
		require.EqualError(t, err, "code=400, error=invalid_request_uri, message=user is not allowed")
	})
}

func TestOIDCLoginHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.OIDCLoginHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		client := RegisterClient(t, ctx, u, "openid", "profile")

		r := models.ClientRedirect{
			ClientID: client.ClientID,
			URL:      "https://example.test",
			Host:     "example.test",
		}
		err := r.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := easy.NewMultipart()

		state, err := lib.RandomStr(10)
		require.NoError(t, err)

		form.Insert("scope", "openid profile email")
		form.Insert("response_type", "code")
		form.Insert("client_id", client.ClientID)
		form.Insert("redirect_uri", "https://example.test")
		form.Insert("state", state)
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
		return m
	})

	t.Run("成功: submitできる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		client := RegisterClient(t, ctx, &u, "openid", "profile")

		r := models.ClientRedirect{
			ClientID: client.ClientID,
			URL:      "https://example.test",
			Host:     "example.test",
		}
		err := r.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()

		state, err := lib.RandomStr(10)
		require.NoError(t, err)

		form.Insert("scope", "openid profile email")
		form.Insert("response_type", "code")
		form.Insert("client_id", client.ClientID)
		form.Insert("redirect_uri", "https://example.test")
		form.Insert("state", state)
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
		m.Cookie(cookies)

		c := m.Echo()

		err = h.OIDCLoginHandler(c)
		require.NoError(t, err)

		response := src.OauthResponse{}
		require.NoError(t, m.Json(&response))

		require.NotEmpty(t, response.RedirectUrl)

		redirectUrl, err := url.Parse(response.RedirectUrl)
		require.NoError(t, err)

		require.Equal(t, redirectUrl.Host, "example.test")
		require.Equal(t, redirectUrl.Query().Get("state"), state, "stateが一致する")

		code := redirectUrl.Query().Get("code")
		require.NotEmpty(t, code, "codeがある")

		existSession, err := models.OauthSessions(
			models.OauthSessionWhere.Code.EQ(code),
		).Exists(ctx, DB)
		require.NoError(t, err)

		require.True(t, existSession, "DBにセッションが存在している")

		operationHistory, err := models.OperationHistories(
			models.OperationHistoryWhere.UserID.EQ(u.ID),
		).One(ctx, DB)
		require.NoError(t, err)
		require.Equal(t, operationHistory.Identifier, int8(1), "ログが保存されている")
	})

	t.Run("失敗: ユーザーチェック不可", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		client := RegisterClient(t, ctx, &u, "openid", "profile")

		// Allow Ruleを設定してユーザーを弾く
		client, err := models.Clients(
			models.ClientWhere.ClientID.EQ(client.ClientID),
		).One(ctx, DB)
		require.NoError(t, err)

		client.IsAllow = true

		_, err = client.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		r := models.ClientRedirect{
			ClientID: client.ClientID,
			URL:      "https://example.test",
			Host:     "example.test",
		}
		err = r.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		// ルール作成
		allowRule := models.ClientAllowRule{
			ClientID:    client.ClientID,
			EmailDomain: null.NewString("nya.test", true),
		}
		require.NoError(t, allowRule.Insert(ctx, DB, boil.Infer()))

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()

		state, err := lib.RandomStr(10)
		require.NoError(t, err)

		form.Insert("scope", "openid profile email")
		form.Insert("response_type", "code")
		form.Insert("client_id", client.ClientID)
		form.Insert("redirect_uri", "https://example.test")
		form.Insert("state", state)
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
		m.Cookie(cookies)

		c := m.Echo()

		err = h.OIDCLoginHandler(c)
		require.EqualError(t, err, "code=400, error=invalid_request_uri, message=user is not allowed")
	})
}

func TestOIDCCancelHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.OIDCCancelHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		client := RegisterClient(t, ctx, u, "openid", "profile")

		r := models.ClientRedirect{
			ClientID: client.ClientID,
			URL:      "https://example.test",
			Host:     "example.test",
		}
		err := r.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		form := easy.NewMultipart()

		state, err := lib.RandomStr(10)
		require.NoError(t, err)

		form.Insert("scope", "openid profile email")
		form.Insert("response_type", "code")
		form.Insert("client_id", client.ClientID)
		form.Insert("redirect_uri", "https://example.test")
		form.Insert("state", state)
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
		return m
	})

	t.Run("成功: キャンセルできる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		client := RegisterClient(t, ctx, &u, "openid", "profile")

		r := models.ClientRedirect{
			ClientID: client.ClientID,
			URL:      "https://example.test",
			Host:     "example.test",
		}
		err := r.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		cookies := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()

		state, err := lib.RandomStr(10)
		require.NoError(t, err)

		form.Insert("scope", "openid profile email")
		form.Insert("response_type", "code")
		form.Insert("client_id", client.ClientID)
		form.Insert("redirect_uri", "https://example.test")
		form.Insert("state", state)
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
		m.Cookie(cookies)

		c := m.Echo()

		err = h.OIDCCancelHandler(c)
		require.NoError(t, err)

		response := src.OauthResponse{}
		require.NoError(t, m.Json(&response))

		require.NotEmpty(t, response.RedirectUrl)

		redirectUrl, err := url.Parse(response.RedirectUrl)
		require.NoError(t, err)

		require.Equal(t, redirectUrl.Host, "example.test")
		require.Equal(t, redirectUrl.Query().Get("state"), state, "stateが一致する")

		require.Empty(t, redirectUrl.Query().Get("code"), "codeがない")
		require.Equal(t, redirectUrl.Query().Get("error"), "access_denied", "errorがaccess_denied")
	})
}
