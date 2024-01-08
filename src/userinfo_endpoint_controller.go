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
