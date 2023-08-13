package apiv1

import (
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
)

func New(log *slog.Logger) http.Handler {
	r := chi.NewRouter()

	usersHandler := NewUsersHandlers(log)

	r.Get("/me", usersHandler.GetMe)

	return r
}
