package main

import (
	"github.com/go-chi/chi"
	"google.golang.org/grpc"

	"github.com/MultiBanker/broker/src/config"
	"github.com/MultiBanker/broker/src/director"
	"github.com/MultiBanker/broker/src/manager"
	grpcsrv "github.com/MultiBanker/broker/src/servers/grpc"
	"github.com/MultiBanker/broker/src/servers/http"
)

func (a *application) services() {
	a.workers()
	a.httpserver(http.Routing)
	a.grpcserver(grpcsrv.Routing)
}

func (a *application) workers(daemons ...director.Daemons) {
	for _, daemon := range daemons {
		a.servers = append(a.servers, daemon)
	}
}

func (a *application) httpserver(fn func(opts *config.Config, man manager.Abstractor) chi.Router) {
	srv := http.NewHTTP(a.opts.HTTP.ListenAddr, fn(a.opts, a.man))
	a.servers = append(a.servers, srv)
}

func (a *application) grpcserver(fn func(server *grpc.Server, man manager.Abstractor)) {
	srv := grpcsrv.NewGRPC(a.opts, a.man, fn)
	a.servers = append(a.servers, srv)
}
