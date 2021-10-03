package http

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/MultiBanker/broker/src/servers/http/resources/health"
	"github.com/go-chi/chi"

	"github.com/MultiBanker/broker/src/config"
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/servers/http/middleware"
	"github.com/MultiBanker/broker/src/servers/http/resources/authresource"
	"github.com/MultiBanker/broker/src/servers/http/resources/orderresource"
)

const (
	readTimeout  = 5 * time.Second
	writeTimeout = 30 * time.Second
	ApiPath      = "/api/v1"
)

func Routing(opts *config.Config, man manager.Abstractor) chi.Router {
	isReady := &atomic.Value{}
	go readyzProbe(isReady)

	r := middleware.Mount(opts.Version, opts.HTTP.FilesDir, opts.HTTP.BasePath)

	// основные роутеры
	r.Route(ApiPath, func(r chi.Router) {
		r.Route("/brokers", func(r chi.Router) {
			r.Mount("/auth", authresource.NewAuth(man.Partnerer(), man.Auther()).Route())
			r.Mount("/orders", orderresource.NewOrder(man.Orderer()).Route())
		})
	})

	r.Mount("/kubernetes", health.NewKuber(isReady, func() error {
		return man.Pinger()
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
