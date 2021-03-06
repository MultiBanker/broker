package auto

import (
	"context"
	"strconv"

	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
)

type Auto interface {
	Get(ctx context.Context, sku string) (models.Auto, error)
	Create(ctx context.Context, auto models.Auto) (string, error)
	Update(ctx context.Context, auto models.Auto) (string, error)
	List(ctx context.Context, query selector.SearchQuery) ([]models.Auto, int64, error)
	Delete(ctx context.Context, sku string) error
}

type autoImpl struct {
	sequence      repository.Sequencer
	autoRepo      repository.Auto
	userApplyRepo repository.UserApplicationRepository
	userAutoRepo  repository.UserAutoRepository
}

func NewAutoImpl(autoRepo repository.Auto) *autoImpl {
	return &autoImpl{autoRepo: autoRepo}
}

func (a autoImpl) Get(ctx context.Context, sku string) (models.Auto, error) {
	return a.autoRepo.Get(ctx, sku)
}

func (a autoImpl) Create(ctx context.Context, auto models.Auto) (string, error) {
	skuInt, err := a.sequence.NextSequenceValue(ctx, models.AutoSequences)
	if err != nil {
		return "", err
	}

	auto.SKU = strconv.Itoa(skuInt)
	_, err = a.autoRepo.Create(ctx, auto)
	if err != nil {
		return "", err
	}

	return auto.SKU, nil
}

func (a autoImpl) Update(ctx context.Context, auto models.Auto) (string, error) {
	return a.autoRepo.Update(ctx, auto)
}

func (a autoImpl) List(ctx context.Context, query selector.SearchQuery) ([]models.Auto, int64, error) {
	return a.autoRepo.List(ctx, query)
}

func (a autoImpl) Delete(ctx context.Context, sku string) error {
	return a.autoRepo.Delete(ctx, sku)
}
