package offer

import (
	"github.com/MultiBanker/broker/src/manager/auth"
	"github.com/MultiBanker/broker/src/manager/offer"
	"github.com/go-chi/chi/v5"
)

type Resource struct {
	authMan  auth.Authenticator
	offerMan offer.Manager
}

func NewResource(authMan auth.Authenticator, offerMan offer.Manager) *Resource {
	return &Resource{authMan: authMan, offerMan: offerMan}
}

func (res Resource) Route() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		//r.Use(jwtauth.Verifier(a.authMan.TokenAuth()))
		//r.Use(middleware.NewUserAccessCtx(a.authMan.JWTKey()).ChiMiddleware)
		r.Post("/", res.create)
		r.Get("/", res.list)
		r.Get("/{code}", res.get)
		r.Put("/{id}", res.update)
	})

	return r
}
