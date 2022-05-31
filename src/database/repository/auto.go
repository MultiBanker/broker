package repository

import (
	"context"

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
