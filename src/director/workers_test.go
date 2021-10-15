package director

import (
	"context"
	"testing"
	"time"
)

func TestWorker(t *testing.T) {
	w := NewWorker("test-worker", 1*time.Second)

	someData := "waow"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go w.Run(ctx, func(ctx context.Context) error {
		t.Logf("got: %s", someData)
		return nil
	})

	time.Sleep(5 * time.Second)
}

