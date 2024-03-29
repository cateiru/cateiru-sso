package src

import (
	"database/sql"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ClientLoginResponse struct {
	CreatedAt time.Time `json:"created_at"`

	Client *ClientHistoryResponse `json:"client"`
}

type ClientLoginHistoryResponse struct {
	ID uint `json:"id"`

	Client *ClientHistoryResponse `json:"client"`

	Device   null.String `json:"device"`
	OS       null.String `json:"os"`
	Browser  null.String `json:"browser"`
	IsMobile null.Bool   `json:"is_mobile"`
	Ip       string      `json:"ip"`

	CreatedAt time.Time `json:"created_at"`
}

type ClientHistoryResponse struct {
	ClientID string `json:"client_id"`

	Name        string      `json:"name"`
	Description null.String `json:"description,omitempty"`
	Image       null.String `json:"image,omitempty"`
}

type LoginDeviceResponse struct {
	ID uint `json:"id"`

	Device   null.String `json:"device"`
	OS       null.String `json:"os"`
	Browser  null.String `json:"browser"`
	IsMobile null.Bool   `json:"is_mobile"`
	Ip       string      `json:"ip"`

	IsCurrent bool `json:"is_current"`

	CreatedAt time.Time `json:"created_at"`
}

type LoginTryHistoryResponse struct {
	ID uint `json:"id"`

	Device   null.String `json:"device"`
	OS       null.String `json:"os"`
	Browser  null.String `json:"browser"`
	IsMobile null.Bool   `json:"is_mobile"`
	Ip       string      `json:"ip"`

	Identifier int8 `json:"identifier"`

	CreatedAt time.Time `json:"created_at"`
}

type OperationHistoryResponse struct {
	ID uint `json:"id"`

	Device   null.String `json:"device"`
	OS       null.String `json:"os"`
	Browser  null.String `json:"browser"`
	IsMobile null.Bool   `json:"is_mobile"`
	Ip       string      `json:"ip"`

	Identifier int8 `json:"identifier"`

	CreatedAt time.Time `json:"created_at"`
}

type LoginHistoriesSlice struct {
	LoginHistory models.LoginHistory `boil:",bind"`
	Refresh      models.Refresh      `boil:",bind"`
}

// ログインしているSSOクライアント
func (h *Handler) HistoryClientLoginHandler(c echo.Context) error {
	ctx := c.Request().Context()

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	clientRefresh, err := models.ClientRefreshes(
		models.ClientRefreshWhere.UserID.EQ(u.ID),
		qm.And("period > NOW()"),
		qm.OrderBy("created_at DESC"),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}

	cacheClient := []*ClientHistoryResponse{}
	cacheEmptyClientId := []string{}
	getClient := func(clientID string) (*ClientHistoryResponse, error) {
		for _, cache := range cacheClient {
			if cache.ClientID == clientID {
				return cache, nil
			}
		}
		for _, id := range cacheEmptyClientId {
			if id == clientID {
				return nil, nil
			}
		}

		clientFromDB, err := models.Clients(
			models.ClientWhere.ClientID.EQ(clientID),
		).One(ctx, h.DB)
		if errors.Is(err, sql.ErrNoRows) {
			cacheEmptyClientId = append(cacheEmptyClientId, clientID)
			return nil, nil
		}
		if err != nil {
			return nil, err
		}

		client := &ClientHistoryResponse{
			ClientID:    clientFromDB.ClientID,
			Name:        clientFromDB.Name,
			Description: clientFromDB.Description,
			Image:       clientFromDB.Image,
		}
		cacheClient = append(cacheClient, client)
		return client, nil
	}

	logins := []ClientLoginResponse{}
	for _, r := range clientRefresh {
		client, err := getClient(r.ClientID)
		if err != nil {
			return err
		}

		logins = append(logins, ClientLoginResponse{
			CreatedAt: r.CreatedAt,

			Client: client,
		})
	}

	return c.JSON(http.StatusOK, logins)
}

// クライアントのログイン履歴
func (h *Handler) HistoryClientHandler(c echo.Context) error {
	ctx := c.Request().Context()

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	histories, err := models.LoginClientHistories(
		models.LoginClientHistoryWhere.UserID.EQ(u.ID),
		qm.Limit(50),
		qm.OrderBy("created_at DESC"),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}

	cacheClient := []*ClientHistoryResponse{}
	cacheEmptyClientId := []string{}
	getClient := func(clientID string) (*ClientHistoryResponse, error) {
		for _, cache := range cacheClient {
			if cache.ClientID == clientID {
				return cache, nil
			}
		}
		for _, id := range cacheEmptyClientId {
			if id == clientID {
				return nil, nil
			}
		}

		clientFromDB, err := models.Clients(
			models.ClientWhere.ClientID.EQ(clientID),
		).One(ctx, h.DB)
		if errors.Is(err, sql.ErrNoRows) {
			cacheEmptyClientId = append(cacheEmptyClientId, clientID)
			return nil, nil
		}
		if err != nil {
			return nil, err
		}

		client := &ClientHistoryResponse{
			ClientID:    clientFromDB.ClientID,
			Name:        clientFromDB.Name,
			Description: clientFromDB.Description,
			Image:       clientFromDB.Image,
		}
		cacheClient = append(cacheClient, client)
		return client, nil
	}

	clientHistories := []ClientLoginHistoryResponse{}
	for _, history := range histories {
		client, err := getClient(history.ClientID)
		if err != nil {
			return err
		}

		clientHistories = append(clientHistories, ClientLoginHistoryResponse{
			ID: history.ID,

			Device:   history.Device,
			OS:       history.Os,
			Browser:  history.Browser,
			IsMobile: history.IsMobile,
			Ip:       net.IP.To16(history.IP).String(),

			CreatedAt: history.CreatedAt,

			Client: client,
		})
	}

	return c.JSON(http.StatusOK, clientHistories)
}

// 現在ログインしているデバイス
func (h *Handler) HistoryLoginDeviceHandler(c echo.Context) error {
	ctx := c.Request().Context()

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	var refreshToken *http.Cookie
	refreshTokenName := fmt.Sprintf("%s-%s", h.C.RefreshCookie.Name, u.ID)
	for _, cookie := range c.Cookies() {
		if cookie.Name == refreshTokenName {
			refreshToken = cookie
		}
	}

	var loginDevices []LoginHistoriesSlice
	// SELECT login_history.*, refresh.* FROM login_history
	// INNER JOIN refresh
	//     on refresh.id = login_history.refresh_id
	// WHERE login_history.user_id = ?
	// AND refresh.period > NOW()
	// ORDER BY login_history.created DESC
	// LIMIT 50;
	err = models.NewQuery(
		qm.Select(
			"login_history.id",
			"login_history.user_id",
			"login_history.refresh_id",
			"login_history.device",
			"login_history.os",
			"login_history.browser",
			"login_history.is_mobile",
			"login_history.ip",
			"login_history.created_at",

			"refresh.id",
			"refresh.user_id",
			"refresh.history_id",
			"refresh.session_id",
			"refresh.period",
			"refresh.created_at",
			"refresh.updated_at",
		),
		qm.From("login_history"),
		qm.InnerJoin("refresh ON refresh.history_id = login_history.refresh_id"),
		qm.Where("login_history.user_id = ?", u.ID),
		qm.And("refresh.period > NOW()"),
		qm.OrderBy("login_history.created_at DESC"),
		qm.Limit(50),
	).Bind(ctx, h.DB, &loginDevices)
	if err != nil {
		return err
	}

	formattedLoginDevices := []LoginDeviceResponse{}
	for _, l := range loginDevices {
		formattedLoginDevices = append(formattedLoginDevices, LoginDeviceResponse{
			ID: l.LoginHistory.ID,

			Device:   l.LoginHistory.Device,
			OS:       l.LoginHistory.Os,
			Browser:  l.LoginHistory.Browser,
			IsMobile: l.LoginHistory.IsMobile,
			Ip:       net.IP.To16(l.LoginHistory.IP).String(),

			IsCurrent: l.Refresh.ID == refreshToken.Value,

			CreatedAt: l.LoginHistory.CreatedAt,
		})
	}

	return c.JSON(http.StatusOK, formattedLoginDevices)
}

// ログイン履歴
func (h *Handler) HistoryLoginHistoryHandler(c echo.Context) error {
	ctx := c.Request().Context()

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	// SELECT * FROM login_history
	// WHERE user_id = ?
	// ORDER BY created DESC
	// LIMIT 50;
	loginHistories, err := models.LoginHistories(
		models.LoginHistoryWhere.UserID.EQ(u.ID),
		qm.OrderBy("created_at DESC"),
		qm.Limit(50),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}

	formattedLoginHistories := []LoginDeviceResponse{}
	for _, l := range loginHistories {
		formattedLoginHistories = append(formattedLoginHistories, LoginDeviceResponse{
			ID: l.ID,

			Device:   l.Device,
			OS:       l.Os,
			Browser:  l.Browser,
			IsMobile: l.IsMobile,
			Ip:       net.IP.To16(l.IP).String(),

			IsCurrent: false,

			CreatedAt: l.CreatedAt,
		})
	}

	return c.JSON(http.StatusOK, formattedLoginHistories)
}

// ログイントライ履歴
func (h *Handler) HistoryLoginTryHistoryHandler(c echo.Context) error {
	ctx := c.Request().Context()

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	// SELECT * FROM login_try_history
	// WHERE user_id = ?
	// ORDER BY created DESC
	// LIMIT 50;
	loginTryHistries, err := models.LoginTryHistories(
		models.LoginTryHistoryWhere.UserID.EQ(u.ID),
		qm.OrderBy("created_at DESC"),
		qm.Limit(50),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}

	formattedLoginTryHistories := []LoginTryHistoryResponse{}
	for _, l := range loginTryHistries {
		formattedLoginTryHistories = append(formattedLoginTryHistories, LoginTryHistoryResponse{
			ID: l.ID,

			Device:   l.Device,
			OS:       l.Os,
			Browser:  l.Browser,
			IsMobile: l.IsMobile,
			Ip:       net.IP.To16(l.IP).String(),

			Identifier: l.Identifier,

			CreatedAt: l.CreatedAt,
		})
	}

	return c.JSON(http.StatusOK, formattedLoginTryHistories)
}

func (h *Handler) HistoryOperationHistoryHandler(c echo.Context) error {
	ctx := c.Request().Context()

	u, err := h.Session.SimpleLogin(ctx, c)
	if err != nil {
		return err
	}

	// SELECT * FROM login_try_history
	// WHERE user_id = ?
	// ORDER BY created DESC
	// LIMIT 50;
	operationHistory, err := models.OperationHistories(
		models.OperationHistoryWhere.UserID.EQ(u.ID),
		qm.OrderBy("created_at DESC"),
		qm.Limit(50),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}

	formattedOperationHistories := []OperationHistoryResponse{}
	for _, l := range operationHistory {
		formattedOperationHistories = append(formattedOperationHistories, OperationHistoryResponse{
			ID: l.ID,

			Device:   l.Device,
			OS:       l.Os,
			Browser:  l.Browser,
			IsMobile: l.IsMobile,
			Ip:       net.IP.To16(l.IP).String(),

			Identifier: l.Identifier,

			CreatedAt: l.CreatedAt,
		})
	}

	return c.JSON(http.StatusOK, formattedOperationHistories)
}
