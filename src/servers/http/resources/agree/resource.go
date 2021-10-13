package agree

import (
	"net/http"

	"github.com/MultiBanker/broker/pkg/httperrors"
	"github.com/MultiBanker/broker/src/manager/agree"
	"github.com/MultiBanker/broker/src/manager/auth"
	"github.com/MultiBanker/broker/src/manager/signature"
	"github.com/MultiBanker/broker/src/servers/http/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
)

type Resource struct {
	auth         auth.Authenticator
	specMan      agree.Specification
	signatureMan signature.Signature
}

func NewResource(auth auth.Authenticator, specMan agree.Specification, signatureMan signature.Signature) *Resource {
	return &Resource{
		auth:         auth,
		specMan:      specMan,
		signatureMan: signatureMan,
	}

}

func (res Resource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Route("/signatures", func(r chi.Router) {
			r.Post("/", res.create)
			r.Get("/", res.list)

			r.Route("/{id}", func(r chi.Router) {
				//r.Use(guard.Allow(errorRender, "moderator"))
				r.Get("/", res.get)
				r.Patch("/resend", res.resend)
				r.Put("/validation", res.token)
			})

		})
	})

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(res.auth.TokenAuth()))
		r.Use(middleware.NewUserAccessCtx(res.auth.JWTKey()).ChiMiddleware)
		// TODO: Add Lib For guarding
		//r.Use(guard.Allow(errorRender, "moderator"))

		r.Get("/code/{code}", res.getByCode)
		r.Get("/{id}", res.getAgree)
		r.Get("/", res.getAgrees)
		r.Post("/", res.createAgree)
		r.Put("/{id}", res.update)


	})


	return r
}

func errorRender(w http.ResponseWriter, r *http.Request, err error) {
	render.Status(r, http.StatusForbidden)
	_ = render.Render(w, r, httperrors.AccessDenied(err))
}
