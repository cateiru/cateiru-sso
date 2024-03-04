package src

import (
	"context"
	"database/sql"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"golang.org/x/exp/slices"
)

// トークンエンドポイントのレスポンス
// https://openid-foundation-japan.github.io/rfc6749.ja.html#token-response
type TokenEndpointResponse struct {
	// 必須 (REQUIRED)。認可サーバーが発行するアクセストークン。
	AccessToken string `json:"access_token"`

	// 必須 (REQUIRED)。トークンのタイプ。値は大文字・小文字を区別しない。詳細は Section 7.1 を参照のこと。
	TokenType string `json:"token_type"`

	// 推奨 (RECOMMENDED)。アクセストークンの有効期間を表す秒数。例えばこの値が 3600 であれば、そのアクセストークンは発行から1時間後に期限切れとなる。
	// 省略された場合、認可サーバはドキュメントまたは他の手段によってデフォルトの有効期間を提示すべきである (SHOULD)。
	ExpiresIn int64 `json:"expires_in"`

	// 任意 (OPTIONAL)。リフレッシュトークン。
	// 同じ認可グラントを用いて新しいアクセストークンを取得するのに利用される。詳細は Section 6 を参照のこと。
	RefreshToken string `json:"refresh_token,omitempty"`

	// クライアントから全く同一のスコープが要求された場合は任意 (OPTIONAL)。その他は必須 (REQUIRED)。アクセストークンのスコープ。詳細は Section 3.3 を参照のこと。
	Scope string `json:"scope,omitempty"`

	// 認証セッションに紐づいた ID Token 値。 OIDC 独自のパラメータ。
	// ref. https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#TokenResponse
	IDToken string `json:"id_token,omitempty"`
}

// Token Endpoint の認証
// `client_secret_basic` と `client_secret_post` に対応している
// ref. https://openid-foundation-japan.github.io/rfc6749.ja.html#client-password
// ref. https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#ClientAuthentication
func (h *Handler) ClientAuthentication(ctx context.Context, c echo.Context) (*models.Client, error) {
	// client_secret_basic
	basic := c.Request().Header.Get("Authorization")
	if basic != "" {
		splitBasic := strings.Split(basic, " ")
		if len(splitBasic) != 2 || splitBasic[0] != "Basic" {
			return nil, NewOIDCError(http.StatusBadRequest, ErrTokenInvalidRequest, "Invalid Authorization Header", "", "")
		}

		// Basic認証のデコード
		decoded, err := base64.StdEncoding.DecodeString(splitBasic[1])
		if err != nil {
			return nil, NewOIDCError(http.StatusBadRequest, ErrTokenInvalidRequest, "Invalid Authorization Header", "", "")
		}

		// client_id:client_secret の形式になっているか確認
		splitDecoded := strings.Split(string(decoded), ":")
		if len(splitDecoded) != 2 {
			return nil, NewOIDCError(http.StatusBadRequest, ErrTokenInvalidRequest, "Invalid Authorization Header", "", "")
		}

		client, err := models.Clients(
			models.ClientWhere.ClientID.EQ(splitDecoded[0]),
		).One(ctx, h.DB)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NewOIDCError(http.StatusBadRequest, ErrTokenInvalidRequest, "Invalid client_id", "", "")
		}
		if err != nil {
			return nil, err
		}

		// シークレットを検証
		if client.ClientSecret != splitDecoded[1] {
			return nil, NewOIDCError(http.StatusBadRequest, ErrTokenInvalidRequest, "Invalid client_secret", "", "")
		}

		return client, nil
	}

	param, err := h.QueryBodyParam(c)
	if err != nil {
		return nil, err
	}

	// client_secret_post
	clientId := param.Get("client_id")
	clientSecret := param.Get("client_secret")
	if clientId != "" || clientSecret != "" {
		client, err := models.Clients(
			models.ClientWhere.ClientID.EQ(clientId),
		).One(ctx, h.DB)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NewOIDCError(http.StatusBadRequest, ErrTokenInvalidRequest, "Invalid client_id", "", "")
		}
		if err != nil {
			return nil, err
		}

		// シークレットを検証
		if client.ClientSecret != clientSecret {
			return nil, NewOIDCError(http.StatusBadRequest, ErrTokenInvalidRequest, "Invalid client_secret", "", "")
		}

		return client, nil
	}

	// どの認証方式でも無い場合は、WWW-Authenticate を付与して 401 を返す
	c.Response().Header().Set("WWW-Authenticate", "Basic")

	return nil, NewOIDCError(http.StatusUnauthorized, ErrTokenInvalidClient, "Invalid client authentication", "", "")
}

// ref. https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#TokenRequest
// validation: https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#TokenRequestValidation
func (h *Handler) TokenEndpointAuthorizationCode(ctx context.Context, c echo.Context, client *models.Client) error {
	param, err := h.QueryBodyParam(c)
	if err != nil {
		return err
	}

	code := param.Get("code")
	redirectUri := param.Get("redirect_uri")

	parsedRedirectUri, redirectUriOk := lib.ValidateURL(redirectUri)
	if !redirectUriOk {
		return NewOIDCError(http.StatusBadRequest, ErrTokenInvalidGrant, "Invalid redirect_uri", "", "")
	}

	existRedirect, err := models.ClientRedirects(
		models.ClientRedirectWhere.ClientID.EQ(client.ClientID),
		models.ClientRedirectWhere.URL.EQ(parsedRedirectUri.String()),
	).Exists(ctx, h.DB)
	if err != nil {
		return err
	}
	if !existRedirect {
		return NewOIDCError(http.StatusBadRequest, ErrTokenInvalidGrant, "Invalid redirect_uri", "", "")
	}

	// code の検証
	oauthSession, err := models.OauthSessions(
		models.OauthSessionWhere.Code.EQ(code),
		models.OauthSessionWhere.Period.GT(time.Now()),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewOIDCError(http.StatusBadRequest, ErrTokenInvalidGrant, "Invalid code", "", "")
	}
	if err != nil {
		return err
	}

	// OauthSession は一度しか使わないものなので削除してしまう
	if _, err := oauthSession.Delete(ctx, h.DB); err != nil {
		return err
	}

	authUser, err := models.Users(
		models.UserWhere.ID.EQ(oauthSession.UserID),
	).One(ctx, h.DB)
	if err != nil {
		return err
	}

	// トークン作成
	refreshToken, err := lib.RandomStr(63)
	if err != nil {
		return err
	}
	sessionToken, err := lib.RandomStr(31)
	if err != nil {
		return err
	}
	clientSession := models.ClientSession{
		ID:       sessionToken,
		UserID:   authUser.ID,
		ClientID: client.ClientID,
		Period:   time.Now().Add(h.C.OAuthAccessTokenPeriod),
	}
	clientRefresh := models.ClientRefresh{
		ID:        refreshToken,
		UserID:    authUser.ID,
		ClientID:  client.ClientID,
		SessionID: sessionToken,
		Period:    time.Now().Add(h.C.OAuthRefreshTokenPeriod),
	}

	err = TxDB(ctx, h.DB, func(tx *sql.Tx) error {
		if err := clientSession.Insert(ctx, h.DB, boil.Infer()); err != nil {
			return err
		}
		if err := clientRefresh.Insert(ctx, h.DB, boil.Infer()); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	clientScopes, err := models.ClientScopes(
		models.ClientScopeWhere.ClientID.EQ(client.ClientID),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}
	scopes := make([]string, len(clientScopes))
	for i, clientScope := range clientScopes {
		scopes[i] = clientScope.Scope
	}

	standardClaims, err := UserToStandardClaims(authUser, scopes)
	if err != nil {
		return err
	}

	// XXX: authorization code flow だけで良い？
	claims := AuthorizationCodeFlowClaims{
		IDTokenClaimsBase: IDTokenClaimsBase{
			Iss:      h.C.SiteHost.String(),
			Sub:      oauthSession.UserID,
			Aud:      client.ClientID,
			Exp:      time.Now().Add(h.C.IDTokenExpire).Unix(),
			Iat:      time.Now().Unix(),
			AuthTime: oauthSession.AuthTime.Unix(),
			Nonce:    oauthSession.Nonce.String,
		},

		StandardClaims: *standardClaims,

		// AtHash: "", TODO: 埋める
	}
	idToken, err := lib.SignJwt(claims, h.C.JWTPrivateKeyFilePath)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &TokenEndpointResponse{
		AccessToken:  sessionToken,
		TokenType:    "Bearer",
		RefreshToken: refreshToken,
		ExpiresIn:    int64(h.C.OAuthAccessTokenPeriod) / 10000000, // time.Duration はマイクロ秒なので秒に変換
		IDToken:      idToken,
	})
}

// リフレッシュトークンの更新
// ref. https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#RefreshingAccessToken
func (h *Handler) TokenEndpointRefreshToken(ctx context.Context, c echo.Context, client *models.Client) error {
	param, err := h.QueryBodyParam(c)
	if err != nil {
		return err
	}

	refreshToken := param.Get("refresh_token")
	if refreshToken == "" {
		return NewOIDCError(http.StatusBadRequest, ErrTokenInvalidGrant, "Invalid refresh_token", "", "")
	}

	clientRefresh, err := models.ClientRefreshes(
		models.ClientRefreshWhere.ID.EQ(refreshToken),
		models.ClientRefreshWhere.Period.GT(time.Now()),
		models.ClientRefreshWhere.ClientID.EQ(client.ClientID),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return NewOIDCError(http.StatusBadRequest, ErrTokenInvalidGrant, "Invalid refresh_token", "", "")
	}
	if err != nil {
		return err
	}

	// 新しいセッショントークンとリフレッシュトークンを用意
	newRefreshToken, err := lib.RandomStr(63)
	if err != nil {
		return err
	}
	newSessionToken, err := lib.RandomStr(31)
	if err != nil {
		return err
	}

	newClientSession := models.ClientSession{
		ID:       newSessionToken,
		UserID:   clientRefresh.UserID,
		ClientID: client.ClientID,
		Period:   time.Now().Add(h.C.OAuthAccessTokenPeriod),
	}
	newClientRefresh := models.ClientRefresh{
		ID:        newRefreshToken,
		UserID:    clientRefresh.UserID,
		ClientID:  clientRefresh.ClientID,
		SessionID: newSessionToken,
		Period:    time.Now().Add(h.C.OAuthRefreshTokenPeriod),
	}
	err = TxDB(ctx, h.DB, func(tx *sql.Tx) error {
		// セッション・リフレッシュは削除しておく
		if _, err := models.ClientSessions(
			models.ClientSessionWhere.ID.EQ(clientRefresh.SessionID),
		).DeleteAll(ctx, h.DB); err != nil {
			return err
		}
		if _, err := clientRefresh.Delete(ctx, h.DB); err != nil {
			return err
		}

		// 新たに作成
		if err := newClientSession.Insert(ctx, h.DB, boil.Infer()); err != nil {
			return err
		}
		if err := newClientRefresh.Insert(ctx, h.DB, boil.Infer()); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &TokenEndpointResponse{
		AccessToken:  newSessionToken,
		TokenType:    "Bearer",
		RefreshToken: newRefreshToken,
		ExpiresIn:    int64(h.C.OAuthAccessTokenPeriod) / 10000000, // time.Duration はマイクロ秒なので秒に変換
	})
}

func UserToStandardClaims(user *models.User, scopes []string) (*StandardClaims, error) {
	standardClaims := &StandardClaims{
		ZoneInfo: "Asia/Tokyo", // 決め打ち
		Locale:   "ja-JP",      // 決め打ち
	}

	// scopes に profile が含まれていたらユーザープロフィールを含める
	if slices.Contains(scopes, "profile") {
		standardClaims.Name = user.UserName
		standardClaims.GivenName = user.GivenName.String
		standardClaims.FamilyName = user.FamilyName.String
		standardClaims.MiddleName = user.MiddleName.String
		standardClaims.Nickname = user.UserName
		standardClaims.PreferredUsername = user.UserName
		standardClaims.Picture = user.Avatar.String
		standardClaims.Gender = user.Gender
		standardClaims.UpdatedAt = user.UpdatedAt.Unix()

		if user.Birthdate.Valid {
			standardClaims.BirthDate = user.Birthdate.Time.Format(time.DateOnly)
		}
	}

	// scopes に email が含まれていたら email を含める
	if slices.Contains(scopes, "email") {
		standardClaims.Email = user.Email
		standardClaims.EmailVerified = true // 必ず確認しているのでtrue
	}

	return standardClaims, nil
}
