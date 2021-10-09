package partner

import (
	"github.com/MultiBanker/broker/src/manager/auth"
	"github.com/MultiBanker/broker/src/manager/partner"
	"github.com/go-chi/chi/v5"
)

const maxOrderHistoryLimit = 100

type Auth struct {
	authMan    auth.Authenticator
	partnerMan partner.Partnerer
}

func NewAuth(partnerMan partner.Partnerer, authMan auth.Authenticator) Auth {
	return Auth{
		partnerMan: partnerMan,
		authMan:    authMan,
	}
}

func (a Auth) Route() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Post("/login", a.auth())
		r.Get("/logout", a.out())
	})

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
