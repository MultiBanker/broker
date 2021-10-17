package victoriaMetrics

import (
	"context"
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
	return p.server.ListenAndServe()
}

func (p victoriaMetricsServer) Stop(ctx context.Context) error {
	return p.server.Shutdown(ctx)
}

func (p *victoriaMetricsServer) Insecure() bool {
	return p.keyFile == "" && p.certFile == ""
}