package logger

import (
	"log/slog"
	"os"
	"sync"
)

var (
	loggerOnce sync.Once
	loggerGlob *slog.Logger
)

type LoggerInterface interface {
}

func GetProdLogger() *slog.Logger {

	loggerOnce.Do(func() {
		loggerGlob = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	})

	return loggerGlob
}
