package loan

import (
	"github.com/MultiBanker/broker/pkg/auth"
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/manager/loan"
	"github.com/VictoriaMetrics/metrics"
	"github.com/go-chi/chi"
)

type resource struct {
	authMan auth.Authenticator
	loanMan loan.Program
	set     *metrics.Set
}

func Newresource(man manager.Managers) *resource {
	return &resource{
		authMan: man.AuthMan,
		loanMan: man.LoanMan,
		set:     man.MetricMan,
	}
}

func (res resource) Route() chi.Router {
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
