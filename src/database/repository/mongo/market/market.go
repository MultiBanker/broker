package market

import (
	"context"

	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) repository.Marketer {
	return Repository{
		collection: collection,
	}
}

func (r Repository) CreateMarket(ctx context.Context, market models.Market) (string, error) {
	panic("implement me")
}

func (r Repository) MarketByID(ctx context.Context, id string) (models.Market, error) {
	panic("implement me")
}

func (r Repository) Markets(ctx context.Context, paging selector.Paging) ([]models.Market, int64, error) {
	panic("implement me")
}

func (r Repository) UpdateMarket(ctx context.Context, market models.Market) (string, error) {
	panic("implement me")
}