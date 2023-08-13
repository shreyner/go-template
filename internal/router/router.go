package router

import (
	"log/slog"
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"

	"go-template/internal/pkg/logger"
)

func New(log *slog.Logger) *chi.Mux {
	log.Info("Init REST API")

	r := chi.NewRouter()

	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(logger.NewStructuredLogger(log))
	r.Use(chiMiddleware.Recoverer)
	r.Use(chiMiddleware.URLFormat)

	r.Use(chiMiddleware.Timeout(60 * time.Second))

	// // TODO: https://github.com/swaggo/http-swagger
	// r.Get("/swagger/*", httpSwagger.Handler())

	return r
}
