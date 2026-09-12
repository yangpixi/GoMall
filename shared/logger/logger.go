package logger

import (
	"log/slog"
	"os"
)

func New(svc string, level slog.Level) *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})).With("service", svc)

	return logger
}
