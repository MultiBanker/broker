package order

import (
	"context"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/models/dto"
	"github.com/MultiBanker/broker/src/models/selector"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	collection *mongo.Collection
}

func NewOrderRepository(collection *mongo.Collection) *OrderRepository {
	return &OrderRepository{collection: collection}
}

func (or OrderRepository) NewOrder(ctx context.Context, order *dto.OrderRequest) (string, error) {
	return "", nil
}

func (or OrderRepository) UpdateOrder(ctx context.Context, order *dto.OrderRequest) (string, error) {
	return "", nil
}

func (or OrderRepository) OrderByID(ctx context.Context, id string) (dto.OrderRequest, error) {
	return dto.OrderRequest{}, nil
}

func (or OrderRepository) Orders(ctx context.Context, paging *selector.Paging) ([]*dto.OrderRequest, error) {
	return nil, nil
}

func (or OrderRepository) OrdersByReferenceID(ctx context.Context, referenceID string) ([]*dto.OrderRequest, error) {
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
