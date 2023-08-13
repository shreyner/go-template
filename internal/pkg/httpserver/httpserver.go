package httpserver

import (
	"context"
	"go-template/internal/pkg/server"
	"log/slog"
	"net"
	"net/http"
	"strconv"
	"time"
)

var _ server.Server = (*HTTPServer)(nil)

type HTTPServer struct {
	server http.Server
	errors chan error
	log    *slog.Logger
}

func NewHttpServer(log *slog.Logger, handler http.Handler, config *HttpConfig) *HTTPServer {
	return &HTTPServer{
		log: log,
		server: http.Server{
			Addr:    net.JoinHostPort("", strconv.Itoa(config.Port)),
			Handler: handler,
		},
		errors: make(chan error),
	}
}

func (hs *HTTPServer) Start(_ context.Context) error {
	go func() {
		hs.log.Info("Http Server listening on ", slog.String("addr", hs.server.Addr))
		hs.errors <- hs.server.ListenAndServe()
		close(hs.errors)
	}()

	return nil
}

func (hs *HTTPServer) Stop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return hs.server.Shutdown(ctx)
}

func (hs *HTTPServer) Notify() <-chan error {
	return hs.errors
}
