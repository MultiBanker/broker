package orders

import (
	"github.com/MultiBanker/broker/pkg/auth"
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/manager/order"
	"github.com/VictoriaMetrics/metrics"
	"github.com/go-chi/chi/v5"
)

const maxOrderHistoryLimit = 100

type Resource struct {
	authMan  auth.Authenticator
	orderMan order.Orderer
	set      *metrics.Set
}

func NewResource(man manager.Managers) Resource {
	return Resource{
		authMan:  man.AuthMan,
		orderMan: man.OrderMan,
		set:      man.MetricMan,
	}
}

func (o Resource) Route() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		//r.Use(jwtauth.Verifier(o.authMan.TokenAuth()))
		//r.Use(middleware.NewUserAccessCtx(o.authMan.JWTKey()).ChiMiddleware)
		r.Post("/", o.neworder())
		r.Post("/markets", o.marketOrderUpdate)
		r.Get("/{reference_id}/partners", o.ordersByReference)
	})

	return r
}
