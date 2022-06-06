package auto

import (
	"context"

	"github.com/MultiBanker/broker/src/database/repository/mongo/transaction"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	coll        *mongo.Collection
	transaction transaction.Func
}

func NewRepository(coll *mongo.Collection, transaction transaction.Func) *Repository {
	return &Repository{coll: coll, transaction: transaction}
}

func (r Repository) Get(ctx context.Context, sku string) (models.Auto, error) {
	panic("implement me")
}

func (r Repository) Create(ctx context.Context, auto models.Auto) (string, error) {
	panic("implement me")
}

func (r Repository) Update(ctx context.Context, auto models.Auto) (string, error) {
	panic("implement me")
}

func (r Repository) List(ctx context.Context, query selector.SearchQuery) ([]models.Auto, int64, error) {
	panic("implement me")
}

func (r Repository) Delete(ctx context.Context, sku string) error {
	panic("implement me")
}