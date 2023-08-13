package logger

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

var _ middleware.LogEntry = (*StructuredLoggerEntry)(nil)

type StructuredLoggerEntry struct {
	Logger *slog.Logger
}

func (s *StructuredLoggerEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	log := s.Logger.With(
		slog.Int("respStatus", status),
		slog.Int("respBytesLength", bytes),
		slog.String("respDuration", elapsed.String()),
	)

	switch {
	case status <= http.StatusBadRequest:
		log.Info("request complete")
	case status <= http.StatusInternalServerError:
		log.Warn("request complete")
	default:
		log.Error("request complete")
	}
}

func (s *StructuredLoggerEntry) Panic(v interface{}, stack []byte) {
	s.Logger = s.Logger.With(
		slog.String("stack", string(stack)),
		slog.String("panic", fmt.Sprintf("%+v", v)),
	)
}

var _ middleware.LogFormatter = (*StructLogger)(nil)

type StructLogger struct {
	Logger *slog.Logger
}

func (s *StructLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &StructuredLoggerEntry{
		Logger: s.Logger,
	}

	schema := "http"
	if r.TLS != nil {
		schema = "https"
	}

	entry.Logger = entry.Logger.WithGroup("http").With(
		slog.String("schema", schema),
		slog.String("method", r.Method),
		slog.String("remoteAddr", r.RemoteAddr),
		slog.String("userAgent", r.UserAgent()),
		slog.String("uri", fmt.Sprintf("%s://%s%s", schema, r.Host, r.RequestURI)),
	)

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		entry.Logger = entry.Logger.With(
			slog.String("reqID", reqID),
		)
	}

	return entry
}

func NewStructuredLogger(log *slog.Logger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&StructLogger{Logger: log})
}
