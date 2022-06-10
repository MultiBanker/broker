package auto

import (
	"github.com/MultiBanker/broker/pkg/auth"
	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/go-chi/chi/v5"
)

type resource struct {
	auther       auth.Authenticator
	userAutoRepo repository.UserAutoRepository
}

func NewResource(auther auth.Authenticator, userAutoRepo repository.UserAutoRepository) *resource {
	return &resource{auther: auther, userAutoRepo: userAutoRepo}
}

func (res resource) Route() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Get("/auto", res.getAuto)
	})

	return r
}
