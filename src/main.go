package src

import (
	"database/sql"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/net/http2"

	_ "time/tzdata"
)

func Main(mode string) {
	config := InitConfig(mode)
	InitLogging(mode, config)

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	time.Local = jst

	if config.DBDebugLog {
		stdOutLogger, err := zap.NewStdLogAt(L, zapcore.DebugLevel)
		if err != nil {
			panic(err)
		}

		// SQLBoilerのログも出力する
		boil.DebugMode = true
		boil.DebugWriter = stdOutLogger.Writer()
	}

	// CloudStorageのために環境変数を設定する
	if config.StorageEmulatorHost.IsValid {
		os.Setenv("STORAGE_EMULATOR_HOST", config.StorageEmulatorHost.Value)
	}

	if err := Server(config); err != nil {
		panic(err)
	}
}

// サーバを実行する
func Server(c *Config) error {
	e := echo.New()
	e.IPExtractor = echo.ExtractIPFromXFFHeader()
	e.HTTPErrorHandler = CustomHTTPErrorHandler

	ServerMiddleWare(e, c)

	db, err := sql.Open("mysql", c.DatabaseConfig.FormatDSN())
	if err != nil {
		return err
	}

	handler, err := NewHandler(db, c)
	if err != nil {
		return err
	}
	// APIのルート設定
	Routes(e, handler, c)

	s := http2.Server{}

	// Start a server
	// connection port is `8080`
	//
	// and, `http://localhist:8080` to access.
	return e.StartH2CServer(":8080", &s)
}

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

// CSRF対策で`Sec-Fetch-Site`ヘッダを検証する
func CSRFHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		secFetchSiteHeader := c.Request().Header.Get("Sec-Fetch-Site")

		// `same-origin`か`same-site`以外の場合はCSRFエラーを出す
		if secFetchSiteHeader != "same-origin" && secFetchSiteHeader != "same-site" {
			return NewHTTPError(403, "CSRF Error")
		}

		return next(c)
	}
}
