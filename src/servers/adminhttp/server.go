package adminhttp

import (
	"context"
	"net/http"
)

type adminServer struct {
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
	defer cancel()
	a.server.RegisterOnShutdown(cancel)
	return a.server.ListenAndServe()
}

func (a *adminServer) Stop(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}
