package main

import (
	"context"
	"database/sql"
	"fmt"
	"go-template/internal/pkg/database"
	"os"
	"os/signal"
	"syscall"

	"go-template/internal/config"
	"go-template/internal/pkg/logger"
	"go-template/internal/router"
	"go-template/internal/server"

	"go.uber.org/zap"
)

func main() {
	log, err := logger.InitLogger()
	if err != nil {
		fmt.Errorf("error initilizing logger: %w", err)
		os.Exit(1)
	}

	defer log.Sync()

	log.Info("Parsing env")
	cfg := config.Config{}

	if err := cfg.Parse(); err != nil {
		log.Fatal("Can't parse env", zap.Error(err))
		os.Exit(1)
	}

	log.Info("Connection to database...")
	db, err := database.New(cfg.DBUrl)

	if err != nil {
		log.Fatal("Can't connection to db", zap.Error(err))
		os.Exit(1)
	}
	defer func(db *sql.DB) {
		log.Info("Close database connection")
		err := db.Close()
		if err != nil {
			log.Error("Error to close connection db", zap.Error(err))
		}
	}(db)

	log.Info("Staring the application...")

	apiMux := router.New(log)

	apiServer := server.New(log, apiMux, cfg.Port)
	log.Info("Staring rest api server...")
	go apiServer.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case x := <-interrupt:
		log.Info("Received a signal.", zap.String("signal", x.String()))
	case err := <-apiServer.Notify():
		log.Error("Received an error from the start rest api server", zap.Error(err))
	}

	log.Info("Stopping server...")

	if err := apiServer.Stop(context.Background()); err != nil {
		log.Error("Got an error while stopping th rest api server", zap.Error(err))
	}

	log.Info("The app is calling the last defers and will be stopped.")
}
