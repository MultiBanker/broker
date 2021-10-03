package http

import (
	"context"
	"log"
	"net/http"
)

type httpServer struct {
	certFile string
	keyFile  string
	server   *http.Server
}

func NewHTTP(addr string, r http.Handler) *httpServer {
	return &httpServer{
		server: &http.Server{
			Addr:         addr,
			Handler:      r,
			WriteTimeout: writeTimeout,
			ReadTimeout:  readTimeout,
		},
	}
}

func (h *httpServer) Name() string {
	return "HTTP"
}

func (h *httpServer) Start(ctx context.Context, cancel context.CancelFunc) error {
	h.server.RegisterOnShutdown(cancel)

	log.Printf("[INFO] Starting http server on %s", h.server.Addr)

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

func (h *httpServer) Insecure() bool {
	return h.keyFile == "" && h.certFile == ""
}

func (h *httpServer) Stop(ctx context.Context) error {
	<-ctx.Done()
	return h.server.Shutdown(context.Background())
}
