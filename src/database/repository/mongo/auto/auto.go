package auto

import (
	"context"

	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	coll *mongo.Collection
}

func NewRepository(coll *mongo.Collection) *Repository {
	return &Repository{coll: coll}
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