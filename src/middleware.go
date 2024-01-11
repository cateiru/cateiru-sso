package src

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

// 共通のミドルウェア
func ServerMiddleWare(e *echo.Echo, c *Config) {
	// リクエストごとにログを出す
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:      true,
		LogStatus:   true,
		LogMethod:   true,
		LogHost:     true,
		LogError:    true,
		LogRemoteIP: true,

		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				ErrorLog(v)
			} else {
				L.Info("request",
					zap.String("URI", v.URI),
					zap.String("method", v.Method),
					zap.Int("status", v.Status),
					zap.String("host", v.Host),
					zap.String("response_time", time.Since(v.StartTime).String()),
					zap.String("ip", v.RemoteIP),
				)
			}
			return nil
		},
	}))

	// CORS設定
	e.Use(middleware.CORSWithConfig(*c.CorsConfig))

	// gzip設定
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 6,
	}))
}

// Basic認証
func BasicAuthMiddleware(c *Config) echo.MiddlewareFunc {
	return middleware.BasicAuth(func(userName, password string, ctx echo.Context) (bool, error) {
		if userName == c.InternalBasicAuthUserName && password == c.InternalBasicAuthPassword {
			return true, nil
		}
		return false, nil
	})
}

// CSRF対策で`Sec-Fetch-Site`ヘッダを検証する
func CSRFMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		secFetchSiteHeader := c.Request().Header.Get("Sec-Fetch-Site")

		// `same-origin`か`same-site`以外の場合はCSRFエラーを出す
		if secFetchSiteHeader != "same-origin" && secFetchSiteHeader != "same-site" {
			return NewHTTPError(403, "CSRF Error")
		}

		return next(c)
	}
}

// FedCMの`Sec-Fetch-Dest`ヘッダーを検証する
func FedCMMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		secFetchDest := c.Request().Header.Get("Sec-Fetch-Dest")

		// `webidentity`以外の場合はエラーを出す
		if secFetchDest != "webidentity" {
			return NewHTTPError(401, "Unauthorized")
		}

		return next(c)
	}
}
