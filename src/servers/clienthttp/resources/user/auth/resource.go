package auth

import (
	"github.com/MultiBanker/broker/pkg/auth"
	"github.com/MultiBanker/broker/src/manager/user"
	"github.com/go-chi/chi/v5"
)

type resource struct {
	auther auth.Authenticator
	user   user.UsersManager
}

func NewResource(auther auth.Authenticator, user user.UsersManager) *resource {
	return &resource{auther: auther, user: user}
}

func (res resource) Route() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Post("/auth/signup", res.signUp)
		r.Post("/auth/signin", res.signInByPhone)
	})

	r.Group(func(r chi.Router) {
		r.Delete("/signout", res.signOut)
	})

	return r
}
