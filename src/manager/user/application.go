package user

import (
	"context"
	"errors"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/models"
)

type ApplicationManager interface {
	Create(ctx context.Context, application models.UserApplication) (string, error)
	Get(ctx context.Context, id string) (models.UserApplication, error)
}

type ApplicationManagerImpl struct {
	tx             drivers.TxStarter
	marketAutoRepo repository.MarketAuto
	userApplyRepo  repository.UserApplicationRepository
	userAutoRepo   repository.UserAutoRepository
}

func (a ApplicationManagerImpl) Create(ctx context.Context, application models.UserApplication) (string, error) {
	// TODO: transaction
	tx, cb, err := a.tx.StartSession(ctx)
	if err != nil {
		return "", err
	}
	defer func() {
		err = cb(err)
	}()

	auto, err := a.marketAutoRepo.Lock(tx, application.ChosenSKU)
	if err != nil && !errors.Is(err, drivers.ErrDoesNotExist) {
		return "", err
	}

	id, err := a.userApplyRepo.Create(tx, application)
	if err != nil {
		return "", err
	}

	_, err = a.userAutoRepo.Create(tx, models.UserAuto{
		ApplicationID: id,
		VIN:           auto.VIN,
	})
	if err != nil {
		return "", err
	}

	return "", nil
}

func (a ApplicationManagerImpl) Get(ctx context.Context, id string) (models.UserApplication, error) {
	panic("implement me")
}
