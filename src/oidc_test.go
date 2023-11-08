package src_test

import (
	"context"
	"testing"

	"github.com/cateiru/cateiru-sso/src"
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
		clientId, _ := RegisterClient(t, ctx, &u, "openid", "profile")

		r := models.ClientRedirect{
			ClientID: clientId,
			URL:      "https://example.test",
			Host:     "example.test",
		}
		err := r.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		cookies := RegisterSession(t, ctx, &u)

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
		m.Cookie(cookies)

		c := m.Echo()

		err = h.OIDCRequireHandler(c)
		require.NoError(t, err)

		response := src.PublicAuthenticationRequest{}
		require.NoError(t, m.Json(&response))

		// GetPreviewResponse のテストて見ているのでここでは必要最低限で見る
		require.Equal(t, response.ClientId, clientId)
	})

	t.Run("ログインしていないと、トークンを返す", func(t *testing.T) {
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
		clientId, _ := RegisterClient(t, ctx, &u, "openid", "profile")

		// Allow Ruleを設定してユーザーを弾く
		client, err := models.Clients(
			models.ClientWhere.ClientID.EQ(clientId),
		).One(ctx, DB)
		require.NoError(t, err)

		client.IsAllow = true

		_, err = client.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		r := models.ClientRedirect{
			ClientID: clientId,
			URL:      "https://example.test",
			Host:     "example.test",
		}
		err = r.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		// ルール作成
		allowRule := models.ClientAllowRule{
			ClientID:    clientId,
			EmailDomain: null.NewString("nya.test", true),
		}
		require.NoError(t, allowRule.Insert(ctx, DB, boil.Infer()))

		cookies := RegisterSession(t, ctx, &u)

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
		m.Cookie(cookies)

		c := m.Echo()

		err = h.OIDCRequireHandler(c)
		require.EqualError(t, err, "code=400, error=invalid_request_uri, message=user is not allowed")
	})
}

// TODO: テスト
func TestOIDCLoginHandler(t *testing.T) {

}

// TODO: テスト
func TestOIDCCancelHandler(t *testing.T) {

}
