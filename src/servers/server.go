package servers

import (
	"context"
	"net/http"
	"time"
)

const (
	readTimeout  = 5 * time.Second
	writeTimeout = 30 * time.Second
)

type Server interface {
	Name() string
	Start(ctx context.Context, cancel context.CancelFunc) error
	Stop(ctx context.Context) error
}

type Service struct {
	name   string
	server *http.Server
}

func (s Service) Name() string {
	return s.name
}

func NewService(name, addr string, r http.Handler) *Service {
	return &Service{
		name: name,
		server: &http.Server{
			Addr:         addr,
			Handler:      r,
			WriteTimeout: writeTimeout,
			ReadTimeout:  readTimeout,
		},
	}
}

func (s *Service) Start(_ context.Context, cancel context.CancelFunc) error {
	defer cancel()
	s.server.RegisterOnShutdown(cancel)
	return s.server.ListenAndServe()
}

func (s *Service) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
