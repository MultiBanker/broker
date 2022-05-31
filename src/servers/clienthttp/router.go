package clienthttp

import (
	"github.com/MultiBanker/broker/pkg/httpresources"
	"github.com/MultiBanker/broker/pkg/metric"
	"github.com/MultiBanker/broker/src/servers/clienthttp/resources/market"
	"github.com/MultiBanker/broker/src/servers/clienthttp/resources/partner"
	"github.com/MultiBanker/broker/src/servers/clienthttp/resources/user/application"
	"github.com/MultiBanker/broker/src/servers/clienthttp/resources/user/auth"
	"github.com/MultiBanker/broker/src/servers/clienthttp/resources/user/recovery"
	"github.com/MultiBanker/broker/src/servers/clienthttp/resources/user/verification"
	"github.com/VictoriaMetrics/metrics"

	"github.com/go-chi/chi/v5"

	"github.com/MultiBanker/broker/src/config"
	"github.com/MultiBanker/broker/src/manager"
)

const (
	ApiPath = "/api/v1"
)

// Routing
// @title                       AutoCredit API
// @version                     1.0
// @host                        api.test.somatic.dev
// @BasePath                    /
// @schemes                     https
// @query.collection.format     multi
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func Routing(opts *config.Config, man manager.Managers) chi.Router {
	r := chi.NewRouter()

	fileRes := httpresources.NewFilesResource("/files", opts.HTTP.FilesDir)
	swaggerRes := httpresources.NewSwaggerResource("/swagger", opts.HTTP.BasePath, "/files")
	versionRes := httpresources.NewVersionResource("/version", opts.Version)

	r.Mount(fileRes.Path(), fileRes.Routes())
	r.Mount(swaggerRes.Path(), swaggerRes.Routes())
	r.Mount(versionRes.Path(), versionRes.Routes())

	mware := metric.NewMetricware(metrics.NewSet())

	// основные роутеры
	r.Route(ApiPath, func(r chi.Router) {
		r.Use(mware.All("/broker")...)
		r.Route("/broker", func(r chi.Router) {
			r.Mount("/partners", partner.NewResource(man).Route())
			r.Mount("/markets", market.NewResource(man).Route())
		})

		r.Route("/users/", func(r chi.Router) {
			r.Mount("/application",
				application.NewResource(man.AuthMan, man.UserApplicationMan).Route())
			r.Mount("/auth",
				auth.NewResource(man.AuthMan, man.UserMan).Route())
			r.Mount("/recovery",
				recovery.NewResource(man.AuthMan, man.RecoveryMan).Route())
			r.Mount("/verify",
				verification.NewResource(man.AuthMan, man.VerifyMan).Route())
		})
	})

	return r
}
