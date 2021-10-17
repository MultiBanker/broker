package adminhttp

import (
	"github.com/MultiBanker/broker/pkg/metric"
	"github.com/MultiBanker/broker/src/config"
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/servers/adminhttp/middleware"
	"github.com/MultiBanker/broker/src/servers/adminhttp/resources/admin"
	"github.com/VictoriaMetrics/metrics"
	"github.com/go-chi/chi/v5"
)

const (
	ApiPath      = "/api/v1"
)

func Routing(opts *config.Config, man manager.Abstractor) chi.Router {

	r := middleware.Mount(opts.Version, opts.HTTP.Admin.FilesDir, opts.HTTP.Admin.BasePath)
	mware := metric.NewMetricware(metrics.NewSet())

	// основные роутеры
	r.Route(ApiPath, func(r chi.Router) {
		r.Use(mware.All("/broker")...)
		r.Route("/broker", func(r chi.Router) {
			r.Mount("/admins", admin.NewResource(man).Route())
		})
	})

	return r
}