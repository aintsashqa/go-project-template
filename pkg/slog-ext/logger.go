package slogext

import (
	"log/slog"
	"os"
)

type Config struct {
	Level int `envconfig:"LOG_LEVEL" default:"-4"`
}

func New(conf Config) *slog.Logger {
	var (
		options    = slog.HandlerOptions{Level: slog.Level(conf.Level)}
		handler    = slog.NewJSONHandler(os.Stderr, &options)
		ctxHandler = NewContextHandler(handler)
		logger     = slog.New(ctxHandler)
	)

	slog.SetDefault(logger)
	return logger
}
