package victoriaMetrics

import (
	"github.com/MultiBanker/broker/src/servers/victoriaMetrics/resources"
	"github.com/go-chi/chi"
)

func Routing() chi.Router {
	r := chi.NewRouter()
	r.Mount("/metrics", resources.NewMetrics().Route())
	return r
}
