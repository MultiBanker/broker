package http

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/MultiBanker/broker/pkg/metric"
	"github.com/MultiBanker/broker/src/servers/http/resources/health"
	"github.com/MultiBanker/broker/src/servers/http/resources/market"
	"github.com/MultiBanker/broker/src/servers/http/resources/offer"
	"github.com/VictoriaMetrics/metrics"

	"github.com/go-chi/chi/v5"

	"github.com/MultiBanker/broker/src/config"
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/servers/http/middleware"
	"github.com/MultiBanker/broker/src/servers/http/resources/orderresource"
	"github.com/MultiBanker/broker/src/servers/http/resources/partner"
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
	mware := metric.NewMetricware(metrics.NewSet())

	// основные роутеры
	r.Route(ApiPath, func(r chi.Router) {
		r.Use(mware.All("/broker")...)
		r.Route("/broker", func(r chi.Router) {
			r.Mount("/partners", partner.NewAuth(man.Auther(), man.Partnerer(), man.Metric()).Route())
			r.Mount("/orders", orderresource.NewOrder(man.Auther(), man.Orderer(), man.Metric()).Route())
			r.Mount("/markets", market.NewResource(man.Auther(), man.Marketer(), man.Metric()).Route())
			r.Mount("/offers", offer.NewResource(man.Auther(), man.Offer(), man.Metric()).Route())
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
