package clienthttp

import (
	"context"
	"net/http"
)

type clientHttpServer struct {
	server   *http.Server
}

func NewClientHTTP(addr string, r http.Handler) *clientHttpServer {
	return &clientHttpServer{
		server: &http.Server{
			Addr:         addr,
			Handler:      r,
			WriteTimeout: writeTimeout,
			ReadTimeout:  readTimeout,
		},
	}
}

func (h *clientHttpServer) Name() string {
	return "client-server"
}

func (h *clientHttpServer) Start(_ context.Context, cancel context.CancelFunc) error {
	defer cancel()
	h.server.RegisterOnShutdown(cancel)
	return h.server.ListenAndServe()
}

func (h *clientHttpServer) Stop(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}
