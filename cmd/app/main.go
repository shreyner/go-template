package main

import (
	"fmt"
	"os"

	"go-template/internal/app"
	"go-template/internal/config"
	"go-template/internal/pkg/logger"
)

func main() {
	cfg := &config.Config{}

	if err := cfg.Parse(); err != nil {
		fmt.Errorf("Can't parse env: %w", err)
		os.Exit(1)
	}

	log, err := logger.InitLogger(cfg)
	if err != nil {
		_ = fmt.Errorf("error initilizing logger: %w", err)
		os.Exit(1)
	}

	defer log.Sync()

	app.Run(log, cfg)
}
