package market

import (
	"github.com/MultiBanker/broker/pkg/auth"
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/manager/market"
	"github.com/VictoriaMetrics/metrics"
	"github.com/go-chi/chi/v5"
)

type resource struct {
	auther auth.Authenticator
	market market.Marketer
	set    *metrics.Set
}

func NewResource(man manager.Managers) *resource {
	return &resource{
		auther: man.AuthMan,
		set:    man.MetricMan,
		market: man.MarketMan,
	}
}

func (res resource) Route() chi.Router {
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
		r.Route("/auto", func(r chi.Router) {
			r.Post("/connect/", res.autoConnect)
		})
	})

	return r
}
