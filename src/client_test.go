package src_test

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/cateiru/cateiru-sso/src"
	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/cateiru/go-http-easy-test/v2/easy"
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestClientHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.ClientHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		return m
	})

	t.Run("成功: client_idを指定するとそのクライアントの詳細を取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, clientSecret := RegisterClient(t, ctx, &u, "openid", "profile")

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientHandler(c)
		require.NoError(t, err)

		response := src.ClientDetailResponse{}
		require.NoError(t, m.Json(&response))

		require.Equal(t, response.ClientID, clientId)

		require.Len(t, response.RedirectUrls, 0)
		require.Len(t, response.ReferrerUrls, 0)
		require.Len(t, response.Scopes, 2)
		require.Equal(t, response.ClientSecret, clientSecret)
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

	t.Run("成功: orgIdを指定するとその組織のすべてのクライアントを取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)

		RegisterClient(t, ctx, &u, "openid", "profile") // 1つだけ個人のクライアントを作る
		RegisterOrgClient(t, ctx, orgId, false, &u, "openid", "profile")
		RegisterOrgClient(t, ctx, orgId, false, &u, "openid", "profile")

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientHandler(c)
		require.NoError(t, err)

		response := []src.ClientResponse{}
		require.NoError(t, m.Json(&response))

		require.Len(t, response, 2, "orgのクライアントのみ取得できる")
	})

	t.Run("成功: clientがorgの場合、client_idを指定して取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		// u2でorgを作る
		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)
		orgId := RegisterOrg(t, ctx, &u2)

		InviteUserInOrg(t, ctx, orgId, &u, "member")

		clientId, clientSecret := RegisterOrgClient(t, ctx, orgId, false, &u, "openid", "profile")

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientHandler(c)
		require.NoError(t, err)

		response := src.ClientDetailResponse{}
		require.NoError(t, m.Json(&response))

		require.Equal(t, response.ClientID, clientId)

		require.Len(t, response.RedirectUrls, 0)
		require.Len(t, response.ReferrerUrls, 0)
		require.Len(t, response.Scopes, 2)
		require.Equal(t, response.ClientSecret, clientSecret)
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
		require.EqualError(t, err, "code=403, message=you are not owner of this client")
	})

	t.Run("失敗: orgIdが存在しない値", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/?org_id=aaaa", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientHandler(c)
		require.EqualError(t, err, "code=403, message=you are not member of this org")
	})

	t.Run("失敗: org_idをしたけどユーザーはorgのメンバーではない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)
		orgId := RegisterOrg(t, ctx, &u2)

		RegisterClient(t, ctx, &u2, "openid", "profile") // 1つだけ個人のクライアントを作る
		RegisterOrgClient(t, ctx, orgId, false, &u2, "openid", "profile")
		RegisterOrgClient(t, ctx, orgId, false, &u2, "openid", "profile")

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientHandler(c)
		require.EqualError(t, err, "code=403, message=you are not member of this org")
	})

	t.Run("失敗: org_idをしたけどユーザーはorgのメンバーだけど権限が無い", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)
		orgId := RegisterOrg(t, ctx, &u2)

		InviteUserInOrg(t, ctx, orgId, &u, "guest") // ゲストにする

		RegisterClient(t, ctx, &u2, "openid", "profile") // 1つだけ個人のクライアントを作る
		RegisterOrgClient(t, ctx, orgId, false, &u2, "openid", "profile")
		RegisterOrgClient(t, ctx, orgId, false, &u2, "openid", "profile")

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?org_id=%s", orgId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientHandler(c)
		require.EqualError(t, err, "code=403, message=you are not authority to access this organization, unique=17")
	})

	t.Run("失敗: client_idをしたけどユーザーはorgのメンバーではない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)
		orgId := RegisterOrg(t, ctx, &u2)

		clientId, _ := RegisterOrgClient(t, ctx, orgId, false, &u2, "openid", "profile")

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientHandler(c)
		require.EqualError(t, err, "code=403, message=you are not member of this org")
	})

	t.Run("失敗: client_idをしたけどユーザーはorgのメンバーだけど権限が無い", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)
		orgId := RegisterOrg(t, ctx, &u2)

		InviteUserInOrg(t, ctx, orgId, &u, "guest") // ゲストにする

		clientId, _ := RegisterOrgClient(t, ctx, orgId, false, &u2, "openid", "profile")

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientHandler(c)
		require.EqualError(t, err, "code=403, message=you are not authority to access this organization, unique=17")
	})
}

func TestClientCreateHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.ClientCreateHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
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

		return m
	})

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

		response := src.ClientDetailResponse{}
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

	t.Run("成功: org_idを指定して新規作成", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)

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

		form.Insert("org_id", orgId)
		form.Insert("org_member_only", "true")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientCreateHandler(c)
		require.NoError(t, err)

		// チェック

		response := src.ClientDetailResponse{}
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
		require.Equal(t, client.OrgID.String, orgId)
		require.Equal(t, client.OrgMemberOnly, true)

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

	t.Run("失敗: org_idが存在しない値", func(t *testing.T) {
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

		form.Insert("org_id", "aaaaa")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientCreateHandler(c)
		require.EqualError(t, err, "code=403, message=you are not member of this org")
	})

	t.Run("失敗: ユーザーはorgのメンバーではない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)
		orgId := RegisterOrg(t, ctx, &u2)

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

		form.Insert("org_id", orgId)

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientCreateHandler(c)
		require.EqualError(t, err, "code=403, message=you are not member of this org")
	})

	t.Run("失敗: ユーザーはorgの権限が無い", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)
		orgId := RegisterOrg(t, ctx, &u2)

		InviteUserInOrg(t, ctx, orgId, &u, "guest")

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

		form.Insert("org_id", orgId)

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientCreateHandler(c)
		require.EqualError(t, err, "code=403, message=you are not member of this org")
	})

	t.Run("失敗: orgの作成上限を超えている", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		orgId := RegisterOrg(t, ctx, &u)

		for i := 0; i <= C.OrgClientMaxCreated; i++ {
			RegisterOrgClient(t, ctx, orgId, false, &u)
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

		form.Insert("org_id", orgId)

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientCreateHandler(c)
		require.EqualError(t, err, "code=400, message=too many clients")
	})
}

func TestClientUpdateHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.ClientUpdateHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		clientId, _ := RegisterClient(t, ctx, u)

		form := easy.NewMultipart()
		form.Insert("client_id", clientId)

		form.Insert("name", "new!!! name")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile email")
		form.Insert("redirect_url_count", "2")
		form.Insert("redirect_url_0", "https://aaaa.test")
		form.Insert("redirect_url_1", "https://bbbb.test")
		form.Insert("referrer_url_count", "2")
		form.Insert("referrer_url_0", "https://aaaa.test")
		form.Insert("referrer_url_1", "https://bbbb.test")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		return m
	})

	t.Run("成功: クライアントを更新できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, clientSecret := RegisterClient(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("client_id", clientId)

		form.Insert("name", "new!!! name")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile email")
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

		err = h.ClientUpdateHandler(c)
		require.NoError(t, err)

		response := src.ClientDetailResponse{}
		require.NoError(t, m.Json(&response))

		require.Equal(t, "new!!! name", response.Name)
		require.Equal(t, response.ClientSecret, clientSecret)

		// スコープ
		scopes, err := models.ClientScopes(
			models.ClientScopeWhere.ClientID.EQ(clientId),
		).Count(ctx, h.DB)
		require.NoError(t, err)
		require.Equal(t, 3, int(scopes))

		// リダイレクトURL
		redirectUrls, err := models.ClientRedirects(
			models.ClientRedirectWhere.ClientID.EQ(clientId),
		).Count(ctx, h.DB)
		require.NoError(t, err)
		require.Equal(t, 2, int(redirectUrls))

		// リファラーURL
		referrerUrls, err := models.ClientReferrers(
			models.ClientReferrerWhere.ClientID.EQ(clientId),
		).Count(ctx, h.DB)
		require.NoError(t, err)
		require.Equal(t, 2, int(referrerUrls))
	})

	t.Run("成功: シークレットが更新できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, clientSecret := RegisterClient(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("client_id", clientId)

		form.Insert("name", "new!!! name")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile email")
		form.Insert("redirect_url_count", "2")
		form.Insert("redirect_url_0", "https://aaaa.test")
		form.Insert("redirect_url_1", "https://bbbb.test")
		form.Insert("referrer_url_count", "2")
		form.Insert("referrer_url_0", "https://aaaa.test")
		form.Insert("referrer_url_1", "https://bbbb.test")

		form.Insert("update_secret", "true")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientUpdateHandler(c)
		require.NoError(t, err)

		response := src.ClientResponse{}
		require.NoError(t, m.Json(&response))

		client, err := models.Clients(
			models.ClientWhere.ClientID.EQ(clientId),
		).One(ctx, h.DB)
		require.NoError(t, err)

		require.NotEqual(t, clientSecret, client.ClientSecret)
	})

	t.Run("成功: 画像を更新", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("client_id", clientId)

		form.Insert("name", "new!!! name")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile email")
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

		err = h.ClientUpdateHandler(c)
		require.NoError(t, err)

		// とりあえずエラーにならなかったらOKとしておく
	})

	t.Run("成功: ユーザーはorgに入っているので更新ができる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)
		orgId := RegisterOrg(t, ctx, &u2)

		// uはorgのメンバー
		InviteUserInOrg(t, ctx, orgId, &u, "member")

		clientId, clientSecret := RegisterOrgClient(t, ctx, orgId, false, &u2)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("client_id", clientId)

		form.Insert("name", "new!!! name")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile email")
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

		err = h.ClientUpdateHandler(c)
		require.NoError(t, err)

		response := src.ClientDetailResponse{}
		require.NoError(t, m.Json(&response))

		require.Equal(t, "new!!! name", response.Name)
		require.Equal(t, response.ClientSecret, clientSecret)

		// スコープ
		scopes, err := models.ClientScopes(
			models.ClientScopeWhere.ClientID.EQ(clientId),
		).Count(ctx, h.DB)
		require.NoError(t, err)
		require.Equal(t, 3, int(scopes))

		// リダイレクトURL
		redirectUrls, err := models.ClientRedirects(
			models.ClientRedirectWhere.ClientID.EQ(clientId),
		).Count(ctx, h.DB)
		require.NoError(t, err)
		require.Equal(t, 2, int(redirectUrls))

		// リファラーURL
		referrerUrls, err := models.ClientReferrers(
			models.ClientReferrerWhere.ClientID.EQ(clientId),
		).Count(ctx, h.DB)
		require.NoError(t, err)
		require.Equal(t, 2, int(referrerUrls))
	})

	t.Run("失敗: クライアントIDが存在しない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()

		form.Insert("name", "new!!! name")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile email")
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

		err = h.ClientUpdateHandler(c)
		require.EqualError(t, err, "code=400, message=client_id is required")
	})

	t.Run("失敗: クライアントIDが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("client_id", "nyancat")

		form.Insert("name", "new!!! name")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile email")
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

		err = h.ClientUpdateHandler(c)
		require.EqualError(t, err, "code=404, message=client not found")
	})

	t.Run("失敗: クライアントは存在するがオーナーではない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		clientId, _ := RegisterClient(t, ctx, &u2)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("client_id", clientId)

		form.Insert("name", "new!!! name")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile email")
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

		err = h.ClientUpdateHandler(c)
		require.EqualError(t, err, "code=403, message=you are not owner of this client")
	})

	t.Run("失敗: orgに入っていない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)
		orgId := RegisterOrg(t, ctx, &u2)

		clientId, _ := RegisterOrgClient(t, ctx, orgId, false, &u2)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("client_id", clientId)

		form.Insert("name", "new!!! name")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile email")
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

		err = h.ClientUpdateHandler(c)
		require.EqualError(t, err, "code=403, message=you are not member of this org")
	})

	t.Run("失敗: orgに入っているが権限が無い", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)
		orgId := RegisterOrg(t, ctx, &u2)

		// uはorgのゲスト
		InviteUserInOrg(t, ctx, orgId, &u, "guest")

		clientId, _ := RegisterOrgClient(t, ctx, orgId, false, &u2)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("client_id", clientId)

		form.Insert("name", "new!!! name")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile email")
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

		err = h.ClientUpdateHandler(c)
		require.EqualError(t, err, "code=403, message=you are not authority to access this organization, unique=17")
	})

	t.Run("失敗: promptの値が不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("client_id", clientId)

		form.Insert("name", "new!!! name")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile email")
		form.Insert("redirect_url_count", "2")
		form.Insert("redirect_url_0", "https://aaaa.test")
		form.Insert("redirect_url_1", "https://bbbb.test")
		form.Insert("referrer_url_count", "2")
		form.Insert("referrer_url_0", "https://aaaa.test")
		form.Insert("referrer_url_1", "https://bbbb.test")

		form.Insert("prompt", "aaaa")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientUpdateHandler(c)
		require.EqualError(t, err, "code=400, message=prompt is invalid")
	})

	t.Run("失敗: スコープの値が不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("client_id", clientId)

		form.Insert("name", "new!!! name")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile aaaaaa")
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

		err = h.ClientUpdateHandler(c)
		require.EqualError(t, err, "code=400, message=scope is invalid")
	})

	t.Run("失敗: リダイレクトURLが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("client_id", clientId)

		form.Insert("name", "new!!! name")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile email")
		form.Insert("redirect_url_count", "2")
		form.Insert("redirect_url_0", "nyan")
		form.Insert("redirect_url_1", "https://bbbb.test")
		form.Insert("referrer_url_count", "2")
		form.Insert("referrer_url_0", "https://aaaa.test")
		form.Insert("referrer_url_1", "https://bbbb.test")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientUpdateHandler(c)
		require.EqualError(t, err, "code=400, message=referrer_url `nyan` is invalid")
	})

	t.Run("失敗: リダイレクトURLのForm指定が違う", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("client_id", clientId)

		form.Insert("name", "new!!! name")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile email")
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

		err = h.ClientUpdateHandler(c)
		require.EqualError(t, err, "code=400, message=redirect_url_2 is required")
	})

	t.Run("失敗: リダイレクトURLの作成上限が超えている", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("client_id", clientId)

		form.Insert("name", "new!!! name")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile email")
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

		err = h.ClientUpdateHandler(c)
		require.EqualError(t, err, "code=400, message=too many redirect urls")
	})

	t.Run("失敗: リファラーURLが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("client_id", clientId)

		form.Insert("name", "new!!! name")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile email")
		form.Insert("redirect_url_count", "2")
		form.Insert("redirect_url_0", "https://aaaa.test")
		form.Insert("redirect_url_1", "https://bbbb.test")
		form.Insert("referrer_url_count", "2")
		form.Insert("referrer_url_0", "https://aaaa.test")
		form.Insert("referrer_url_1", "aaaaa")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientUpdateHandler(c)
		require.EqualError(t, err, "code=400, message=referrer_url `aaaaa` is invalid")
	})

	t.Run("失敗: リファラーURLのForm指定が違う", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("client_id", clientId)

		form.Insert("name", "new!!! name")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile email")
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

		err = h.ClientUpdateHandler(c)
		require.EqualError(t, err, "code=400, message=referrer_url_2 is required")
	})

	t.Run("失敗: リファラーURLの作成上限が超えている", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("client_id", clientId)

		form.Insert("name", "new!!! name")
		form.Insert("is_allow", "false")
		form.Insert("scopes", "openid profile email")
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

		err = h.ClientUpdateHandler(c)
		require.EqualError(t, err, "code=400, message=too many referrer urls")
	})
}

func TestClientDeleteHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.ClientDeleteHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		clientId, _ := RegisterClient(t, ctx, u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodDelete, "")
		require.NoError(t, err)

		return m
	})

	t.Run("成功: クライアントを削除できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientDeleteHandler(c)
		require.NoError(t, err)

		clientExists, err := models.Clients(
			models.ClientWhere.ClientID.EQ(clientId),
		).Exists(ctx, h.DB)
		require.NoError(t, err)
		require.False(t, clientExists)

		scopes, err := models.ClientScopes(
			models.ClientScopeWhere.ClientID.EQ(clientId),
		).Exists(ctx, h.DB)
		require.NoError(t, err)
		require.False(t, scopes)

		redirects, err := models.ClientRedirects(
			models.ClientRedirectWhere.ClientID.EQ(clientId),
		).Exists(ctx, h.DB)
		require.NoError(t, err)
		require.False(t, redirects)

		referrers, err := models.ClientReferrers(
			models.ClientReferrerWhere.ClientID.EQ(clientId),
		).Exists(ctx, h.DB)
		require.NoError(t, err)
		require.False(t, referrers)

		allows, err := models.ClientAllowRules(
			models.ClientAllowRuleWhere.ClientID.EQ(clientId),
		).Exists(ctx, h.DB)
		require.NoError(t, err)
		require.False(t, allows)
	})

	t.Run("成功: ユーザーはorgに入っているので削除ができる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		orgId := RegisterOrg(t, ctx, &u2)

		// uはorgのメンバー
		InviteUserInOrg(t, ctx, orgId, &u, "member")

		clientId, _ := RegisterOrgClient(t, ctx, orgId, false, &u2)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientDeleteHandler(c)
		require.NoError(t, err)

		clientExists, err := models.Clients(
			models.ClientWhere.ClientID.EQ(clientId),
		).Exists(ctx, h.DB)
		require.NoError(t, err)
		require.False(t, clientExists)

		scopes, err := models.ClientScopes(
			models.ClientScopeWhere.ClientID.EQ(clientId),
		).Exists(ctx, h.DB)
		require.NoError(t, err)
		require.False(t, scopes)

		redirects, err := models.ClientRedirects(
			models.ClientRedirectWhere.ClientID.EQ(clientId),
		).Exists(ctx, h.DB)
		require.NoError(t, err)
		require.False(t, redirects)

		referrers, err := models.ClientReferrers(
			models.ClientReferrerWhere.ClientID.EQ(clientId),
		).Exists(ctx, h.DB)
		require.NoError(t, err)
		require.False(t, referrers)

		allows, err := models.ClientAllowRules(
			models.ClientAllowRuleWhere.ClientID.EQ(clientId),
		).Exists(ctx, h.DB)
		require.NoError(t, err)
		require.False(t, allows)
	})

	t.Run("失敗: クライアントIDが存在しない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientDeleteHandler(c)
		require.EqualError(t, err, "code=400, message=client_id is required")
	})

	t.Run("失敗: クライアントIDが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/?client_id=invalid", http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientDeleteHandler(c)
		require.EqualError(t, err, "code=404, message=client not found")
	})

	t.Run("失敗: クライアントは存在するがオーナーではない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		cookie := RegisterSession(t, ctx, &u2)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientDeleteHandler(c)
		require.EqualError(t, err, "code=403, message=you are not owner of this client")
	})

	t.Run("失敗: orgに入っていないので削除できない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		orgId := RegisterOrg(t, ctx, &u2)

		clientId, _ := RegisterOrgClient(t, ctx, orgId, false, &u2)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientDeleteHandler(c)
		require.EqualError(t, err, "code=403, message=you are not member of this org")
	})

	t.Run("失敗: orgに入っているが権限が無いので削除できない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		orgId := RegisterOrg(t, ctx, &u2)

		InviteUserInOrg(t, ctx, orgId, &u, "guest")

		clientId, _ := RegisterOrgClient(t, ctx, orgId, false, &u2)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientDeleteHandler(c)
		require.EqualError(t, err, "code=403, message=you are not authority to access this organization, unique=17")
	})
}

func TestClientDeleteImageHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.ClientDeleteImageHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		clientId, _ := RegisterClient(t, ctx, u)

		path := filepath.Join("client_icon", clientId)
		url := &url.URL{
			Scheme: C.CDNHost.Scheme,
			Host:   C.CDNHost.Host,
			Path:   path,
		}

		client, err := models.Clients(
			models.ClientWhere.ClientID.EQ(clientId),
		).One(ctx, DB)
		require.NoError(t, err)

		client.Image = null.NewString(url.String(), true)

		_, err = client.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodDelete, "")
		require.NoError(t, err)

		return m
	})

	t.Run("成功: 画像を削除できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u)

		path := filepath.Join("client_icon", clientId)
		url := &url.URL{
			Scheme: C.CDNHost.Scheme,
			Host:   C.CDNHost.Host,
			Path:   path,
		}

		client, err := models.Clients(
			models.ClientWhere.ClientID.EQ(clientId),
		).One(ctx, DB)
		require.NoError(t, err)

		client.Image = null.NewString(url.String(), true)

		_, err = client.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientDeleteImageHandler(c)
		require.NoError(t, err)

		client, err = models.Clients(
			models.ClientWhere.ClientID.EQ(clientId),
		).One(ctx, DB)
		require.NoError(t, err)

		require.False(t, client.Image.Valid)
	})

	t.Run("成功: ユーザーはorgに入っているので削除ができる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		orgId := RegisterOrg(t, ctx, &u2)

		// uはorgのメンバー
		InviteUserInOrg(t, ctx, orgId, &u, "member")

		clientId, _ := RegisterOrgClient(t, ctx, orgId, false, &u2)

		path := filepath.Join("client_icon", clientId)
		url := &url.URL{
			Scheme: C.CDNHost.Scheme,
			Host:   C.CDNHost.Host,
			Path:   path,
		}

		client, err := models.Clients(
			models.ClientWhere.ClientID.EQ(clientId),
		).One(ctx, DB)
		require.NoError(t, err)

		client.Image = null.NewString(url.String(), true)

		_, err = client.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientDeleteImageHandler(c)
		require.NoError(t, err)

		client, err = models.Clients(
			models.ClientWhere.ClientID.EQ(clientId),
		).One(ctx, DB)
		require.NoError(t, err)

		require.False(t, client.Image.Valid)
	})

	t.Run("失敗: そもそも画像が設定されていない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientDeleteImageHandler(c)
		require.EqualError(t, err, "code=404, message=image is not set")
	})

	t.Run("失敗: クライアントIDが存在しない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientDeleteImageHandler(c)
		require.EqualError(t, err, "code=400, message=client_id is required")
	})

	t.Run("失敗: クライアントIDが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/?client_id=aaaaaa", http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientDeleteImageHandler(c)
		require.EqualError(t, err, "code=404, message=client not found")
	})

	t.Run("失敗: クライアントは存在するがオーナーではない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		cookie := RegisterSession(t, ctx, &u2)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientDeleteImageHandler(c)
		require.EqualError(t, err, "code=403, message=you are not owner of this client")
	})

	t.Run("失敗: orgに入っていないので削除できない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		orgId := RegisterOrg(t, ctx, &u2)

		clientId, _ := RegisterOrgClient(t, ctx, orgId, false, &u2)

		path := filepath.Join("client_icon", clientId)
		url := &url.URL{
			Scheme: C.CDNHost.Scheme,
			Host:   C.CDNHost.Host,
			Path:   path,
		}

		client, err := models.Clients(
			models.ClientWhere.ClientID.EQ(clientId),
		).One(ctx, DB)
		require.NoError(t, err)

		client.Image = null.NewString(url.String(), true)

		_, err = client.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientDeleteImageHandler(c)
		require.EqualError(t, err, "code=403, message=you are not member of this org")
	})

	t.Run("失敗: orgに入っているが権限が無いので削除できない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		orgId := RegisterOrg(t, ctx, &u2)

		InviteUserInOrg(t, ctx, orgId, &u, "guest")

		clientId, _ := RegisterOrgClient(t, ctx, orgId, false, &u2)

		path := filepath.Join("client_icon", clientId)
		url := &url.URL{
			Scheme: C.CDNHost.Scheme,
			Host:   C.CDNHost.Host,
			Path:   path,
		}

		client, err := models.Clients(
			models.ClientWhere.ClientID.EQ(clientId),
		).One(ctx, DB)
		require.NoError(t, err)

		client.Image = null.NewString(url.String(), true)

		_, err = client.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodDelete, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientDeleteImageHandler(c)
		require.EqualError(t, err, "code=403, message=you are not authority to access this organization, unique=17")
	})
}

func TestClientAllowUserHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.ClientAllowUserHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		clientId, _ := RegisterClient(t, ctx, u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodDelete, "")
		require.NoError(t, err)

		return m
	})

	t.Run("成功: ルールを取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u)

		for i := 0; i < 2; i++ {
			RegisterAllowRules(t, ctx, clientId, false, fmt.Sprintf("%daaa.test", i))
		}

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientAllowUserHandler(c)
		require.NoError(t, err)

		response := []src.ClientAllowUserRuleResponse{}
		require.NoError(t, m.Json(&response))

		require.Len(t, response, 2)
	})

	t.Run("成功: ユーザーはorgに入っているの取得できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		orgId := RegisterOrg(t, ctx, &u2)

		// uはorgのメンバー
		InviteUserInOrg(t, ctx, orgId, &u, "member")

		clientId, _ := RegisterOrgClient(t, ctx, orgId, false, &u2)

		for i := 0; i < 2; i++ {
			RegisterAllowRules(t, ctx, clientId, false, fmt.Sprintf("%daaa.test", i))
		}

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientAllowUserHandler(c)
		require.NoError(t, err)

		response := []src.ClientAllowUserRuleResponse{}
		require.NoError(t, m.Json(&response))

		require.Len(t, response, 2)
	})

	t.Run("失敗: クライアントIDが存在しない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientAllowUserHandler(c)
		require.EqualError(t, err, "code=400, message=client_id is required")
	})

	t.Run("失敗: クライアントIDが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/?client_id=invalid", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientAllowUserHandler(c)
		require.EqualError(t, err, "code=404, message=client not found")
	})

	t.Run("失敗: クライアントは存在するがオーナーではない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		cookie := RegisterSession(t, ctx, &u2)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientAllowUserHandler(c)
		require.EqualError(t, err, "code=403, message=you are not owner of this client")
	})

	t.Run("失敗: orgに入っていないので取得できない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		orgId := RegisterOrg(t, ctx, &u2)

		clientId, _ := RegisterOrgClient(t, ctx, orgId, false, &u2)

		for i := 0; i < 2; i++ {
			RegisterAllowRules(t, ctx, clientId, false, fmt.Sprintf("%daaa.test", i))
		}

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientAllowUserHandler(c)
		require.EqualError(t, err, "code=403, message=you are not member of this org")
	})

	t.Run("失敗: orgに入っているが権限が無いので取得できない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		orgId := RegisterOrg(t, ctx, &u2)

		InviteUserInOrg(t, ctx, orgId, &u, "guest")

		clientId, _ := RegisterOrgClient(t, ctx, orgId, false, &u2)

		for i := 0; i < 2; i++ {
			RegisterAllowRules(t, ctx, clientId, false, fmt.Sprintf("%daaa.test", i))
		}

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?client_id=%s", clientId), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientAllowUserHandler(c)
		require.EqualError(t, err, "code=403, message=you are not authority to access this organization, unique=17")
	})
}

func TestClientAddAllowUserHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.ClientAddAllowUserHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		clientId, _ := RegisterClient(t, ctx, u)

		form := easy.NewMultipart()
		form.Insert("client_id", clientId)
		form.Insert("user_id", "test")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)

		return m
	})

	t.Run("成功: ルールにユーザーIDを指定して追加できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("client_id", clientId)
		form.Insert("user_id", "test")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientAddAllowUserHandler(c)
		require.NoError(t, err)

		rule, err := models.ClientAllowRules(
			models.ClientAllowRuleWhere.ClientID.EQ(clientId),
		).One(ctx, h.DB)
		require.NoError(t, err)

		require.Equal(t, rule.UserID.String, "test")
		require.False(t, rule.EmailDomain.Valid)
	})

	t.Run("成功: ルールにメールアドレスのドメインを指定して追加できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("client_id", clientId)
		form.Insert("email_domain", "cateiru.test")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientAddAllowUserHandler(c)
		require.NoError(t, err)

		rule, err := models.ClientAllowRules(
			models.ClientAllowRuleWhere.ClientID.EQ(clientId),
		).One(ctx, h.DB)
		require.NoError(t, err)

		require.False(t, rule.UserID.Valid)
		require.Equal(t, rule.EmailDomain.String, "cateiru.test")
	})

	t.Run("成功: ユーザーはorgに入っているの追加できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		orgId := RegisterOrg(t, ctx, &u2)

		// uはorgのメンバー
		InviteUserInOrg(t, ctx, orgId, &u, "member")

		clientId, _ := RegisterOrgClient(t, ctx, orgId, false, &u2)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("client_id", clientId)
		form.Insert("user_id", "test")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientAddAllowUserHandler(c)
		require.NoError(t, err)

		rule, err := models.ClientAllowRules(
			models.ClientAllowRuleWhere.ClientID.EQ(clientId),
		).One(ctx, h.DB)
		require.NoError(t, err)

		require.Equal(t, rule.UserID.String, "test")
		require.False(t, rule.EmailDomain.Valid)
	})

	t.Run("失敗: クライアントIDが存在しない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("email_domain", "cateiru.test")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientAddAllowUserHandler(c)
		require.EqualError(t, err, "code=400, message=client_id is required")
	})

	t.Run("失敗: クライアントIDが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("client_id", "invalid")
		form.Insert("email_domain", "cateiru.test")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientAddAllowUserHandler(c)
		require.EqualError(t, err, "code=404, message=client not found")
	})

	t.Run("失敗: クライアントは存在するがオーナーではない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		cookie := RegisterSession(t, ctx, &u2)

		form := easy.NewMultipart()
		form.Insert("client_id", clientId)
		form.Insert("email_domain", "cateiru.test")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientAddAllowUserHandler(c)
		require.EqualError(t, err, "code=403, message=you are not owner of this client")
	})

	t.Run("失敗: user_idとemail_domainどちらも指定しない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("client_id", clientId)

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientAddAllowUserHandler(c)
		require.EqualError(t, err, "code=400, message=user_id or email_domain is required")
	})

	t.Run("失敗: user_idとemail_domainどちらも指定してしまっている", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("client_id", clientId)
		form.Insert("user_id", "test")
		form.Insert("email_domain", "cateiru.test")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientAddAllowUserHandler(c)
		require.EqualError(t, err, "code=400, message=user_id and email_domain cannot be set at the same time")
	})

	t.Run("失敗: orgに入っていないので追加できない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		orgId := RegisterOrg(t, ctx, &u2)

		clientId, _ := RegisterOrgClient(t, ctx, orgId, false, &u2)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("client_id", clientId)
		form.Insert("user_id", "test")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientAddAllowUserHandler(c)
		require.EqualError(t, err, "code=403, message=you are not member of this org")
	})

	t.Run("失敗: orgに入っているが権限が無いので追加できない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		orgId := RegisterOrg(t, ctx, &u2)

		// uはorgのメンバー
		InviteUserInOrg(t, ctx, orgId, &u, "guest")

		clientId, _ := RegisterOrgClient(t, ctx, orgId, false, &u2)

		cookie := RegisterSession(t, ctx, &u)

		form := easy.NewMultipart()
		form.Insert("client_id", clientId)
		form.Insert("user_id", "test")

		m, err := easy.NewFormData("/", http.MethodPost, form)
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientAddAllowUserHandler(c)
		require.EqualError(t, err, "code=403, message=you are not authority to access this organization, unique=17")
	})
}

func TestClientDeleteAllowUserHandler(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	SessionTest(t, h.ClientDeleteAllowUserHandler, func(ctx context.Context, u *models.User) *easy.MockHandler {
		clientId, _ := RegisterClient(t, ctx, u)

		RegisterAllowRules(t, ctx, clientId, false, "cateiru.test")
		rule, err := models.ClientAllowRules(
			models.ClientAllowRuleWhere.ClientID.EQ(clientId),
		).One(ctx, h.DB)
		require.NoError(t, err)

		m, err := easy.NewMock(fmt.Sprintf("/?id=%d", rule.ID), http.MethodGet, "")
		require.NoError(t, err)

		return m
	})

	t.Run("成功: ルールからIDを指定して削除できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u)

		RegisterAllowRules(t, ctx, clientId, false, "cateiru.test")
		rule, err := models.ClientAllowRules(
			models.ClientAllowRuleWhere.ClientID.EQ(clientId),
		).One(ctx, h.DB)
		require.NoError(t, err)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?id=%d", rule.ID), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientDeleteAllowUserHandler(c)
		require.NoError(t, err)

		existRule, err := models.ClientAllowRules(
			models.ClientAllowRuleWhere.ID.EQ(rule.ID),
		).Exists(ctx, h.DB)
		require.NoError(t, err)
		require.False(t, existRule)
	})

	t.Run("成功: ユーザーはorgに入っているので削除できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		orgId := RegisterOrg(t, ctx, &u2)

		// uはorgのメンバー
		InviteUserInOrg(t, ctx, orgId, &u, "member")

		clientId, _ := RegisterOrgClient(t, ctx, orgId, false, &u2)

		RegisterAllowRules(t, ctx, clientId, false, "cateiru.test")
		rule, err := models.ClientAllowRules(
			models.ClientAllowRuleWhere.ClientID.EQ(clientId),
		).One(ctx, h.DB)
		require.NoError(t, err)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?id=%d", rule.ID), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientDeleteAllowUserHandler(c)
		require.NoError(t, err)

		existRule, err := models.ClientAllowRules(
			models.ClientAllowRuleWhere.ID.EQ(rule.ID),
		).Exists(ctx, h.DB)
		require.NoError(t, err)
		require.False(t, existRule)
	})

	t.Run("失敗: idが不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/?id=invalid", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientDeleteAllowUserHandler(c)
		require.EqualError(t, err, "code=400, message=id is invalid")
	})

	t.Run("失敗: idが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientDeleteAllowUserHandler(c)
		require.EqualError(t, err, "code=400, message=id is required")
	})

	t.Run("失敗: そのルールのクライアントのオーナーではない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		clientId, _ := RegisterClient(t, ctx, &u)

		RegisterAllowRules(t, ctx, clientId, false, "cateiru.test")
		rule, err := models.ClientAllowRules(
			models.ClientAllowRuleWhere.ClientID.EQ(clientId),
		).One(ctx, h.DB)
		require.NoError(t, err)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		cookie := RegisterSession(t, ctx, &u2)

		m, err := easy.NewMock(fmt.Sprintf("/?id=%d", rule.ID), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientDeleteAllowUserHandler(c)
		require.EqualError(t, err, "code=403, message=you are not owner of this client")
	})

	t.Run("失敗: orgに入っていないので削除できない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		orgId := RegisterOrg(t, ctx, &u2)

		clientId, _ := RegisterOrgClient(t, ctx, orgId, false, &u2)

		RegisterAllowRules(t, ctx, clientId, false, "cateiru.test")
		rule, err := models.ClientAllowRules(
			models.ClientAllowRuleWhere.ClientID.EQ(clientId),
		).One(ctx, h.DB)
		require.NoError(t, err)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?id=%d", rule.ID), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientDeleteAllowUserHandler(c)
		require.EqualError(t, err, "code=403, message=you are not member of this org")
	})

	t.Run("失敗: orgに入っているが権限が無いので削除できない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		email2 := RandomEmail(t)
		u2 := RegisterUser(t, ctx, email2)

		orgId := RegisterOrg(t, ctx, &u2)

		// uはorgのメンバー
		InviteUserInOrg(t, ctx, orgId, &u, "guest")

		clientId, _ := RegisterOrgClient(t, ctx, orgId, false, &u2)

		RegisterAllowRules(t, ctx, clientId, false, "cateiru.test")
		rule, err := models.ClientAllowRules(
			models.ClientAllowRuleWhere.ClientID.EQ(clientId),
		).One(ctx, h.DB)
		require.NoError(t, err)

		cookie := RegisterSession(t, ctx, &u)

		m, err := easy.NewMock(fmt.Sprintf("/?id=%d", rule.ID), http.MethodGet, "")
		require.NoError(t, err)
		m.Cookie(cookie)

		c := m.Echo()

		err = h.ClientDeleteAllowUserHandler(c)
		require.EqualError(t, err, "code=403, message=you are not authority to access this organization, unique=17")
	})
}

func TestClientLoginUsersHandler(t *testing.T) {
	// TODO
}
