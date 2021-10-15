package order

import (
	"context"
	"log"
	"time"

	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/director"
)

const (
	defaultTicker = 300 * time.Millisecond
)

type Deadline struct {
	partnerOrderRepo repository.PartnerOrderer
	timer            *time.Ticker
}

func NewDeadline(partnerOrderRepo repository.PartnerOrderer) director.Daemons {
	return Deadline{
		partnerOrderRepo: partnerOrderRepo,
		timer:            time.NewTicker(defaultTicker),
	}
}

func (d Deadline) Name() string {
	return "order-deadliner"
}

func (d Deadline) Start(ctx context.Context, cancel context.CancelFunc) error {
	for {
		select {
		case <-d.timer.C:
			if err := d.partnerOrderRepo.UpdateInitStatusByTimeOut(ctx); err != nil {
				log.Println("[ERROR] Database ", err)
			}
		case <-ctx.Done():
			d.timer.Stop()
			return nil
		}
	}
}

func (d Deadline) Stop(_ context.Context) error {
	return nil
}
