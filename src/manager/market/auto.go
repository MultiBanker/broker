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
	marketRepo     repository.Marketer
	autoRepo       repository.Auto
	marketAutoRepo repository.MarketAuto
	sequence       repository.Sequencer
	userApplyRepo  repository.UserApplicationRepository
	userAutoRepo   repository.UserAutoRepository
}

func NewAuto(marketRepo repository.Marketer, marketAutoRepo repository.MarketAuto, sequence repository.Sequencer) *auto {
	return &auto{marketRepo: marketRepo, marketAutoRepo: marketAutoRepo, sequence: sequence}
}

type Auto interface {
	Create(ctx context.Context, auto models.MarketAuto) (string, error)
	Get(ctx context.Context, sku string) (models.MarketAuto, error)
}

func (a auto) Create(ctx context.Context, auto models.MarketAuto) (string, error) {
	// TODO: transaction
	userAppl, err := a.userApplyRepo.Lock(ctx, auto.AutoSKU)
	if err != nil && !errors.Is(err, drivers.ErrDoesNotExist) {
		return "", fmt.Errorf("user application: %w", err)
	}

	id, err := a.marketAutoRepo.Create(ctx, auto)
	if err != nil {
		return "", err
	}

	_, err = a.userAutoRepo.Create(ctx, models.UserAuto{
		ApplicationID: userAppl.ApplicationID,
		VIN:           auto.VIN,
	})

	return id, nil
}

func (a auto) Get(ctx context.Context, sku string) (models.MarketAuto, error) {
	panic("implement me")
}
