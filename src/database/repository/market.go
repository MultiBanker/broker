package repository

import (
	"context"

	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
)

func (r Repository) MarketRepo() Marketer {
	return r.Market
}

type Marketer interface {
	CreateMarket(ctx context.Context, market models.Market) (string, error)
	MarketByID(ctx context.Context, id string) (models.Market, error)
	MarketByUsername(ctx context.Context, username string) (models.Market, error)
	Markets(ctx context.Context, paging selector.Paging) ([]models.Market, int64, error)
	UpdateMarket(ctx context.Context, market models.Market) (string, error)
	MarketByCode(ctx context.Context, code string) (models.Market, error)
}
