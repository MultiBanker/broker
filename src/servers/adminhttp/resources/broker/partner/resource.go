package partner

import (
	"github.com/MultiBanker/broker/pkg/auth"
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/manager/partner"
	"github.com/VictoriaMetrics/metrics"
	"github.com/go-chi/chi/v5"
)

const maxOrderHistoryLimit = 100

type resource struct {
	authMan    auth.Authenticator
	partnerMan partner.Partnerer
	set        *metrics.Set
}

func NewResource(man manager.Managers) resource {
	return resource{
		authMan:    man.AuthMan,
		set:        man.MetricMan,
		partnerMan: man.PartnerMan,
	}
}

func (a resource) Route() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		//r.Use(jwtauth.Verifier(a.authMan.TokenAuth()))
		//r.Use(middleware.NewUserAccessCtx(a.authMan.JWTKey()).ChiMiddleware)
		r.Post("/", a.newpartner())
		r.Get("/{id}", a.partner())
		r.Put("/{id}", a.update())
		r.Get("/", a.partners())
	})
	return r
}
