package orderresource

import (
	"github.com/go-chi/chi"

	"github.com/MultiBanker/broker/src/manager/order"
)

const maxOrderHistoryLimit = 100

type Order struct {
	orderMan order.Orderer
}

func NewOrder(orderMan order.Orderer) Order {
	return Order{
		orderMan: orderMan,
	}
}

func (o Order) Route() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Post("/", o.neworder())
		r.Put("/{id}", o.updateorder())
		r.Get("/{id}", o.order())
		r.Get("/", o.orders())
		r.Get("/reference/{id}", o.referenceOrders())

	})

	return r
}


