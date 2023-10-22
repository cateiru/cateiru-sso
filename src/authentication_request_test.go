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
	"github.com/stretchr/testify/require"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

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

		require.Equal(t, *response, src.PreviewResponse{
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

		require.Equal(t, *response, src.PreviewResponse{
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
