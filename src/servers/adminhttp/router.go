package adminhttp

import (
	"github.com/MultiBanker/broker/pkg/metric"
	"github.com/MultiBanker/broker/src/config"
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/servers/adminhttp/middleware"
	"github.com/MultiBanker/broker/src/servers/adminhttp/resources/auto"
	"github.com/MultiBanker/broker/src/servers/adminhttp/resources/broker"
	"github.com/MultiBanker/broker/src/servers/adminhttp/resources/market"
	"github.com/MultiBanker/broker/src/servers/adminhttp/resources/user/application"
	"github.com/go-chi/chi/v5"
)

const (
	ApiPath = "/api/v1"
)

// Routing
// @title                       AutoCredit API
// @version                     1.0
// @host                        admin-api.test.somatic.dev
// @BasePath                    /
// @schemes                     https
// @query.collection.format     multi
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func Routing(opts *config.Config, man manager.Managers) chi.Router {

	r := middleware.Mount(opts.Version, opts.HTTP.FilesDir, opts.HTTP.BasePath)
	mware := metric.NewMetricware(man.MetricMan)

	// основные роутеры
	r.Route(ApiPath, func(r chi.Router) {
		r.Use(mware.All("/brokers")...)

		r.Mount("/brokers", broker.NewResource(man).Route())
		r.Mount("/markets", market.NewResource(man).Route())
		r.Mount("/auto", auto.NewResource(man.AuthMan, man.AutoMan).Route())
		r.Route("/users/", func(r chi.Router) {
			r.Mount("/application",
				application.NewResource(man.AuthMan, man.UserApplicationMan).Route())
		})
	})

	return r
}
