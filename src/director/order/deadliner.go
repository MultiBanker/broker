package order

import (
	"context"
	"time"

	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/director"
)

const (
	defaultTicker = 300 * time.Millisecond
)

type Deadline struct {
	partnerOrderRepo repository.PartnerOrderer
}

func NewDeadline(partnerOrderRepo repository.PartnerOrderer) director.Daemons {
	return Deadline{
		partnerOrderRepo: partnerOrderRepo,
	}
}

func (d Deadline) Name() string {
	return "order-deadliner"
}

func (d Deadline) Start(ctx context.Context, _ context.CancelFunc) error {
	orderTimeKiller := director.NewWorker(d.Name(), defaultTicker)
	go orderTimeKiller.Run(ctx, d.InitTimeOutKill)
	return nil
}

func (d Deadline) InitTimeOutKill(ctx context.Context) error {
	return d.partnerOrderRepo.UpdateInitStatusByTimeOut(ctx)
}

func (d Deadline) Stop(_ context.Context) error {
	return nil
}
