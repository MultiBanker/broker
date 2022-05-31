package auth

import (
	"github.com/MultiBanker/broker/pkg/auth"
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/manager/market"
	"github.com/VictoriaMetrics/metrics"
	"github.com/go-chi/chi/v5"
)

type Resource struct {
	authMan   auth.Authenticator
	makretMan market.Marketer
	set       *metrics.Set
}

func NewResource(man manager.Managers) *Resource {
	return &Resource{
		authMan:   man.AuthMan,
		makretMan: man.MarketMan,
		set:       man.MetricMan,
	}
}

func (res Resource) Route() chi.Router {
	r := chi.NewRouter()

	r.Post("/login", res.auth())
	r.Get("/logout", res.out())

	return r
}
