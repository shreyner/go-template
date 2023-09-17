package apiv1

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func New(log *slog.Logger) http.Handler {
	r := chi.NewRouter()

	usersHandler := NewUsersHandlers(log)

	r.Get("/me", usersHandler.GetMe)

	return r
}
