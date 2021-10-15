package main

import (
	"github.com/MultiBanker/broker/src/director/order"
	"github.com/MultiBanker/broker/src/servers/adminhttp"
	"github.com/MultiBanker/broker/src/servers/victoriaMetrics"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"

	"github.com/MultiBanker/broker/src/config"
	"github.com/MultiBanker/broker/src/director"
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/servers/clienthttp"
	grpcsrv "github.com/MultiBanker/broker/src/servers/grpc"
)

func (a *application) services() {
	a.workers(
		a.orderDeadliner(),
		)
	a.clienthttpServer(clienthttp.Routing)
	a.adminhttpserver(adminhttp.Routing)
	a.grpcserver(grpcsrv.Routing)
	a.victoriaMetricsServer()
}

func (a *application) workers(daemons ...director.Daemons) {
	for _, daemon := range daemons {
		a.servers = append(a.servers, daemon)
	}
}

func (a *application) clienthttpServer(fn func(opts *config.Config, man manager.Abstractor) chi.Router) {
	srv := clienthttp.NewClientHTTP(a.opts.HTTP.Client.ListenAddr, fn(a.opts, a.man))
	a.servers = append(a.servers, srv)
}

func (a *application) grpcserver(fn func(server *grpc.Server, man manager.Abstractor)) {
	srv := grpcsrv.NewGRPC(a.opts, a.man, fn)
	a.servers = append(a.servers, srv)
}

func (a *application) adminhttpserver(fn func(opts *config.Config, man manager.Abstractor) chi.Router) {
	srv := adminhttp.NewAdminServer(a.opts.HTTP.Admin.ListenAddr, fn(a.opts, a.man))
	a.servers = append(a.servers, srv)
}

func (a *application) victoriaMetricsServer() {
	srv := victoriaMetrics.NewVictoriaM(a.metric, a.opts.VictoriaMetrics.ListenAddr, victoriaMetrics.Routing())
	a.servers = append(a.servers, srv)
}

func (a *application) orderDeadliner() director.Daemons {
	return order.NewDeadline(a.repo.PartnerOrderRepo())
}
