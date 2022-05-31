package verification

import (
	"github.com/MultiBanker/broker/pkg/auth"
	"github.com/MultiBanker/broker/src/manager/user"
	"github.com/go-chi/chi/v5"
)

type resource struct {
	auther auth.Authenticator
	user   user.VerificationManager
}

func NewResource(auther auth.Authenticator, user   user.VerificationManager) *resource {
	return &resource{auther: auther, user: user}
}

func (res resource) Route() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Put("/phone", res.verifyPhone)
		r.Put("/phone/otp", res.verifyPhoneOTP)
	})

	return r
}