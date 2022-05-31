package application

import (
	"github.com/MultiBanker/broker/pkg/auth"
	"github.com/MultiBanker/broker/src/manager/user"
	"github.com/go-chi/chi/v5"
)

type resource struct {
	auther auth.Authenticator
	user   user.ApplicationManager
}

func NewResource(auther auth.Authenticator, user user.ApplicationManager) *resource {
	return &resource{auther: auther, user: user}
}

func (res resource) Route() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Get("/", res.getApplication)
		r.Post("/", res.createApplication)
		r.Get("/list/", res.listApplications)
	})

	return r
}