package logging

import (
	"github.com/cateiru/cateiru-sso/api/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger
var Sugar *zap.SugaredLogger

// ログモードを選択する
func init() {
	switch utils.DEPLOY_MODE {
	case "production":
		loggerProduction, err := zap.NewProduction()
		if err != nil {
			panic(err)
		}
		Logger = loggerProduction
	default:
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		loggingDebug, err := config.Build()
		if err != nil {
			panic(err)
		}
		Logger = loggingDebug
	}

	Sugar = Logger.Sugar()
}
