package adminhttp

import (
	"github.com/MultiBanker/broker/pkg/metric"
	"github.com/MultiBanker/broker/src/config"
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/servers/adminhttp/middleware"
	"github.com/MultiBanker/broker/src/servers/adminhttp/resources/admin"
	"github.com/go-chi/chi/v5"
)

const (
	ApiPath = "/api/v1"
)

func Routing(opts *config.Config, man manager.Wrapper) chi.Router {

	r := middleware.Mount(opts.Version, opts.HTTP.FilesDir, opts.HTTP.BasePath)
	mware := metric.NewMetricware(man.Metric())

	// основные роутеры
	r.Route(ApiPath, func(r chi.Router) {
		r.Use(mware.All("/broker")...)
		r.Route("/broker", func(r chi.Router) {
			r.Mount("/admins", admin.NewResource(man).Route())
		})
	})

	return r
}
