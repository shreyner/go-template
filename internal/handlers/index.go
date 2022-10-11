package handlers

import (
	"net/http"

	"go.uber.org/zap"
)

type IndexHandler struct {
	log *zap.Logger
}

func NewIndexHandlers(logger *zap.Logger) *IndexHandler {
	return &IndexHandler{
		log: logger.Named("IndexHandlers"),
	}
}

func (i *IndexHandler) Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("index"))
}
