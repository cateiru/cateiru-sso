package src

import (
	"net"
	"net/http"
	"time"

	"github.com/cateiru/cateiru-sso/src/models"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ClientLoginResponse struct {
	ClientID string `json:"client_id"`

	Scope   []string
	Created time.Time `json:"created"`

	// client
	Name        string      `json:"name"`
	Description null.String `json:"description,omitempty"`
	Image       null.String `json:"image,omitempty"`
}

type ClientLoginHistoryResponse struct {
	ID       uint   `json:"id"`
	ClientID string `json:"client_id"`

	// client
	Name        string      `json:"name"`
	Description null.String `json:"description,omitempty"`
	Image       null.String `json:"image,omitempty"`

	Device   null.String `json:"device"`
	OS       null.String `json:"os"`
	Browser  null.String `json:"browser"`
	IsMobile null.Bool   `json:"is_mobile"`
	Ip       string      `json:"ip"`

	Created time.Time `json:"created"`
}

type LoginDeviceResponse struct {
	ID uint `json:"id"`

	Device   null.String `json:"device"`
	OS       null.String `json:"os"`
	Browser  null.String `json:"browser"`
	IsMobile null.Bool   `json:"is_mobile"`
	Ip       string      `json:"ip"`

	Created time.Time `json:"time"`
}

type LoginTryHistoryResponse struct {
	ID uint `json:"id"`

	Device   null.String `json:"device"`
	OS       null.String `json:"os"`
	Browser  null.String `json:"browser"`
	IsMobile null.Bool   `json:"is_mobile"`
	Ip       string      `json:"ip"`

	Identifier int8 `json:"identifier"`

	Created time.Time `json:"time"`
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
		qm.OrderBy("created DESC"),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}

	clients := []*models.Client{}
	logins := []ClientLoginResponse{}
	for _, r := range clientRefresh {
		// キャッシュしてDBから取得する
		var client *models.Client = nil
		for _, mem := range clients {
			if mem.ClientID == r.ClientID {
				client = mem
				break
			}
		}
		if client == nil {
			clientFromDB, err := models.Clients(
				models.ClientWhere.ClientID.EQ(r.ClientID),
			).One(ctx, h.DB)
			if err != nil {
				return err
			}
			clients = append(clients, clientFromDB)
			client = clientFromDB
		}

		scope := []string{}
		if err := r.Scopes.Unmarshal(&scope); err != nil {
			return err
		}

		logins = append(logins, ClientLoginResponse{
			ClientID: client.ClientID,

			Scope:   scope,
			Created: r.Created,

			Name:        client.Name,
			Description: client.Description,
			Image:       client.Image,
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
		qm.OrderBy("created DESC"),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}

	clients := []*models.Client{}
	clientHistories := []ClientLoginHistoryResponse{}
	for _, history := range histories {
		// キャッシュしてDBから取得する
		var client *models.Client = nil
		for _, mem := range clients {
			if mem.ClientID == history.ClientID {
				client = mem
				break
			}
		}
		if client == nil {
			clientFromDB, err := models.Clients(
				models.ClientWhere.ClientID.EQ(history.ClientID),
			).One(ctx, h.DB)
			if err != nil {
				return err
			}
			clients = append(clients, clientFromDB)
			client = clientFromDB
		}

		clientHistories = append(clientHistories, ClientLoginHistoryResponse{
			ID:       history.ID,
			ClientID: client.ClientID,

			Name:        client.Name,
			Description: client.Description,
			Image:       client.Image,

			Device:   history.Device,
			OS:       history.Os,
			Browser:  history.Browser,
			IsMobile: history.IsMobile,
			Ip:       net.IP.To16(history.IP).String(),

			Created: history.Created,
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

	// SELECT login_history.* FROM login_history
	// INNER JOIN refresh
	//     on refresh.id = login_history.refresh_id
	// WHERE login_history.user_id = ?
	// AND refresh.period > NOW()
	// ORDER BY login_history.created DESC;
	loginDevices, err := models.LoginHistories(
		qm.Select("login_history.*"),
		qm.InnerJoin("refresh on refresh.id = login_history.refresh_id"),
		qm.Where("login_history.user_id", u.ID),
		qm.And("refresh.period > NOW()"),
		qm.OrderBy("login_history.created DESC"),
	).All(ctx, h.DB)
	if err != nil {
		return err
	}

	formattedLoginDevices := []LoginDeviceResponse{}
	for _, l := range loginDevices {
		formattedLoginDevices = append(formattedLoginDevices, LoginDeviceResponse{
			ID: l.ID,

			Device:   l.Device,
			OS:       l.Os,
			Browser:  l.Browser,
			IsMobile: l.IsMobile,
			Ip:       net.IP.To16(l.IP).String(),

			Created: l.Created,
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
		qm.OrderBy("created DESC"),
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

			Created: l.Created,
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
		qm.OrderBy("created DESC"),
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

			Created: l.Created,
		})
	}

	return c.JSON(http.StatusOK, formattedLoginTryHistories)
}
