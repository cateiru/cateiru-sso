package logging

import (
	"os"

	"go.uber.org/zap"
)

var Logger *zap.Logger
var Sugar *zap.SugaredLogger

// ログモードを選択する
func init() {
	switch os.Getenv("DEPLOY_MODE") {
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
