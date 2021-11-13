package admin

import (
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/servers/adminhttp/resources/admin/loan"
	"github.com/MultiBanker/broker/src/servers/adminhttp/resources/admin/market"
	"github.com/MultiBanker/broker/src/servers/adminhttp/resources/admin/offer"
	"github.com/MultiBanker/broker/src/servers/adminhttp/resources/admin/orders"
	"github.com/MultiBanker/broker/src/servers/adminhttp/resources/admin/partner"
	"github.com/go-chi/chi/v5"
)

type Resource struct {
	man manager.Abstractor
}

func NewResource(man manager.Abstractor) *Resource {
	return &Resource{man: man}
}

func (res Resource) Route() chi.Router {
	r := chi.NewRouter()

	r.Mount("/markets", market.NewAdminResource(res.man).Route())
	r.Mount("/partners", partner.NewAdminResource(res.man).Route())
	r.Mount("/offers", offer.NewAdminResource(res.man).Route())
	r.Mount("/orders", orders.NewAdminResource(res.man).Route())
	r.Mount("/loan-programs", loan.NewAdminResource(res.man).Route())

	return r
}
