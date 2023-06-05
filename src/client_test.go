package src_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/stretchr/testify/require"
)

func TestClientHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.ClientHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		return m
	})

	t.Run("成功: client_idを指定するとそのクライアントを取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u, "openid", "profile")

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientHandler(c)
		require.NoError(t, err)

		response := src.ClientResponse{}
		require.NoError(t, m.Json(&response))

		require.Equal(t, response.ClientID, clientId)
	})

	t.Run("成功: client_idを指定しないと自分のすべてのクライアントを取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		RegisterClient(t, ctx, &u, "openid", "profile")
		RegisterClient(t, ctx, &u, "openid", "profile")

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientHandler(c)
		require.NoError(t, err)

		response := []src.ClientResponse{}
		require.NoError(t, m.Json(&response))

		require.Len(t, response, 2)
	})

	t.Run("失敗: client_idが存在しない値", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		noExistClientId, err := lib.RandomStr(32)
		require.NoError(t, err)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", noExistClientId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientHandler(c)
		require.EqualError(t, err, "code=404, message=client not found")
	})

	t.Run("失敗: client_idが指定するクライアントが自分のものではない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		clientId, _ := RegisterClient(t, ctx, &u2, "openid", "profile")

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientHandler(c)
		require.EqualError(t, err, "code=404, message=client not found")
	})
}

func TestClientCreateHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	t.Run("成功: クライアントを新規作成できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("name", "test")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile")
		form.Insert("redirect_url_count", "2")
		form.Insert("redirect_url_0", "https://aaaa.test")
		form.Insert("redirect_url_1", "https://bbbb.test")
		form.Insert("referrer_url_count", "2")
		form.Insert("referrer_url_0", "https://aaaa.test")
		form.Insert("referrer_url_1", "https://bbbb.test")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientCreateHandler(c)
		require.NoError(t, err)

		// チェック

		response := src.ClientResponse{}
		require.NoError(t, m.Json(&response))

		require.Equal(t, response.Name, "test")
		require.Equal(t, response.IsAllow, false)

		// クライアント
		client, err := models.Clients(
			models.ClientWhere.ClientID.EQ(response.ClientID),
		).One(ctx, DB)
		require.NoError(t, err)

		require.Equal(t, client.Name, "test")
		require.Equal(t, client.IsAllow, false)

		// スコープ
		scopes, err := models.ClientScopes(
			models.ClientScopeWhere.ClientID.EQ(response.ClientID),
		).All(ctx, DB)
		require.NoError(t, err)

		require.Len(t, scopes, 2)

		// リダイレクトURL
		redirectUrls, err := models.ClientRedirects(
			models.ClientRedirectWhere.ClientID.EQ(response.ClientID),
		).All(ctx, DB)
		require.NoError(t, err)

		require.Len(t, redirectUrls, 2)

		// リファラーURL
		referrerUrls, err := models.ClientReferrers(
			models.ClientReferrerWhere.ClientID.EQ(response.ClientID),
		).All(ctx, DB)
		require.NoError(t, err)

		require.Len(t, referrerUrls, 2)
	})

	t.Run("成功: 画像を設定して新規作成", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("name", "test")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile")
		form.Insert("redirect_url_count", "2")
		form.Insert("redirect_url_0", "https://aaaa.test")
		form.Insert("redirect_url_1", "https://bbbb.test")
		form.Insert("referrer_url_count", "2")
		form.Insert("referrer_url_0", "https://aaaa.test")
		form.Insert("referrer_url_1", "https://bbbb.test")

		image, err := os.Open("./test_sample_image.png")
		require.NoError(t, err)
		defer image.Close()
		form.InsertFile("image", image)

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientCreateHandler(c)
		require.NoError(t, err)

		// チェック

		response := src.ClientResponse{}
		require.NoError(t, m.Json(&response))

		require.Equal(t, response.Name, "test")
		require.Equal(t, response.IsAllow, false)

		// クライアント
		client, err := models.Clients(
			models.ClientWhere.ClientID.EQ(response.ClientID),
		).One(ctx, DB)
		require.NoError(t, err)

		require.Equal(t, client.Name, "test")
		require.Equal(t, client.IsAllow, false)

		// 画像のリンクが入っている
		require.True(t, client.Image.Valid)

		// スコープ
		scopes, err := models.ClientScopes(
			models.ClientScopeWhere.ClientID.EQ(response.ClientID),
		).All(ctx, DB)
		require.NoError(t, err)

		require.Len(t, scopes, 2)

		// リダイレクトURL
		redirectUrls, err := models.ClientRedirects(
			models.ClientRedirectWhere.ClientID.EQ(response.ClientID),
		).All(ctx, DB)
		require.NoError(t, err)

		require.Len(t, redirectUrls, 2)

		// リファラーURL
		referrerUrls, err := models.ClientReferrers(
			models.ClientReferrerWhere.ClientID.EQ(response.ClientID),
		).All(ctx, DB)
		require.NoError(t, err)

		require.Len(t, referrerUrls, 2)
	})

	t.Run("失敗: promptの値が不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("name", "test")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile")
		form.Insert("redirect_url_count", "2")
		form.Insert("redirect_url_0", "https://aaaa.test")
		form.Insert("redirect_url_1", "https://bbbb.test")
		form.Insert("referrer_url_count", "2")
		form.Insert("referrer_url_0", "https://aaaa.test")
		form.Insert("referrer_url_1", "https://bbbb.test")

		form.Insert("prompt", "hogehoge")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientCreateHandler(c)
		require.EqualError(t, err, "code=400, message=prompt is invalid")
	})

	t.Run("失敗: スコープの値が不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("name", "test")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile hogehoge")
		form.Insert("redirect_url_count", "2")
		form.Insert("redirect_url_0", "https://aaaa.test")
		form.Insert("redirect_url_1", "https://bbbb.test")
		form.Insert("referrer_url_count", "2")
		form.Insert("referrer_url_0", "https://aaaa.test")
		form.Insert("referrer_url_1", "https://bbbb.test")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientCreateHandler(c)
		require.EqualError(t, err, "code=400, message=scope `hogehoge` is invalid")
	})

	t.Run("失敗: クライアントの作成上限が超えている", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		for i := 0; i <= C.ClientMaxCreated; i++ {
			RegisterClient(t, ctx, &u)
		}

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("name", "test")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile")
		form.Insert("redirect_url_count", "2")
		form.Insert("redirect_url_0", "https://aaaa.test")
		form.Insert("redirect_url_1", "https://bbbb.test")
		form.Insert("referrer_url_count", "2")
		form.Insert("referrer_url_0", "https://aaaa.test")
		form.Insert("referrer_url_1", "https://bbbb.test")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientCreateHandler(c)
		require.EqualError(t, err, "code=400, message=too many clients")
	})

	t.Run("失敗: リダイレクトURLが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("name", "test")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile")
		form.Insert("redirect_url_count", "2")
		form.Insert("redirect_url_0", "nyancat")
		form.Insert("redirect_url_1", "https://bbbb.test")
		form.Insert("referrer_url_count", "2")
		form.Insert("referrer_url_0", "https://aaaa.test")
		form.Insert("referrer_url_1", "https://bbbb.test")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientCreateHandler(c)
		require.EqualError(t, err, "code=400, message=referrer_url `nyancat` is invalid")
	})

	t.Run("失敗: リダイレクトURLのForm指定が違う", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("name", "test")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile")
		form.Insert("redirect_url_count", "3")
		form.Insert("redirect_url_0", "https://aaaa.test")
		form.Insert("redirect_url_1", "https://bbbb.test")
		form.Insert("referrer_url_count", "2")
		form.Insert("referrer_url_0", "https://aaaa.test")
		form.Insert("referrer_url_1", "https://bbbb.test")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientCreateHandler(c)
		require.EqualError(t, err, "code=400, message=redirect_url_2 is required")
	})

	t.Run("失敗: リダイレクトURLの作成上限が超えている", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("name", "test")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile")
		form.Insert("referrer_url_count", "2")
		form.Insert("referrer_url_0", "https://aaaa.test")
		form.Insert("referrer_url_1", "https://bbbb.test")

		form.Insert("redirect_url_count", fmt.Sprint(C.ClientRedirectURLMaxCreated+2))
		for i := 0; i <= C.ClientRedirectURLMaxCreated+2; i++ {
			form.Insert(fmt.Sprintf("redirect_url_%d", i), "https://aaaa.test")
		}

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientCreateHandler(c)
		require.EqualError(t, err, "code=400, message=too many redirect urls")
	})

	t.Run("失敗: リファラーURLが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("name", "test")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile")
		form.Insert("redirect_url_count", "2")
		form.Insert("redirect_url_0", "https://aaaa.test")
		form.Insert("redirect_url_1", "https://bbbb.test")
		form.Insert("referrer_url_count", "2")
		form.Insert("referrer_url_0", "https://aaaa.test")
		form.Insert("referrer_url_1", "nyancat")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientCreateHandler(c)
		require.EqualError(t, err, "code=400, message=referrer_url `nyancat` is invalid")
	})

	t.Run("失敗: リファラーURLのForm指定が違う", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("name", "test")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile")
		form.Insert("redirect_url_count", "2")
		form.Insert("redirect_url_0", "https://aaaa.test")
		form.Insert("redirect_url_1", "https://bbbb.test")
		form.Insert("referrer_url_count", "3")
		form.Insert("referrer_url_0", "https://aaaa.test")
		form.Insert("referrer_url_1", "https://bbbb.test")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientCreateHandler(c)
		require.EqualError(t, err, "code=400, message=referrer_url_2 is required")
	})

	t.Run("失敗: リファラーURLの作成上限が超えている", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("name", "test")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile")
		form.Insert("redirect_url_count", "2")
		form.Insert("redirect_url_0", "https://aaaa.test")
		form.Insert("redirect_url_1", "https://bbbb.test")

		form.Insert("referrer_url_count", fmt.Sprint(C.ClientReferrerURLMaxCreated+2))
		for i := 0; i <= C.ClientReferrerURLMaxCreated+2; i++ {
			form.Insert(fmt.Sprintf("referrer_url_%d", i), "https://aaaa.test")
		}

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientCreateHandler(c)
		require.EqualError(t, err, "code=400, message=too many referrer urls")
	})
}

func TestClientUpdateHandler(t *testing.T) {
	t.Run("成功: クライアントを更新できる", func(t *testing.T) {})

	t.Run("成功: スコープはすべて置き換わる", func(t *testing.T) {})

	t.Run("成功: シークレットが更新できる", func(t *testing.T) {})

	t.Run("成功: 画像を更新", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが存在しない", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが不正", func(t *testing.T) {})

	t.Run("失敗: クライアントは存在するがオーナーではない", func(t *testing.T) {})

	t.Run("失敗: promptの値が不正", func(t *testing.T) {})

	t.Run("失敗: スコープの値が不正", func(t *testing.T) {})

	t.Run("失敗: リダイレクトURLが不正", func(t *testing.T) {})

	t.Run("失敗: リダイレクトURLのForm指定が違う", func(t *testing.T) {})

	t.Run("失敗: リダイレクトURLの作成上限が超えている", func(t *testing.T) {})

	t.Run("失敗: リファラーURLが不正", func(t *testing.T) {})

	t.Run("失敗: リファラーURLのForm指定が違う", func(t *testing.T) {})

	t.Run("失敗: リファラーURLの作成上限が超えている", func(t *testing.T) {})
}

func TestClientDeleteHandler(t *testing.T) {
	t.Run("成功: クライアントを削除できる", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが存在しない", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが不正", func(t *testing.T) {})

	t.Run("失敗: クライアントは存在するがオーナーではない", func(t *testing.T) {})
}

func TestClientDeleteImageHandler(t *testing.T) {
	t.Run("成功: 画像を削除できる", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが存在しない", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが不正", func(t *testing.T) {})

	t.Run("失敗: クライアントは存在するがオーナーではない", func(t *testing.T) {})
}

func TestClientAllowUserHandler(t *testing.T) {
	t.Run("成功: ルールを取得できる", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが存在しない", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが不正", func(t *testing.T) {})

	t.Run("失敗: クライアントは存在するがオーナーではない", func(t *testing.T) {})
}

func TestClientAddAllowUserHandler(t *testing.T) {
	t.Run("成功: ルールにユーザーIDを指定して追加できる", func(t *testing.T) {})

	t.Run("成功: ルールにメールアドレスのドメインを指定して追加できる", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが存在しない", func(t *testing.T) {})

	t.Run("失敗: クライアントIDが不正", func(t *testing.T) {})

	t.Run("失敗: クライアントは存在するがオーナーではない", func(t *testing.T) {})

	t.Run("失敗: user_idとemail_domainどちらも指定しない", func(t *testing.T) {})

	t.Run("失敗: user_idとemail_domainどちらも指定してしまっている", func(t *testing.T) {})
}

func TestClientDeleteAllowUserHandler(t *testing.T) {
	t.Run("成功: ルールからIDを指定して削除できる", func(t *testing.T) {})

	t.Run("失敗: idが不正", func(t *testing.T) {})

	t.Run("失敗: idが空", func(t *testing.T) {})

	t.Run("失敗: そのルールのクライアントのオーナーではない", func(t *testing.T) {})
}

func TestClientLoginUsersHandler(t *testing.T) {
	// TODO
}
