package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/MultiBanker/broker/src/config"
	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/manager"
	"github.com/MultiBanker/broker/src/servers"
)

type application struct {
	opts    *config.Config
	version string
	testing bool

	servers []servers.Server
	ds      drivers.Datastore
	repo    repository.Repositories
	man     manager.Abstractor
}

func initApp(version string) *application {
	return &application{
		version: version,
		opts:    config.ParseConfig(),
	}
}

func (a *application) run() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	for _, server := range a.servers {
		go func(server servers.Server) {
			if err := server.Start(ctx, cancel); err != nil {
				switch errors.Is(err, http.ErrServerClosed) {
				case true:
					log.Printf("[INFO] Shutting down %s server", server.Name())
				default:
					log.Printf("[ERROR] Server %s not start, or closed", server.Name())
				}
			}
		}(server)
	}

	resetEnv(a.opts.JWTKey, a.opts.DSURL)

	a.shutdown(ctx)
}

func (a *application) shutdown(ctx context.Context) {
	defer a.ds.Close(ctx)

	log.Printf("[INFO] Disable all services")

	wg := sync.WaitGroup{}
	for _, server := range a.servers {
		wg.Add(1)
		go func(server servers.Server) {
			defer wg.Done()
			if err := server.Stop(ctx); err != nil {
				log.Printf("[ERROR] Can't stop %s server", server.Name())
			}
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