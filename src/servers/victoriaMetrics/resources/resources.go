package resources

import (
	"net/http"

	"github.com/VictoriaMetrics/metrics"
	"github.com/go-chi/chi/v5"
)

type Metrics struct {
}

func NewMetrics() Metrics {
	return Metrics{}
}

func (res Metrics) Route() chi.Router {
	r := chi.NewRouter()

	r.HandleFunc("/", res.Metric)

	return r
}

func (res Metrics) Metric(w http.ResponseWriter, r *http.Request) {
	metrics.WritePrometheus(w, true)
}
