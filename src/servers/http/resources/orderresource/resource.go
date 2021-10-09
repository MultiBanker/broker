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
		r.Put("/{reference_id}/partners", o.updateClientOrder)
		r.Get("/{reference_id}/partners", o.ordersByReference)
	})

	r.Group(func(r chi.Router) {
		// Admin Api
		r.Get("/", o.orders())
		r.Post("/", o.neworder())
		r.Get("/{id}", o.order())
		r.Put("/{id}", o.updateorder())
	})

	r.Group(func(r chi.Router) {
		// Partner API
		r.Post("/partners", o.updateorderpartner())
	})

	return r
}
