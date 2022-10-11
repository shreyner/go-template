package handlers

import (
	"net/http"

	"go.uber.org/zap"
)

type UsersHandler struct {
	log *zap.Logger
}

func NewUsersHandlers(logger *zap.Logger) *UsersHandler {
	return &UsersHandler{
		log: logger.Named("UsersHandlers"),
	}
}

func (h *UsersHandler) GetMe(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte("Hello world"))

	if err != nil {
		h.log.Error("Get error:", zap.Error(err))

		return
	}
}
