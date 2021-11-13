package clienthttp

import (
	"github.com/MultiBanker/broker/pkg/metric"
	"github.com/MultiBanker/broker/src/servers/clienthttp/resources/market"
	"github.com/MultiBanker/broker/src/servers/clienthttp/resources/partner"
	"github.com/VictoriaMetrics/metrics"

	"github.com/go-chi/chi/v5"

	"github.com/MultiBanker/broker/src/config"
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/servers/clienthttp/middleware"
)

const (
	ApiPath = "/api/v1"
)

func Routing(opts *config.Config, man manager.Abstractor) chi.Router {
	r := middleware.Mount(opts.Version, opts.HTTP.FilesDir, opts.HTTP.BasePath)
	mware := metric.NewMetricware(metrics.NewSet())

	// основные роутеры
	r.Route(ApiPath, func(r chi.Router) {
		r.Use(mware.All("/broker")...)
		r.Route("/broker", func(r chi.Router) {
			r.Mount("/partners", partner.NewResource(man).Route())
			r.Mount("/markets", market.NewResource(man).Route())
		})
	})

	return r
}
