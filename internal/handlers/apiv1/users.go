package apiv1

import (
	"go-template/internal/pkg/logger"
	"log/slog"
	"net/http"
)

type UsersHandler struct {
	log *slog.Logger
}

func NewUsersHandlers(logger *slog.Logger) *UsersHandler {
	return &UsersHandler{
		log: logger.With(slog.String("moduleName", "UsersHandlers")),
	}
}

func (h *UsersHandler) GetMe(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello world"))

	if err != nil {
		h.log.Error("Get error:", logger.Err(err))

		return
	}
}
