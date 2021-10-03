package metric

import (
	"net/http"
	"strconv"
	"time"

	"github.com/VictoriaMetrics/metrics"
)

type Metricware struct {
	set *metrics.Set
}

func NewMetricware(set *metrics.Set) *Metricware {
	return &Metricware{set: set}
}

const ( //names
	latency = "prometheus_http_request_duration_seconds"
	traffic = "prometheus_http_requests_total"
)

const (
	method  = "Method"
	pathKey = "Path"
	status  = "Status"
)

func (m *Metricware) Latency(path string) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lrw := newLoggingResponseWriter(w)
			start := time.Now()
			handler.ServeHTTP(lrw, r)
			latency := m.set.GetOrCreateHistogram(
				(&NameBuilder{}).Name(latency).
					Add(method, r.Method).
					Add(pathKey, path).
					Add(status, strconv.Itoa(lrw.statusCode)).
					String(),
			)
			latency.UpdateDuration(start)
		})
	}
}

func (m *Metricware) Traffic(path string) func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lrw := newLoggingResponseWriter(w)
			handler.ServeHTTP(lrw, r)
			reqs := m.set.GetOrCreateCounter(
				(&NameBuilder{}).Name(traffic).
					Add(method, r.Method).
					Add(pathKey, path).
					Add(status, strconv.Itoa(lrw.statusCode)).
					String(),
			)
			reqs.Inc()
		})
	}
}

func (m *Metricware) All(path string) []func(handler http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		m.Traffic(path),
		m.Latency(path),
	}
}

