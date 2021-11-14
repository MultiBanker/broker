package market

import (
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/servers/clienthttp/resources/market/auth"
	"github.com/MultiBanker/broker/src/servers/clienthttp/resources/market/offer"
	"github.com/MultiBanker/broker/src/servers/clienthttp/resources/market/orders"
	"github.com/go-chi/chi/v5"
)

type Resource struct {
	man manager.Wrapper
}

func NewResource(man manager.Wrapper) Resource {
	return Resource{
		man: man,
	}
}

func (res Resource) Route() chi.Router {
	r := chi.NewRouter()

	r.Mount("/", auth.NewResource(res.man).Route())
	r.Mount("/orders", orders.NewResource(res.man).Route())
	r.Mount("/offers", offer.NewResource(res.man).Route())

	return r
}
