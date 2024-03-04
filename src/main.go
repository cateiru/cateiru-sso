package src

import (
	"database/sql"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/net/http2"

	_ "time/tzdata"
)

func Main(mode string, path string) {
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

	if err := Server(config, path); err != nil {
		panic(err)
	}
}

// サーバを実行する
func Server(c *Config, path string) error {
	e := echo.New()
	e.IPExtractor = c.IPExtractor
	e.HTTPErrorHandler = CustomHTTPErrorHandler

	ServerMiddleWare(e, c)

	db, err := sql.Open("mysql", c.DatabaseConfig.FormatDSN())
	if err != nil {
		return err
	}

	handler, err := NewHandler(db, c, path)
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
