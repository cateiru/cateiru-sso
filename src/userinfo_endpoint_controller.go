package src

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/labstack/echo/v4"
)

type UserinfoResponse struct {
	Sub string `json:"sub"`

	StandardClaims
}

// userinfo endpoint の認証
func (h *Handler) UserinfoAuthentication(ctx context.Context, c echo.Context) (*models.ClientSession, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		c.Response().Header().Set("WWW-Authenticate", "Bearer")
		return nil, NewOIDCError(http.StatusBadRequest, ErrUserinfoInvalidRequest, "Invalid code", "", "")
	}

	token := authHeader[len("Bearer "):]
	if token == "" {
		c.Response().Header().Set("WWW-Authenticate", "Bearer")
		return nil, NewOIDCError(http.StatusBadRequest, ErrUserinfoInvalidRequest, "Invalid code", "", "")
	}

	clientSession, err := models.ClientSessions(
		models.ClientSessionWhere.ID.EQ(token),
		models.ClientSessionWhere.Period.GT(time.Now()),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		c.Response().Header().Set("WWW-Authenticate", "Bearer")
		return nil, NewOIDCError(http.StatusBadRequest, ErrUserinfoInvalidRequest, "Invalid code", "", "")
	}
	if err != nil {
		return nil, err
	}

	return clientSession, nil
}

// userinfo のレスポンスを返す。JWTによる署名はしない
// ref. https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#UserInfoResponse
func (h *Handler) ResponseStandardClaims(ctx context.Context, clientId string, userId string) (*UserinfoResponse, error) {
	client, err := models.Clients(
		models.ClientWhere.ClientID.EQ(clientId),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, NewOIDCError(http.StatusBadRequest, ErrUserinfoInvalidRequest, "Client not found", "", "")
	}
	if err != nil {
		return nil, err
	}

	user, err := models.Users(
		models.UserWhere.ID.EQ(userId),
	).One(ctx, h.DB)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, NewOIDCError(http.StatusBadRequest, ErrUserinfoInvalidRequest, "User not found", "", "")
	}
	if err != nil {
		return nil, err
	}

	clientScopes, err := models.ClientScopes(
		models.ClientScopeWhere.ClientID.EQ(client.ClientID),
	).All(ctx, h.DB)
	if err != nil {
		return nil, err
	}
	scopes := make([]string, len(clientScopes))
	for i, clientScope := range clientScopes {
		scopes[i] = clientScope.Scope
	}

	standardClaims, err := UserToStandardClaims(user, scopes)
	if err != nil {
		return nil, err
	}

	return &UserinfoResponse{
		Sub:            user.ID,
		StandardClaims: *standardClaims,
	}, nil
}
