package market

import (
	"github.com/MultiBanker/broker/src/manager/auth"
	"github.com/MultiBanker/broker/src/manager/market"
	"github.com/go-chi/chi"
)

type Resource struct {
	auth   auth.Authenticator
	market market.Marketer
}

func NewResource(auth auth.Authenticator, market market.Marketer) Resource {
	return Resource{
		auth:   auth,
		market: market,
	}
}

func (res Resource) Route() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Post("/", res.create)
		r.Get("/", res.list)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", res.get)
			r.Put("/", res.update)
		})
	})

	return r
}
