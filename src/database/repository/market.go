package repository

import (
	"context"

	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
)

type Marketer interface {
	CreateMarket(ctx context.Context, market models.Market) (string, error)
	MarketByID(ctx context.Context, id string) (models.Market, error)
	MarketByUsername(ctx context.Context, username string) (models.Market, error)
	Markets(ctx context.Context, paging selector.Paging) ([]models.Market, int64, error)
	UpdateMarket(ctx context.Context, market models.Market) (string, error)
	MarketByCode(ctx context.Context, code string) (models.Market, error)
}

type MarketAuto interface {
	Create(ctx context.Context, auto models.MarketAuto) (string, error)
	Get(ctx context.Context, sku string) (models.MarketAuto, error)
	Lock(ctx context.Context, id string) (models.MarketAuto, error)
}
