package logger

import (
	"log/slog"
	"os"

	"go-template/internal/config"
)

func InitLogger(cfg *config.Config) (*slog.Logger, error) {
	logOptions := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	var logHandler slog.Handler = slog.NewJSONHandler(os.Stdout, logOptions)

	if cfg.Env == config.Local {
		logOptions.Level = slog.LevelDebug
		logHandler = slog.NewTextHandler(os.Stdout, logOptions)
	}

	if cfg.Env == config.Test {
		logOptions.Level = slog.LevelError
		logHandler = slog.NewTextHandler(os.Stdout, logOptions)
	}

	log := slog.New(logHandler)

	return log, nil
}
