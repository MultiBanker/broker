package victoriaMetrics

import (
	"net/http"

	"github.com/VictoriaMetrics/metrics"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func Routing() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		metrics.WritePrometheus(w, true)
	})
	return r
}
