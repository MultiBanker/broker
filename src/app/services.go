package main

import (
	"github.com/MultiBanker/broker/src/config"
	"github.com/MultiBanker/broker/src/director"
	"github.com/MultiBanker/broker/src/director/notifier"
	"github.com/MultiBanker/broker/src/director/order"
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/servers"
	"github.com/MultiBanker/broker/src/servers/adminhttp"
	"github.com/MultiBanker/broker/src/servers/clienthttp"
	"github.com/MultiBanker/broker/src/servers/healthz"
	"github.com/MultiBanker/broker/src/servers/victoriaMetrics"
	"github.com/go-chi/chi/v5"
)

func (a *application) services() {
	a.workers(
		a.orderDeadliner(),
		a.notifier(),
	)
	a.healthServer(healthz.Routing)
	a.clienthttpServer(clienthttp.Routing)
	a.adminhttpserver(adminhttp.Routing)
	a.victoriaMetricsServer(victoriaMetrics.Routing)
}

func (a *application) workers(daemons ...director.Daemons) {
	for _, daemon := range daemons {
		a.servers = append(a.servers, daemon)
	}
}

func (a *application) clienthttpServer(fn func(opts *config.Config, man manager.Managers) chi.Router) {
	srv := servers.NewService("client-broker-http", a.opts.HTTP.Client.ListenAddr, fn(a.opts, a.man))
	a.servers = append(a.servers, srv)
}

func (a *application) adminhttpserver(fn func(opts *config.Config, man manager.Managers) chi.Router) {
	srv := servers.NewService("broker-broker-http", a.opts.HTTP.Admin.ListenAddr, fn(a.opts, a.man))
	a.servers = append(a.servers, srv)
}

func (a *application) victoriaMetricsServer(fn func() chi.Router) {
	srv := servers.NewService("victoria-metrics", a.opts.VictoriaMetrics.ListenAddr, fn())
	a.servers = append(a.servers, srv)
}

func (a *application) healthServer(fn func(man manager.Managers) chi.Router) {
	srv := servers.NewService("health-server", a.opts.HTTP.HealthPort, fn(a.man))
	a.servers = append(a.servers, srv)
}

func (a *application) orderDeadliner() director.Daemons {
	return order.NewDeadline(a.repo.PartnerOrder)
}

func (a *application) notifier() director.Daemons {
	return notifier.NewNotifier(a.repo.User, a.repo.Verify, a.repo.Recovery, *a.opts.NotifyConfig)
}
