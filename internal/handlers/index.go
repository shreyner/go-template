package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
)

type IndexHandler struct {
	log *slog.Logger
}

func NewIndexHandlers(logger *slog.Logger) *IndexHandler {
	return &IndexHandler{
		log: logger.With(slog.String("moduleName", "IndexHandlers")),
	}
}

func (i *IndexHandler) Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello world")
}
