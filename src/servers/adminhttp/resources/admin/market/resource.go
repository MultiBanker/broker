package market

import (
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/manager/auth"
	"github.com/MultiBanker/broker/src/manager/market"
	"github.com/VictoriaMetrics/metrics"
	"github.com/go-chi/chi/v5"
)

type AdminResource struct {
	auther auth.Authenticator
	market market.Marketer
	set    *metrics.Set
}

func NewAdminResource(man manager.Wrapper) AdminResource {
	return AdminResource{
		auther: man.Auther(),
		set:    man.Metric(),
		market: man.Marketer(),
	}
}

func (res AdminResource) Route() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		//r.Use(jwtauth.Verifier(a.authMan.TokenAuth()))
		//r.Use(middleware.NewUserAccessCtx(a.authMan.JWTKey()).ChiMiddleware)
		r.Post("/", res.create)
		r.Get("/", res.list)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", res.get)
			r.Put("/", res.update)
		})
	})

	return r
}
