package logger

import (
	"time"

	"go.uber.org/zap"
)

var zapLog *zap.SugaredLogger

func NewLogger() *zap.SugaredLogger {
	url := ""
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync() // flushes buffer, if any
	zapLog := logger.Sugar()
	zapLog.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	zapLog.Infof("Failed to fetch URL: %s", url)
	return zapLog
}

func Sync() {
	zapLog.Sync()
}
