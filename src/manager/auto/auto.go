package auto

import (
	"context"

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
	sequence repository.Sequencer
	autoRepo repository.Auto
}

func NewAutoImpl(autoRepo repository.Auto) *autoImpl {
	return &autoImpl{autoRepo: autoRepo}
}

func (a autoImpl) Get(ctx context.Context, sku string) (models.Auto, error) {
	panic("implement me")
}

func (a autoImpl) Create(ctx context.Context, auto models.Auto) (string, error) {
	panic("implement me")
}

func (a autoImpl) Update(ctx context.Context, auto models.Auto) (string, error) {
	panic("implement me")
}

func (a autoImpl) List(ctx context.Context, query selector.SearchQuery) ([]models.Auto, int64, error) {
	panic("implement me")
}

func (a autoImpl) Delete(ctx context.Context, sku string) error {
	panic("implement me")
}
