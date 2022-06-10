package market

import (
	"context"
	"errors"
	"fmt"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/models"
)

type auto struct {
	tx             drivers.TxStarter
	marketRepo     repository.Marketer
	autoRepo       repository.Auto
	marketAutoRepo repository.MarketAuto
	sequence       repository.Sequencer
	userApplyRepo  repository.UserApplicationRepository
	userAutoRepo   repository.UserAutoRepository
}

func NewAuto(marketRepo repository.Marketer, marketAutoRepo repository.MarketAuto, sequence repository.Sequencer, tx drivers.TxStarter) *auto {
	return &auto{marketRepo: marketRepo, marketAutoRepo: marketAutoRepo, sequence: sequence, tx: tx}
}

type Auto interface {
	Create(ctx context.Context, auto models.MarketAuto) (string, error)
	Get(ctx context.Context, sku string) (models.MarketAuto, error)
}

func (a auto) Create(ctx context.Context, auto models.MarketAuto) (id string, err error) {
	tx, cb, err := a.tx.StartSession(ctx)
	if err != nil {
		return
	}
	defer func() {
		err = cb(err)
	}()

	userAppl, err := a.userApplyRepo.Lock(tx, auto.AutoSKU)
	if err != nil && !errors.Is(err, drivers.ErrDoesNotExist) {
		return "", fmt.Errorf("user application: %w", err)
	}

	id, err = a.marketAutoRepo.Create(tx, auto)
	if err != nil {
		return
	}

	_, err = a.userAutoRepo.Create(tx, models.UserAuto{
		UserID:        userAppl.UserID,
		ApplicationID: userAppl.ApplicationID,
		VIN:           auto.VIN,
	})

	return
}

func (a auto) Get(ctx context.Context, sku string) (models.MarketAuto, error) {
	return a.marketAutoRepo.Get(ctx, sku)
}
