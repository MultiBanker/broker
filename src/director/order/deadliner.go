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

type Order struct {
	partnerOrderRepo repository.PartnerOrderer
}

func NewDeadline(partnerOrderRepo repository.PartnerOrderer) director.Daemons {
	return Order{
		partnerOrderRepo: partnerOrderRepo,
	}
}

func (d Order) Name() string {
	return "order-worker"
}

func (d Order) Start(ctx context.Context, _ context.CancelFunc) error {
	orderTimeKiller := director.NewWorker(d.Name(), defaultTicker)
	go orderTimeKiller.Run(ctx, d.InitTimeOutKill)
	return nil
}

func (d Order) InitTimeOutKill(ctx context.Context) error {
	return d.partnerOrderRepo.UpdateInitStatusByTimeOut(ctx)
}

func (d Order) Stop(_ context.Context) error {
	return nil
}
