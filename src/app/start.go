package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/MultiBanker/broker/src/config"
	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/servers"
	"github.com/VictoriaMetrics/metrics"
)

const killTimeOut = 35 * time.Second

type application struct {
	opts    *config.Config
	version string
	testing bool

	servers []servers.Server
	ds      drivers.Datastore
	repo    repository.Repositories
	man     manager.Abstractor
	metric  *metrics.Set
}

func initApp(version string) (*application, error) {
	opts, err := config.ParseConfig()
	if err != nil {
		return nil, err
	}
	resetEnv(opts.JWTKey, opts.DSURL)
	return &application{
		version: version,
		opts:    opts,
		metric:  metrics.NewSet(),
	}, nil
}

func (a *application) run() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	for _, server := range a.servers {
		log.Printf("[INFO] Starting %s server", server.Name())
		go func(server servers.Server) {
			server.Start(ctx, cancel)
		}(server)
	}

	a.shutdown(ctx)
}

func (a *application) shutdown(ctx context.Context) {
	<-ctx.Done()

	killContext, cancel := context.WithTimeout(context.Background(), killTimeOut)
	defer cancel()
	defer a.ds.Close(killContext)

	log.Printf("[INFO] Disable all services")

	wg := sync.WaitGroup{}
	for _, server := range a.servers {
		wg.Add(1)
		go func(server servers.Server) {
			defer wg.Done()
			if err := server.Stop(killContext); err != nil {
				log.Printf("[ERROR] Can't stop %s server", server.Name())
			}

			log.Printf("[INFO] Shutting down %s server", server.Name())

		}(server)
	}
	wg.Wait()

	log.Println("[INFO] Terminated")
}

// resetEnv() - сбрасывает чувствительные переменные окружения.
func resetEnv(envs ...string) {
	for _, env := range envs {
		if err := os.Unsetenv(env); err != nil {
			log.Printf("[WARN] can't unset env %s, %s", env, err)
		}
	}
}
