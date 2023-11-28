package src

import (
	"context"
	"database/sql"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"

	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/labstack/echo/v4"
)

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
func (h *Handler) TokenEndpointAuthorizationCode(ctx context.Context, c echo.Context, client *models.Client) error {
	_ = c.QueryParam("code")

	return nil
}

// TODO
// ref. https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#RefreshingAccessToken
func (h *Handler) TokenEndpointRefreshToken(ctx context.Context, c echo.Context, client *models.Client) error {
	return nil
}
