package src

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	DB *sql.DB
	C  *Config
}

func NewHandler(db *sql.DB, config *Config) *Handler {
	return &Handler{
		DB: db,
		C:  config,
	}
}

func (h *Handler) Root(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"mode": h.C.Mode,
	})
}
