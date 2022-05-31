package healthz

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/servers/healthz/health"
	"github.com/go-chi/chi/v5"
)

func Routing(man manager.Managers) chi.Router {
	isReady := &atomic.Value{}
	go readyzProbe(isReady)

	r := chi.NewRouter()

	r.Mount("/healthcheck", health.New(isReady, func() error {
		return man.DB.Ping()
	}).Route())

	return r
}

func readyzProbe(isReady *atomic.Value) {
	isReady.Store(false)
	log.Printf("Readyz probe is negative by default...")
	time.Sleep(10 * time.Second)
	isReady.Store(true)
	log.Printf("Readyz probe is positive.")
}
