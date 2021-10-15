package clienthttp

import (
	"context"
	"net/http"
)

type clientHttpServer struct {
	certFile string
	keyFile  string
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
	h.server.RegisterOnShutdown(cancel)

	if h.Insecure() {
		if err := h.server.ListenAndServe(); err != nil {
			return err
		}
	}

	if !h.Insecure() {
		if err := h.server.ListenAndServeTLS(h.certFile, h.keyFile); err != nil {
			return err
		}
	}
	panic("SOMETHING WRONG WITH CERT FILES")
}

func (h *clientHttpServer) Insecure() bool {
	return h.keyFile == "" && h.certFile == ""
}

func (h *clientHttpServer) Stop(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}
