package auth

import (
	"github.com/MultiBanker/broker/pkg/auth"
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/manager/partner"
	"github.com/VictoriaMetrics/metrics"
	"github.com/go-chi/chi/v5"
)

type Resource struct {
	authMan    auth.Authenticator
	partnerMan partner.Partnerer
	set        *metrics.Set
}

func NewResource(man manager.Managers) *Resource {
	return &Resource{
		authMan:    man.AuthMan,
		partnerMan: man.PartnerMan,
		set:        man.MetricMan,
	}
}

func (res Resource) Route() chi.Router {
	r := chi.NewRouter()

	r.Post("/login", res.auth())
	r.Get("/logout", res.out())

	return r
}
