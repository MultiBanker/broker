package health

import (
	"net/http"
	"sync/atomic"

	"github.com/go-chi/chi"
)

type Kubernetes struct {
	isReady *atomic.Value
	pinger  func() error
}

func NewKuber(isReady *atomic.Value, pinger func() error) *Kubernetes {
	return &Kubernetes{
		isReady: isReady,
		pinger:  pinger,
	}
}

func (k Kubernetes) Route() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Get("/healthz", k.healthz())
		r.Get("/readyz", k.readyz())
	})

	return r
}

func (k Kubernetes) healthz() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := k.pinger(); err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
		_, _ = w.Write([]byte("System is running"))
		w.WriteHeader(http.StatusOK)
	}
}

func (k Kubernetes) readyz() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isready, ok := k.isReady.Load().(bool)
		if !ok {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if k.isReady == nil || !isready {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
