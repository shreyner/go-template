package app

import (
	"context"
	"database/sql"
	"go-template/internal/handlers/apiv1"
	"go-template/internal/pkg/database"
	"go-template/internal/pkg/logger"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"go-template/internal/config"
	"go-template/internal/pkg/httpserver"
	"go-template/internal/router"
)

func Run(log *slog.Logger, cfg *config.Config) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Info("Starting application", slog.String("env", cfg.Env))

	log.Info("Connection to database...")
	db, err := database.New(&cfg.DataBase)

	if err != nil {
		log.Error("Can't connection to db", logger.Err(err))

		os.Exit(1)
		return
	}

	defer func(db *sql.DB) {
		log.Info("Close database connection")
		err := db.Close()
		if err != nil {
			log.Error("Error to close connection db", logger.Err(err))
		}
	}(db)

	log.Info("Staring the application...")

	router := router.New(log)
	routerAPIv1 := apiv1.New(log)

	router.Mount("/api/v1/", routerAPIv1)

	apiServer := httpserver.NewHttpServer(log, router, &cfg.Http)

	log.Info("Staring rest api server...")
	if err := apiServer.Start(ctx); err != nil {
		log.Error("can't start rest api server", logger.Err(err))
		return
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case x := <-interrupt:
		log.Info("Received a signal.", slog.String("signal", x.String()))
	case err := <-apiServer.Notify():
		log.Error("Received an error from the start rest api server", logger.Err(err))
	}

	log.Info("Stopping server...")

	if err := apiServer.Stop(ctx); err != nil {
		log.Error("Got an error while stopping th rest api server", logger.Err(err))
	}

	log.Info("The app is calling the last defers and will be stopped.")
}
