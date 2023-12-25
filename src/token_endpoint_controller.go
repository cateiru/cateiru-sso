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
	ExpiresIn int `json:"expires_in"`

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

	// client_secret_post
	clientId := c.QueryParam("client_id")
	clientSecret := c.QueryParam("client_secret")
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

// TODO
// ref. https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#TokenRequest
// validation: https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#TokenRequestValidation
func (h *Handler) TokenEndpointAuthorizationCode(ctx context.Context, c echo.Context, client *models.Client) error {
	code := c.QueryParam("code")
	redirectUri := c.QueryParam("redirect_uri")
	clientId := c.QueryParam("client_id")

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

	// `client_id` が指定されている場合は、検証する
	// https://openid-foundation-japan.github.io/rfc6749.ja.html#token-req では、「認可サーバーよってクライアントが認証されていなければ必須 」
	// となっているが、本プロジェクトではクライアントは認証されているはずなのでオプショナルとしている
	// 認証: https://openid-foundation-japan.github.io/rfc6749.ja.html#token-endpoint-auth
	// XXX: 本当に？
	if clientId != "" {
		if clientId != client.ClientID {
			return NewOIDCError(http.StatusBadRequest, ErrTokenInvalidGrant, "Invalid client_id", "", "")
		}
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

	// XXX; authorization code flow だけで良い？
	// TODO: claims埋める
	claims := AuthorizationCodeFlowClaims{
		IDTokenClaimsBase: IDTokenClaimsBase{
			Iss:      h.C.OIDCIssuer,
			Sub:      oauthSession.UserID,
			Aud:      client.ClientID,
			Exp:      0, // TODO
			Iat:      0, // TODO
			AuthTime: time.Now().Unix(),
			Nonce:    oauthSession.Nonce.String,
		},
	}

	idToken, err := lib.SignJwt(claims, h.C.JWTPrivateKeyFilePath)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &TokenEndpointResponse{
		AccessToken:  "TODO",
		TokenType:    "Bearer",
		RefreshToken: "TODO",
		ExpiresIn:    00, // TODO
		IDToken:      idToken,
	})
}

// TODO
// ref. https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#RefreshingAccessToken
func (h *Handler) TokenEndpointRefreshToken(ctx context.Context, c echo.Context, client *models.Client) error {
	return nil
}
