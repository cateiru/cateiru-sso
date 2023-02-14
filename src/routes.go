package src

import "github.com/labstack/echo/v4"

func Routes(e *echo.Echo, h *Handler) {
	e.GET("/", h.Root)
}
