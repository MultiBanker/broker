package director

import (
	"context"
	"log"
	"time"
)

type TaskFn func(ctx context.Context) error

type Worker struct {
	name  string
	every *time.Ticker
}

func NewWorker(name string, every time.Duration) *Worker {
	return &Worker{
		name:  name,
		every: time.NewTicker(every),
	}
}

func (w Worker) Run(ctx context.Context, fn TaskFn) {
	log.Printf("[INFO] run task: %s\n", w.name)

	defer w.every.Stop()
	for {
		select {
		case <-ctx.Done():
			log.Printf("[INFO] stop task: %s\n", w.name)
			return
		case <-w.every.C:
			if err := fn(ctx); err != nil {
				log.Printf("[WARN] run task %s error: %s\n", w.name, err)
			}
		}
	}
}
