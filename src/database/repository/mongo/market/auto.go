package market

import (
	"context"

	"github.com/MultiBanker/broker/src/database/repository/mongo/transaction"
	"github.com/MultiBanker/broker/src/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type AutoRepository struct {
	coll        *mongo.Collection
	transaction transaction.Func
}

func NewAutoRepository(coll *mongo.Collection, transaction transaction.Func) *AutoRepository {
	return &AutoRepository{coll: coll, transaction: transaction}
}

func (a AutoRepository) Create(ctx context.Context, auto models.MarketAuto) (string, error) {
	panic("implement me")
}

func (a AutoRepository) Get(ctx context.Context, sku string) (models.MarketAuto, error) {
	panic("implement me")
}

func (a AutoRepository) Lock(ctx context.Context, id string) (models.MarketAuto, error) {
	panic("implement me")
}
