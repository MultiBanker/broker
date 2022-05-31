package auto

import (
	"github.com/MultiBanker/broker/pkg/auth"
	"github.com/MultiBanker/broker/src/manager/auto"
	"github.com/go-chi/chi/v5"
)

type Resource struct {
	authMan auth.Authenticator
	autoMan auto.Auto
}

func NewResource(authMan auth.Authenticator, autoMan auto.Auto) *Resource {
	return &Resource{authMan: authMan, autoMan: autoMan}
}

func (res Resource) Route() chi.Router {
	r := chi.NewRouter()

	r.Get("/{sku}", res.getCar)
	r.Get("/list/", res.listCar)

	return r
}
