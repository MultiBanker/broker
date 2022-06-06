package userauto

import (
	"context"

	"github.com/MultiBanker/broker/src/database/repository/mongo/transaction"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
	"go.mongodb.org/mongo-driver/mongo"
)

type userAutoRepository struct {
	coll        *mongo.Collection
	transaction transaction.Func
}

func NewUserAutoRepository(coll *mongo.Collection, transaction transaction.Func) *userAutoRepository {
	return &userAutoRepository{coll: coll, transaction: transaction}
}

func (u userAutoRepository) Get(ctx context.Context, sku string) (models.UserAuto, error) {
	panic("implement me")
}

func (u userAutoRepository) List(ctx context.Context, search selector.SearchQuery) ([]models.UserAuto, int64, error) {
	panic("implement me")
}

func (u userAutoRepository) Create(ctx context.Context, auto models.UserAuto) (string, error) {
	panic("implement me")
}
