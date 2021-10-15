package adminhttp

import (
	"context"
	"net/http"
)

type adminServer struct {
	certFile string
	keyFile  string
	server   *http.Server
}

func NewAdminServer(addr string, r http.Handler) *adminServer {
	return &adminServer{
		server: &http.Server{
			Addr:         addr,
			Handler:      r,
			WriteTimeout: writeTimeout,
			ReadTimeout:  readTimeout,
		},
	}
}

func (a adminServer) Name() string {
	return "admin-server"
}

func (a adminServer) Start(_ context.Context, cancel context.CancelFunc) error {
	a.server.RegisterOnShutdown(cancel)

	if a.Insecure() {
		if err := a.server.ListenAndServe(); err != nil {
			return err
		}
	}

	if !a.Insecure() {
		if err := a.server.ListenAndServeTLS(a.certFile, a.keyFile); err != nil {
			return err
		}
	}
	panic("SOMETHING WRONG WITH CERT FILES")
}

func (a *adminServer) Insecure() bool {
	return a.keyFile == "" && a.certFile == ""
}

func (a *adminServer) Stop(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}
