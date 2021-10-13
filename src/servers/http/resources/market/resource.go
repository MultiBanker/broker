package market

import (
	"github.com/MultiBanker/broker/src/manager/auth"
	"github.com/MultiBanker/broker/src/manager/market"
	"github.com/VictoriaMetrics/metrics"
	"github.com/go-chi/chi/v5"
)

type Resource struct {
	auther auth.Authenticator
	market market.Marketer
	set    *metrics.Set
}

func NewResource(auth auth.Authenticator, market market.Marketer, set *metrics.Set) Resource {
	return Resource{
		auther: auth,
		market: market,
		set:    set,
	}
}

func (res Resource) Route() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Post("/login", res.auth())
		r.Get("/logout", res.out())
	})

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
