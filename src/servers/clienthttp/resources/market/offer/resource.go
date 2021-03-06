package offer

import (
	"github.com/MultiBanker/broker/pkg/auth"
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/manager/offer"
	"github.com/VictoriaMetrics/metrics"
	"github.com/go-chi/chi/v5"
)

type Resource struct {
	authMan  auth.Authenticator
	offerMan offer.Manager
	set      *metrics.Set
}

func NewResource(man manager.Managers) *Resource {
	return &Resource{
		authMan:  man.AuthMan,
		offerMan: man.OfferMan,
		set:      man.MetricMan,
	}
}

func (res Resource) Route() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		//r.Use(jwtauth.Verifier(a.authMan.TokenAuth()))
		//r.Use(middleware.NewUserAccessCtx(a.authMan.JWTKey()).ChiMiddleware)
		r.Post("/", res.offers)
	})

	return r
}
