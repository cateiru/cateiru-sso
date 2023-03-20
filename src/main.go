package src

import (
	"database/sql"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
)

func Init(mode string) {
	InitLogging(mode)
}

func Main(mode string) {
	config := InitConfig(mode)

	if err := Server(config); err != nil {
		S.Fatal(err)
	}
}

// サーバを実行する
func Server(c *Config) error {
	e := echo.New()
	e.IPExtractor = echo.ExtractIPFromXFFHeader()
	e.HTTPErrorHandler = CustomHTTPErrorHandler

	ServerMiddleWare(e)

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

	// オレオレ証明書で実行する
	// ローカル環境用
	if c.SelfSignedCert {
		return e.StartTLS(":8080", "./certificates/tls.crt", "./certificates/tls.key")
	}

	s := http2.Server{}

	// Start a server
	// connection port is `8080`
	//
	// and, `http://localhist:8080` to access.
	return e.StartH2CServer(":8080", &s)
}

// サーバ設定など
func ServerMiddleWare(e *echo.Echo) {
	// リクエストごとにログを出す
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogMethod: true,
		LogHost:   true,
		LogError:  true,

		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				// エラーコードが400番台の場合はInfo
				if v.Status >= 400 && v.Status < 500 {
					L.Info("request",
						zap.String("URI", v.URI),
						zap.String("method", v.Method),
						zap.Int("status", v.Status),
						zap.String("host", v.Host),
						zap.String("response_time", time.Since(v.StartTime).String()),
						zap.String("error_message", v.Error.Error()),
					)
					return nil
				}
				L.Error("request",
					zap.String("URI", v.URI),
					zap.String("method", v.Method),
					zap.Int("status", v.Status),
					zap.String("host", v.Host),
					zap.String("response_time", time.Since(v.StartTime).String()),
					zap.String("error_message", v.Error.Error()),
				)
			} else {
				L.Info("request",
					zap.String("URI", v.URI),
					zap.String("method", v.Method),
					zap.Int("status", v.Status),
					zap.String("host", v.Host),
					zap.String("response_time", time.Since(v.StartTime).String()),
				)
			}
			return nil
		},
	}))
}
