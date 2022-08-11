package logger

import (
	"fmt"
	"go.uber.org/zap"
)

func InitLogger() (*zap.Logger, error) {

	cfg := zap.NewDevelopmentConfig()

	cfg.DisableStacktrace = true

	logger, err := cfg.Build()

	if err != nil {
		return nil, fmt.Errorf("logger build: %w", err)
	}

	return logger, nil
}
