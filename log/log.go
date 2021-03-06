package log

import (
	"fmt"
	"time"

	"github.com/den19980107/go-fiber-gorm-starter/config"
	"go.uber.org/zap"
)

var Zap *zap.Logger

// Setup Zap logger.
func SetupLogger() {
	var (
		err    error
		zapCfg zap.Config
	)

	switch config.App.Env {
	case "production":
		zapCfg = zap.NewProductionConfig()
	default:
		zapCfg = zap.NewDevelopmentConfig()
	}

	date := time.Now().Format("2006-01-02")

	zapCfg.OutputPaths = []string{
		"stdout",
		fmt.Sprintf(config.App.Log.FilePath, config.App.Env, date),
	}

	Zap, err = zapCfg.Build()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer Zap.Sync()
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Debug(msg string, fields ...zap.Field) {
	Zap.Debug(msg, fields...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Info(msg string, fields ...zap.Field) {
	Zap.Info(msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Warn(msg string, fields ...zap.Field) {
	Zap.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Error(msg string, fields ...zap.Field) {
	Zap.Error(msg, fields...)
}
