package src

import (
	"net/http"

	"github.com/cateiru/cateiru-sso/src/lib"
	"github.com/labstack/echo/v4"
)

// TODO: テスト
// OIDC Token Endpoint
// ref. https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#TokenEndpoint
func (h *Handler) TokenEndpointHandler(c echo.Context) error {
	ctx := c.Request().Context()

	// レスポンスのキャッシュを無効化
	// ref. https://openid-foundation-japan.github.io/rfc6749.ja.html#token-response
	c.Response().Header().Set("Cache-Control", "no-store")
	c.Response().Header().Set("Pragma", "no-cache")

	// 認証
	client, err := h.ClientAuthentication(ctx, c)
	if err != nil {
		return err
	}

	param, err := h.QueryBodyParam(c)
	if err != nil {
		return err
	}

	grantType := param.Get("grant_type")
	formattedGrantType := lib.ValidateTokenEndpointGrantType(grantType)

	switch formattedGrantType {
	case lib.TokenEndpointGrantTypeAuthorizationCode:
		return h.TokenEndpointAuthorizationCode(ctx, c, client)

	case lib.TokenEndpointGrantTypeRefreshToken:
		return h.TokenEndpointRefreshToken(ctx, c, client)

	default:
		return NewOIDCError(http.StatusBadRequest, ErrTokenUnsupportedGrantType, "unsupported grant type", "", "")

	}
}

// TODO: テスト
// OIDC Userinfo Endpoint
// ref. https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#UserInfoRequest
func (h *Handler) UserinfoEndpointHandler(c echo.Context) error {
	ctx := c.Request().Context()

	clientSession, err := h.UserinfoAuthentication(ctx, c)
	if err != nil {
		return err
	}

	response, err := h.ResponseStandardClaims(ctx, clientSession.ClientID, clientSession.UserID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}
