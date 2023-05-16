package src

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var L *zap.Logger = nil
var S *zap.SugaredLogger = nil

// Cloud Logging LOG Level Severity
var logLevelSeverity = map[zapcore.Level]string{
	zapcore.DebugLevel:  "DEBUG",
	zapcore.InfoLevel:   "INFO",
	zapcore.WarnLevel:   "WARNING",
	zapcore.ErrorLevel:  "ERROR",
	zapcore.DPanicLevel: "CRITICAL",
	zapcore.PanicLevel:  "ALERT",
	zapcore.FatalLevel:  "EMERGENCY",
}

func EncodeLevel(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(logLevelSeverity[l])
}

func newProductionEncoderConfig() zapcore.EncoderConfig {
	cfg := zap.NewProductionEncoderConfig()

	cfg.TimeKey = "time"
	cfg.LevelKey = "severity"
	cfg.MessageKey = "message"
	cfg.EncodeLevel = EncodeLevel
	cfg.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	return cfg
}

func InitLogging(mode string) {
	var logConfig zap.Config

	switch mode {
	case "test":
		logConfig = zap.NewDevelopmentConfig()
		logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	case "local":
		logConfig = zap.NewDevelopmentConfig()
		logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	case "cloudrun":
		logConfig = zap.NewProductionConfig()
		// Cloud Loggerに対応するための設定
		logConfig.EncoderConfig = newProductionEncoderConfig()
	case "cloudrun-staging":
		logConfig = zap.NewProductionConfig()
		// ステージングではDebugLevel以上のログを出力する
		logConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		// Cloud Loggerに対応するための設定
		logConfig.EncoderConfig = newProductionEncoderConfig()
	default:
		panic(fmt.Sprintf("Unknown mode: %s", mode))
	}

	logger, err := logConfig.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	sugar := logger.Sugar()

	L = logger
	S = sugar
}
