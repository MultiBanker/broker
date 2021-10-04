package order

import (
	"context"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/models/dto"
	"github.com/MultiBanker/broker/src/models/selector"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) repository.Orderer {
	return &Repository{collection: collection}
}

func (or Repository) NewOrder(ctx context.Context, order *dto.OrderRequest) (string, error) {
	return "", nil
}

func (or Repository) UpdateOrder(ctx context.Context, order *dto.OrderRequest) (string, error) {
	return "", nil
}

func (or Repository) OrderByID(ctx context.Context, id string) (dto.OrderRequest, error) {
	return dto.OrderRequest{}, nil
}

func (or Repository) Orders(ctx context.Context, paging *selector.Paging) ([]*dto.OrderRequest, int64, error) {
	return nil, 0, nil
}

func (or Repository) OrdersByReferenceID(ctx context.Context, referenceID string) ([]*dto.OrderRequest, error) {
	filter := bson.D{
		{"reference_id", referenceID},
	}

	res, err := or.collection.Find(ctx, filter)
	switch err {
	case mongo.ErrNoDocuments:
		return nil, drivers.ErrDoesNotExist
	case nil:
		orders := make([]*dto.OrderRequest, res.RemainingBatchLength())
		if err := res.All(ctx, orders); err != nil {
			return nil, err
		}
		return orders, nil
	default:
		return nil, err
	}
}
