package market

import (
	"context"
	"fmt"
	"strconv"

	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/manager/auth"
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
	MarketByUsername(ctx context.Context, username, password string) (models.Market, error)
}

var _ Marketer = (*Market)(nil)

func (m Market) CreateMarket(ctx context.Context, market models.Market) (string, error) {
	bytePass, err := auth.HashPassword(market.Password)
	if err != nil {
		return "", err
	}
	market.Password = string(bytePass)
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
	res, err := m.marketColl.MarketByID(ctx, market.ID)
	if err != nil {
		return "", err
	}
	if !auth.CheckPasswordHash(market.Password, []byte(res.Password)) {
		return "", fmt.Errorf("[ERROR] Wrong Password")
	}
	return m.marketColl.UpdateMarket(ctx, market)
}

func (m Market) MarketByUsername(ctx context.Context, username, password string) (models.Market, error) {
	var market models.Market
	market, err := m.marketColl.MarketByUsername(ctx, username)
	if err != nil {
		return market, err
	}
	if !auth.CheckPasswordHash(password, []byte(market.Password)) {
		return market, fmt.Errorf("[ERROR] Wrong Password")
	}

	return market, nil
}
