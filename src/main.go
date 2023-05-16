package src

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"golang.org/x/net/http2"

	_ "time/tzdata"
)

func Init(mode string) {
	InitLogging(mode)
}

func Main(mode string) {
	config := InitConfig(mode)

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	time.Local = jst

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

// サーバ設定など
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
				code := http.StatusInternalServerError
				he, ok := v.Error.(*HTTPError)
				echohe, eok := v.Error.(*echo.HTTPError)
				if ok {
					code = he.Code
				} else if eok {
					code = echohe.Code
				}

				// エラーコードが400番台の場合はInfo
				if code >= 400 && code < 500 {
					L.Info("request",
						zap.String("URI", v.URI),
						zap.String("method", v.Method),
						zap.Int("status", code),
						zap.String("host", v.Host),
						zap.String("response_time", time.Since(v.StartTime).String()),
						zap.String("ip", v.RemoteIP),
						zap.String("error_message", v.Error.Error()),
					)
					return nil
				}
				L.Error("request",
					zap.String("URI", v.URI),
					zap.String("method", v.Method),
					zap.Int("status", code),
					zap.String("host", v.Host),
					zap.String("response_time", time.Since(v.StartTime).String()),
					zap.String("ip", v.RemoteIP),
					zap.String("error_message", v.Error.Error()),
				)
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
}
