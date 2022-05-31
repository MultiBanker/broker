package orders

import (
	"github.com/MultiBanker/broker/pkg/auth"
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/manager/order"
	"github.com/VictoriaMetrics/metrics"
	"github.com/go-chi/chi/v5"
)

const maxOrderHistoryLimit = 100

type resource struct {
	authMan  auth.Authenticator
	orderMan order.Orderer
	set      *metrics.Set
}

func NewResource(man manager.Managers) resource {
	return resource{
		authMan:  man.AuthMan,
		orderMan: man.OrderMan,
		set:      man.MetricMan,
	}
}

func (o resource) Route() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		//r.Use(jwtauth.Verifier(o.authMan.TokenAuth()))
		//r.Use(middleware.NewUserAccessCtx(o.authMan.JWTKey()).ChiMiddleware)
		// Admin Api
		r.Get("/", o.orders())
		r.Get("/{id}", o.order())
		r.Put("/{id}", o.updateorder())
	})

	return r
}
