package repository

import (
	"context"

	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
)

type UserAutoRepository interface {
	Get(ctx context.Context, sku string) (models.UserAuto, error)
	List(ctx context.Context, search selector.SearchQuery) ([]models.UserAuto, int64, error)
	Create(ctx context.Context, auto models.UserAuto) (string, error)
}
