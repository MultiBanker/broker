package partner

import (
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/servers/clienthttp/resources/partner/auth"
	"github.com/MultiBanker/broker/src/servers/clienthttp/resources/partner/orders"
	"github.com/go-chi/chi/v5"
)

const maxOrderHistoryLimit = 100

type Resource struct {
	man manager.Managers
}

func NewResource(man manager.Managers) Resource {
	return Resource{
		man: man,
	}
}

func (a Resource) Route() chi.Router {
	r := chi.NewRouter()

	r.Mount("/", auth.NewResource(a.man).Route())
	r.Mount("/orders", orders.NewResource(a.man).Route())

	return r
}
