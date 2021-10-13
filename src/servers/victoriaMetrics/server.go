package victoriaMetrics

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/VictoriaMetrics/metrics"
)

const (
	readTimeout  = 5 * time.Second
	writeTimeout = 30 * time.Second
)

type victoriaMetricsServer struct {
	certFile string
	keyFile  string
	server   *http.Server
	metric   *metrics.Set
}

func NewVictoriaM(metric *metrics.Set, addr string, r http.Handler) *victoriaMetricsServer {
	return &victoriaMetricsServer{
		metric: metric,
		server: &http.Server{
			Addr:         addr,
			Handler:      r,
			WriteTimeout: writeTimeout,
			ReadTimeout:  readTimeout,
		},
	}
}

func (p victoriaMetricsServer) Name() string {
	return "Metrics"
}

func (p victoriaMetricsServer) Start(_ context.Context, cancel context.CancelFunc) error {
	p.server.RegisterOnShutdown(cancel)

	log.Printf("[INFO] Starting prometheus server on %s", p.server.Addr)

	if p.Insecure() {
		if err := p.server.ListenAndServe(); err != nil {
			return err
		}
	}

	if !p.Insecure() {
		if err := p.server.ListenAndServeTLS(p.certFile, p.keyFile); err != nil {
			return err
		}
	}
	panic("SOMETHING WRONG WITH CERT FILES")
}

func (p victoriaMetricsServer) Stop(ctx context.Context) error {
	<-ctx.Done()
	return p.server.Shutdown(context.Background())
}

func (p *victoriaMetricsServer) Insecure() bool {
	return p.keyFile == "" && p.certFile == ""
}