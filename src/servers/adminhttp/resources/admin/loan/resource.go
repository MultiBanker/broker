package loan

import (
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/manager/auth"
	"github.com/MultiBanker/broker/src/manager/loan"
	"github.com/VictoriaMetrics/metrics"
	"github.com/go-chi/chi/v5"
)

type AdminResource struct {
	authMan auth.Authenticator
	loanMan loan.Program
	set     *metrics.Set
}

func NewAdminResource(man manager.Wrapper) *AdminResource {
	return &AdminResource{
		authMan: man.Auther(),
		loanMan: man.LoanProgram(),
		set:     man.Metric(),
	}
}

func (res AdminResource) Route() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		//r.Use(jwtauth.Verifier(a.authMan.TokenAuth()))
		//r.Use(middleware.NewUserAccessCtx(a.authMan.JWTKey()).ChiMiddleware)
		r.Post("/", res.create)
		r.Get("/", res.list)
		r.Get("/{code}", res.get)
		r.Put("/{code}", res.update)
	})

	return r
}
