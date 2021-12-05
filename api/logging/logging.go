package logging

import (
	"github.com/cateiru/cateiru-sso/api/utils"
	"go.uber.org/zap"
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
		loggingDebug, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		Logger = loggingDebug
	}

	Sugar = Logger.Sugar()
}
