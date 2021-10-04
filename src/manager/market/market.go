package market

import (
	"context"
	"strconv"

	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
)

type Market struct {
	marketColl repository.Marketer
	seqColl    repository.Sequencer
}

func NewMarket(repo repository.Repositories) Marketer {
	return Market{
		marketColl: repo.MarketRepo(),
		seqColl:    repo.SequenceRepo(),
	}
}

type Marketer interface {
	CreateMarket(ctx context.Context, market models.Market) (string, error)
	MarketByID(ctx context.Context, id string) (models.Market, error)
	Markets(ctx context.Context, paging selector.Paging) ([]models.Market, int64, error)
	UpdateMarket(ctx context.Context, market models.Market) (string, error)
}

func (m Market) CreateMarket(ctx context.Context, market models.Market) (string, error) {
	idInt, err := m.seqColl.NextSequenceValue(ctx, models.MarketSequences)
	if err != nil {
		return "", err
	}
	market.ID = strconv.Itoa(idInt)
	return m.marketColl.CreateMarket(ctx, market)
}

func (m Market) MarketByID(ctx context.Context, id string) (models.Market, error) {
	return m.marketColl.MarketByID(ctx, id)
}

func (m Market) Markets(ctx context.Context, paging selector.Paging) ([]models.Market, int64, error) {
	return m.marketColl.Markets(ctx, paging)
}

func (m Market) UpdateMarket(ctx context.Context, market models.Market) (string, error) {
	return m.marketColl.UpdateMarket(ctx, market)
}
