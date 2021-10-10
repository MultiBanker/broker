package orderresource

import (
	"github.com/MultiBanker/broker/src/manager/auth"
	"github.com/MultiBanker/broker/src/servers/http/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"

	"github.com/MultiBanker/broker/src/manager/order"
)

const maxOrderHistoryLimit = 100

type Order struct {
	authMan  auth.Authenticator
	orderMan order.Orderer
}

func NewOrder(authMan auth.Authenticator, orderMan order.Orderer) Order {
	return Order{
		authMan:  authMan,
		orderMan: orderMan,
	}
}

func (o Order) Route() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(o.authMan.TokenAuth()))
		r.Use(middleware.NewUserAccessCtx(o.authMan.JWTKey()).ChiMiddleware)
		r.Post("/", o.neworder())
		r.Post("/markets/", o.marketOrderUpdate)
		r.Get("/{reference_id}/partners", o.ordersByReference)
	})

	r.Group(func(r chi.Router) {
		// Admin Api
		r.Get("/{id}", o.order())
		r.Put("/{id}", o.updateorder())
	})

	r.Group(func(r chi.Router) {
		// Partner API
		r.Post("/partners", o.updateorderpartner())
	})

	return r
}
