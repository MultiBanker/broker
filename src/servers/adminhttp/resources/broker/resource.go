package broker

import (
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/servers/adminhttp/resources/broker/loan"
	"github.com/MultiBanker/broker/src/servers/adminhttp/resources/broker/offer"
	"github.com/MultiBanker/broker/src/servers/adminhttp/resources/broker/orders"
	"github.com/MultiBanker/broker/src/servers/adminhttp/resources/broker/partner"
	"github.com/go-chi/chi/v5"
)

type Resource struct {
	man manager.Managers
}

func NewResource(man manager.Managers) *Resource {
	return &Resource{man: man}
}

func (res Resource) Route() chi.Router {
	r := chi.NewRouter()

	r.Mount("/partners", partner.NewResource(res.man).Route())
	r.Mount("/offers", offer.NewResource(res.man).Route())
	r.Mount("/orders", orders.NewResource(res.man).Route())
	r.Mount("/loan-programs", loan.NewResource(res.man).Route())

	return r
}
