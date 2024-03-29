package src_test

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
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

func TestClientAuthentication(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	t.Run("成功: Basic認証", func(t *testing.T) {
		client := RegisterClient(t, ctx, nil)

		base64Value := base64.StdEncoding.EncodeToString([]byte(client.ClientID + ":" + client.ClientSecret))

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		m.R.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64Value))

		c := m.Echo()

		returnClient, err := h.ClientAuthentication(ctx, c)
		require.NoError(t, err)

		require.Equal(t, client.ClientID, returnClient.ClientID)
	})

	t.Run("成功: POST", func(t *testing.T) {
		client := RegisterClient(t, ctx, nil)

		pathParam := fmt.Sprintf("/?client_id=%s&client_secret=%s", client.ClientID, client.ClientSecret)

		m, err := easy.NewMock(pathParam, http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()

		returnClient, err := h.ClientAuthentication(ctx, c)
		require.NoError(t, err)

		require.Equal(t, client.ClientID, returnClient.ClientID)
	})

	t.Run("失敗: どの認証も無い", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()

		_, err = h.ClientAuthentication(ctx, c)
		require.EqualError(t, err, "code=401, error=invalid_client, message=Invalid client authentication")

		wwwAuthenticate := m.Response().Header.Get("WWW-Authenticate")
		require.Equal(t, wwwAuthenticate, "Basic")
	})

	t.Run("失敗: Basic認証でAuthorizationの形式が不正", func(t *testing.T) {
		client := RegisterClient(t, ctx, nil)

		base64Value := base64.StdEncoding.EncodeToString([]byte(client.ClientID + ":" + client.ClientSecret))

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		m.R.Header.Set("Authorization", fmt.Sprintf("Basicaaa %s", base64Value))

		c := m.Echo()

		_, err = h.ClientAuthentication(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request, message=Invalid Authorization Header")
	})

	t.Run("失敗: Basic認証でBase64デコードに失敗", func(t *testing.T) {
		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		m.R.Header.Set("Authorization", "Basic hogehoge")

		c := m.Echo()

		_, err = h.ClientAuthentication(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request, message=Invalid Authorization Header")

	})

	t.Run("失敗: Basic認証でクライアントが存在しない", func(t *testing.T) {
		client := RegisterClient(t, ctx, nil)

		base64Value := base64.StdEncoding.EncodeToString([]byte("invalid_client" + ":" + client.ClientSecret))

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		m.R.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64Value))

		c := m.Echo()

		_, err = h.ClientAuthentication(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request, message=Invalid client_id")
	})

	t.Run("失敗: POSTでクライアントが存在しない", func(t *testing.T) {
		client := RegisterClient(t, ctx, nil)

		pathParam := fmt.Sprintf("/?client_id=%s&client_secret=%s", "invalid_client", client.ClientSecret)

		m, err := easy.NewMock(pathParam, http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()

		_, err = h.ClientAuthentication(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request, message=Invalid client_id")
	})

	t.Run("失敗: Basic認証でクライアントシークレットが不正", func(t *testing.T) {
		client := RegisterClient(t, ctx, nil)

		base64Value := base64.StdEncoding.EncodeToString([]byte(client.ClientID + ":" + "invalid_client_secret"))

		m, err := easy.NewMock("/", http.MethodGet, "")
		require.NoError(t, err)

		m.R.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64Value))

		c := m.Echo()

		_, err = h.ClientAuthentication(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request, message=Invalid client_secret")
	})

	t.Run("失敗: POSTでクライアントシークレットが不正", func(t *testing.T) {
		client := RegisterClient(t, ctx, nil)

		pathParam := fmt.Sprintf("/?client_id=%s&client_secret=%s", client.ClientID, "invalid_client_secret")

		m, err := easy.NewMock(pathParam, http.MethodGet, "")
		require.NoError(t, err)

		c := m.Echo()

		_, err = h.ClientAuthentication(ctx, c)
		require.EqualError(t, err, "code=400, error=invalid_request, message=Invalid client_secret")
	})
}

func TestTokenEndpointAuthorizationCode(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	t.Run("成功: レスポンスを受け取れる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		client := RegisterClient(t, ctx, nil, "openid", "email", "profile")

		redirectUri, err := url.Parse("https://example.test/hogehoge")
		require.NoError(t, err)
		redirect := models.ClientRedirect{
			ClientID: client.ClientID,
			Host:     redirectUri.Host,
			URL:      redirectUri.String(),
		}
		err = redirect.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		oauthSession := RegisterOauthSession(t, ctx, client.ClientID, &u)

		query := url.Values{}
		query.Set("code", oauthSession.Code)
		query.Set("redirect_uri", redirectUri.String())

		m, err := easy.NewURLEncoded("/", http.MethodPost, query)
		require.NoError(t, err)

		c := m.Echo()

		err = h.TokenEndpointAuthorizationCode(ctx, c, client)
		require.NoError(t, err)

		// oauthSession は削除されている
		existOauthSession, err := models.OauthSessions(
			models.OauthSessionWhere.Code.EQ(oauthSession.Code),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.False(t, existOauthSession)

		response := src.TokenEndpointResponse{}
		require.NoError(t, m.Json(&response))

		require.Equal(t, response.TokenType, "Bearer")
		require.Equal(t, response.ExpiresIn, int64(h.C.OAuthAccessTokenPeriod)/10000000)

		// アクセストークン有効確認
		accessToken, err := models.ClientSessions(
			models.ClientSessionWhere.ID.EQ(response.AccessToken),
		).One(ctx, DB)
		require.NoError(t, err)
		require.Equal(t, accessToken.ClientID, client.ClientID)
		require.Equal(t, accessToken.UserID, u.ID)
		require.True(t, accessToken.Period.After(time.Now()))

		// リフレッシュトークン有効確認
		refreshToken, err := models.ClientRefreshes(
			models.ClientRefreshWhere.ID.EQ(response.RefreshToken),
		).One(ctx, DB)
		require.NoError(t, err)
		require.Equal(t, refreshToken.ClientID, client.ClientID)
		require.Equal(t, refreshToken.UserID, u.ID)
		require.Equal(t, refreshToken.SessionID, accessToken.ID)
		require.True(t, refreshToken.Period.After(time.Now()))

		// IDToken の検証
		idToken := response.IDToken
		require.NotEmpty(t, idToken)

		claims := src.AuthorizationCodeFlowClaims{}
		token := DecodeJWT(t, idToken, &claims)
		require.True(t, token.Valid)

		require.Equal(t, claims.Iss, h.C.SiteHost.String())
		require.Equal(t, claims.Sub, u.ID)
		require.Equal(t, claims.Nonce, oauthSession.Nonce.String)
		require.Equal(t, claims.StandardClaims.PreferredUsername, u.UserName)
	})

	t.Run("成功: スコープが無いとその情報は取得できない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		client := RegisterClient(t, ctx, nil, "openid") // profile と email が無い

		redirectUri, err := url.Parse("https://example.test/hogehoge")
		require.NoError(t, err)
		redirect := models.ClientRedirect{
			ClientID: client.ClientID,
			Host:     redirectUri.Host,
			URL:      redirectUri.String(),
		}
		err = redirect.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		oauthSession := RegisterOauthSession(t, ctx, client.ClientID, &u)

		query := url.Values{}
		query.Set("code", oauthSession.Code)
		query.Set("redirect_uri", redirectUri.String())

		m, err := easy.NewURLEncoded("/", http.MethodPost, query)
		require.NoError(t, err)

		c := m.Echo()

		err = h.TokenEndpointAuthorizationCode(ctx, c, client)
		require.NoError(t, err)

		response := src.TokenEndpointResponse{}
		require.NoError(t, m.Json(&response))

		// IDToken の検証
		idToken := response.IDToken
		require.NotEmpty(t, idToken)

		claims := src.AuthorizationCodeFlowClaims{}
		token := DecodeJWT(t, idToken, &claims)
		require.True(t, token.Valid)

		require.Equal(t, claims.Iss, h.C.SiteHost.String())
		require.Equal(t, claims.Sub, u.ID)
		require.Equal(t, claims.Nonce, oauthSession.Nonce.String)

		require.Equal(t, claims.StandardClaims.PreferredUsername, "", "profile スコープが無いので取得できない")
		require.Equal(t, claims.StandardClaims.Email, "", "email スコープが無いので取得できない")
	})

	t.Run("失敗: codeが存在しない値", func(t *testing.T) {
		client := RegisterClient(t, ctx, nil)

		redirectUri, err := url.Parse("https://example.test/hogehoge")
		require.NoError(t, err)
		redirect := models.ClientRedirect{
			ClientID: client.ClientID,
			Host:     redirectUri.Host,
			URL:      redirectUri.String(),
		}
		err = redirect.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		query := url.Values{}
		query.Set("code", "hogehoge")
		query.Set("redirect_uri", redirectUri.String())

		m, err := easy.NewURLEncoded("/", http.MethodPost, query)
		require.NoError(t, err)

		c := m.Echo()

		err = h.TokenEndpointAuthorizationCode(ctx, c, client)
		require.EqualError(t, err, "code=400, error=invalid_grant, message=Invalid code")
	})

	t.Run("失敗: codeが空", func(t *testing.T) {
		client := RegisterClient(t, ctx, nil)

		redirectUri, err := url.Parse("https://example.test/hogehoge")
		require.NoError(t, err)
		redirect := models.ClientRedirect{
			ClientID: client.ClientID,
			Host:     redirectUri.Host,
			URL:      redirectUri.String(),
		}
		err = redirect.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		query := url.Values{}
		query.Set("code", "")
		query.Set("redirect_uri", redirectUri.String())

		m, err := easy.NewURLEncoded("/", http.MethodPost, query)
		require.NoError(t, err)

		c := m.Echo()

		err = h.TokenEndpointAuthorizationCode(ctx, c, client)
		require.EqualError(t, err, "code=400, error=invalid_grant, message=Invalid code")
	})

	t.Run("失敗: リダイレクトURIが存在しない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		client := RegisterClient(t, ctx, nil)

		redirectUri, err := url.Parse("https://example.test/hogehoge")
		require.NoError(t, err)

		oauthSession := RegisterOauthSession(t, ctx, client.ClientID, &u)

		query := url.Values{}
		query.Set("code", oauthSession.Code)
		query.Set("redirect_uri", redirectUri.String())

		m, err := easy.NewURLEncoded("/", http.MethodPost, query)
		require.NoError(t, err)

		c := m.Echo()

		err = h.TokenEndpointAuthorizationCode(ctx, c, client)
		require.EqualError(t, err, "code=400, error=invalid_grant, message=Invalid redirect_uri")
	})

	t.Run("失敗: リダイレクトURIの形式が不正", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		client := RegisterClient(t, ctx, nil)

		oauthSession := RegisterOauthSession(t, ctx, client.ClientID, &u)

		query := url.Values{}
		query.Set("code", oauthSession.Code)
		query.Set("redirect_uri", "hogehoge")

		m, err := easy.NewURLEncoded("/", http.MethodPost, query)
		require.NoError(t, err)

		c := m.Echo()

		err = h.TokenEndpointAuthorizationCode(ctx, c, client)
		require.EqualError(t, err, "code=400, error=invalid_grant, message=Invalid redirect_uri")
	})

	t.Run("失敗: リダイレクトURIが空", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		client := RegisterClient(t, ctx, nil)

		oauthSession := RegisterOauthSession(t, ctx, client.ClientID, &u)

		query := url.Values{}
		query.Set("code", oauthSession.Code)

		m, err := easy.NewURLEncoded("/", http.MethodPost, query)
		require.NoError(t, err)

		c := m.Echo()

		err = h.TokenEndpointAuthorizationCode(ctx, c, client)
		require.EqualError(t, err, "code=400, error=invalid_grant, message=Invalid redirect_uri")
	})
}

func TestTokenEndpointRefreshToken(t *testing.T) {
	ctx := context.Background()
	h := NewTestHandler(t)

	t.Run("成功: 更新できる", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		client := RegisterClient(t, ctx, nil)

		refreshToken, err := lib.RandomStr(63)
		require.NoError(t, err)
		sessionToken, err := lib.RandomStr(31)
		require.NoError(t, err)

		clientSession := models.ClientSession{
			ID:       sessionToken,
			UserID:   u.ID,
			ClientID: client.ClientID,
			Period:   time.Now().Add(h.C.OAuthAccessTokenPeriod),
		}
		clientRefresh := models.ClientRefresh{
			ID:        refreshToken,
			UserID:    u.ID,
			ClientID:  client.ClientID,
			SessionID: sessionToken,
			Period:    time.Now().Add(h.C.OAuthRefreshTokenPeriod),
		}
		err = clientSession.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)
		err = clientRefresh.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		query := url.Values{}
		query.Set("refresh_token", refreshToken)

		m, err := easy.NewURLEncoded("/", http.MethodPost, query)
		require.NoError(t, err)

		c := m.Echo()

		err = h.TokenEndpointRefreshToken(ctx, c, client)
		require.NoError(t, err)

		response := src.TokenEndpointResponse{}
		require.NoError(t, m.Json(&response))

		require.Equal(t, response.TokenType, "Bearer")
		require.Equal(t, response.ExpiresIn, int64(h.C.OAuthAccessTokenPeriod)/10000000)

		// トークンが生成されているかチェック
		existNewClientSession, err := models.ClientSessions(
			models.ClientSessionWhere.ID.EQ(response.AccessToken),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.True(t, existNewClientSession)
		existNewClientRefresh, err := models.ClientRefreshes(
			models.ClientRefreshWhere.ID.EQ(response.RefreshToken),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.True(t, existNewClientRefresh)

		// 古いトークンは削除されているかチェック
		existOldClientSession, err := models.ClientSessions(
			models.ClientSessionWhere.ID.EQ(sessionToken),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.False(t, existOldClientSession)
		existOldClientRefresh, err := models.ClientRefreshes(
			models.ClientRefreshWhere.ID.EQ(refreshToken),
		).Exists(ctx, DB)
		require.NoError(t, err)
		require.False(t, existOldClientRefresh)
	})

	t.Run("失敗: リフレッシュトークンのクライアントが一致しない", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		client := RegisterClient(t, ctx, nil)
		client2 := RegisterClient(t, ctx, nil)

		refreshToken, err := lib.RandomStr(63)
		require.NoError(t, err)
		sessionToken, err := lib.RandomStr(31)
		require.NoError(t, err)

		clientSession := models.ClientSession{
			ID:       sessionToken,
			UserID:   u.ID,
			ClientID: client.ClientID,
			Period:   time.Now().Add(h.C.OAuthAccessTokenPeriod),
		}
		clientRefresh := models.ClientRefresh{
			ID:        refreshToken,
			UserID:    u.ID,
			ClientID:  client.ClientID,
			SessionID: sessionToken,
			Period:    time.Now().Add(h.C.OAuthRefreshTokenPeriod),
		}
		err = clientSession.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)
		err = clientRefresh.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		query := url.Values{}
		query.Set("refresh_token", refreshToken)

		m, err := easy.NewURLEncoded("/", http.MethodPost, query)
		require.NoError(t, err)

		c := m.Echo()

		err = h.TokenEndpointRefreshToken(ctx, c, client2)
		require.EqualError(t, err, "code=400, error=invalid_grant, message=Invalid refresh_token")
	})

	t.Run("失敗: リフレッシュトークンが存在しない値", func(t *testing.T) {
		client := RegisterClient(t, ctx, nil)

		query := url.Values{}
		query.Set("refresh_token", "invalid")

		m, err := easy.NewURLEncoded("/", http.MethodPost, query)
		require.NoError(t, err)

		c := m.Echo()

		err = h.TokenEndpointRefreshToken(ctx, c, client)
		require.EqualError(t, err, "code=400, error=invalid_grant, message=Invalid refresh_token")
	})

	t.Run("失敗: リフレッシュトークンが空", func(t *testing.T) {
		client := RegisterClient(t, ctx, nil)

		query := url.Values{}
		query.Set("refresh_token", "")

		m, err := easy.NewURLEncoded("/", http.MethodPost, query)
		require.NoError(t, err)

		c := m.Echo()

		err = h.TokenEndpointRefreshToken(ctx, c, client)
		require.EqualError(t, err, "code=400, error=invalid_grant, message=Invalid refresh_token")
	})

	t.Run("失敗: リフレッシュトークンの有効期限が切れている", func(t *testing.T) {
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)
		client := RegisterClient(t, ctx, nil)

		refreshToken, err := lib.RandomStr(63)
		require.NoError(t, err)
		sessionToken, err := lib.RandomStr(31)
		require.NoError(t, err)

		clientSession := models.ClientSession{
			ID:       sessionToken,
			UserID:   u.ID,
			ClientID: client.ClientID,
			Period:   time.Now().Add(h.C.OAuthAccessTokenPeriod),
		}
		clientRefresh := models.ClientRefresh{
			ID:        refreshToken,
			UserID:    u.ID,
			ClientID:  client.ClientID,
			SessionID: sessionToken,
			Period:    time.Now().Add(-1 * time.Hour), // 有効期限切れ
		}
		err = clientSession.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)
		err = clientRefresh.Insert(ctx, DB, boil.Infer())
		require.NoError(t, err)

		query := url.Values{}
		query.Set("refresh_token", refreshToken)

		m, err := easy.NewURLEncoded("/", http.MethodPost, query)
		require.NoError(t, err)

		c := m.Echo()

		err = h.TokenEndpointRefreshToken(ctx, c, client)
		require.EqualError(t, err, "code=400, error=invalid_grant, message=Invalid refresh_token")
	})
}

func TestUserToStandardClaims(t *testing.T) {
	t.Run("StandardClaimsに変換できる", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		claims, err := src.UserToStandardClaims(&u, []string{"openid", "profile", "email"})
		require.NoError(t, err)

		require.NotNil(t, claims)

		require.Equal(t, claims.Name, u.UserName)
		require.Equal(t, claims.GivenName, u.GivenName.String)
		require.Equal(t, claims.FamilyName, u.FamilyName.String)
		require.Equal(t, claims.MiddleName, u.MiddleName.String)
		require.Equal(t, claims.Nickname, u.UserName)
		require.Equal(t, claims.PreferredUsername, u.UserName)
		require.Equal(t, claims.Picture, u.Avatar.String)
		require.Equal(t, claims.Email, u.Email)
		require.True(t, claims.EmailVerified)
		require.Equal(t, claims.Gender, u.Gender)
		require.Equal(t, claims.ZoneInfo, "Asia/Tokyo")
		require.Equal(t, claims.Locale, "ja-JP")
		require.Equal(t, claims.UpdatedAt, u.UpdatedAt.Unix())

		require.Equal(t, claims.BirthDate, "")
	})

	t.Run("BirthDateが設定されている", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		birthDate := time.Now()

		u.Birthdate = null.TimeFrom(birthDate)
		_, err := u.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		claims, err := src.UserToStandardClaims(&u, []string{"openid", "profile", "email"})
		require.NoError(t, err)

		require.Equal(t, claims.BirthDate, birthDate.Format(time.DateOnly))
	})

	t.Run("scopesにemailが含まれていない場合はemailは空", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		birthDate := time.Now()

		u.Birthdate = null.TimeFrom(birthDate)
		_, err := u.Update(ctx, DB, boil.Infer())
		require.NoError(t, err)

		claims, err := src.UserToStandardClaims(&u, []string{"openid", "profile"})
		require.NoError(t, err)

		require.Equal(t, claims.Email, "")
		require.False(t, claims.EmailVerified)
	})

	t.Run("scopesにprofileが含まれていない場合はプロフィールは空", func(t *testing.T) {
		ctx := context.Background()
		email := RandomEmail(t)
		u := RegisterUser(t, ctx, email)

		claims, err := src.UserToStandardClaims(&u, []string{"openid", "email"})
		require.NoError(t, err)

		require.NotNil(t, claims)

		require.Equal(t, claims.Name, "")
		require.Equal(t, claims.GivenName, "")
		require.Equal(t, claims.FamilyName, "")
		require.Equal(t, claims.MiddleName, "")
		require.Equal(t, claims.Nickname, "")
		require.Equal(t, claims.PreferredUsername, "")
		require.Equal(t, claims.Picture, "")
		require.Equal(t, claims.Email, u.Email)
		require.True(t, claims.EmailVerified)
		require.Equal(t, claims.Gender, "")
		require.Equal(t, claims.ZoneInfo, "Asia/Tokyo")
		require.Equal(t, claims.Locale, "ja-JP")
		require.Equal(t, claims.UpdatedAt, int64(0))
		require.Equal(t, claims.BirthDate, "")
	})
}
