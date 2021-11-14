package orders

import (
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/manager/auth"
	"github.com/MultiBanker/broker/src/manager/order"
	"github.com/VictoriaMetrics/metrics"
	"github.com/go-chi/chi/v5"
)

const maxOrderHistoryLimit = 100

type AdminResource struct {
	authMan  auth.Authenticator
	orderMan order.Orderer
	set      *metrics.Set
}

func NewAdminResource(man manager.Wrapper) AdminResource {
	return AdminResource{
		authMan:  man.Auther(),
		orderMan: man.Orderer(),
		set:      man.Metric(),
	}
}

func (o AdminResource) Route() chi.Router {
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
