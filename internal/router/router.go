package router

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"go-template/internal/handlers"
	"go.uber.org/zap"
)

func New(log *zap.Logger) *chi.Mux {
	log.Info("Initilize REST API")
	userH := handlers.NewUsersHandlers(log)

	r := chi.NewRouter()

	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	r.Get("/", Index)
	r.Get("/me", userH.GetMe)

	return r
}
