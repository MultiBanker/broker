package orders

import (
	"github.com/MultiBanker/broker/pkg/auth"
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/servers/clienthttp/middleware"
	"github.com/VictoriaMetrics/metrics"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"

	"github.com/MultiBanker/broker/src/manager/order"
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
		r.Use(jwtauth.Verifier(o.authMan.TokenAuth()))
		r.Use(middleware.NewUserAccessCtx(o.authMan.JWTKey()).ChiMiddleware)
		// Partner API
		r.Post("/", o.updateorderpartner())
	})

	return r
}
