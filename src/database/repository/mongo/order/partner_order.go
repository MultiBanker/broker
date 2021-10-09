package order

import (
	"context"
	"errors"
	"time"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/models/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PartnerOrderRepository struct {
	coll *mongo.Collection
}

func NewPartnerOrderRepository(coll *mongo.Collection) *PartnerOrderRepository {
	return &PartnerOrderRepository{coll: coll}
}

func (p PartnerOrderRepository) NewOrder(ctx context.Context, request dto.OrderResponse) (string, error) {
	request.ID = primitive.NewObjectID().Hex()
	_, err := p.coll.InsertOne(ctx, request)
	return request.ID, err
}

func (p PartnerOrderRepository) UpdateOrder(ctx context.Context, response dto.OrderPartnerUpdateRequest) (string, error) {
	var order dto.OrderResponse

	filter := bson.D{
		{"partner_code", response.PartnerCode},
		{"reference_id", response.ReferenceID},
	}

	update := bson.D{
		{"state", response.State},
		{"state_title", response.StateTitle},
		{"customer", response.Customer},
		{"offers", response.Offers},
		{"updated_at", time.Now().UTC()},
	}

	err := p.coll.FindOneAndUpdate(ctx, filter, update).Decode(&order)
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return order.ID, drivers.ErrDoesNotExist
	case errors.Is(err, nil):
		return order.ID, nil
	}
	return order.ID, nil
}

func (p PartnerOrderRepository) OrdersByReferenceID(ctx context.Context, marketCode, referenceID string) ([]*dto.OrderResponse, error) {
	filter := bson.D{
		{"market_code", marketCode},
		{"reference_id", referenceID},
	}

	cur, err := p.coll.Find(ctx, filter)
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return nil, drivers.ErrDoesNotExist
	case errors.Is(err, nil):
		orders := make([]*dto.OrderResponse, cur.RemainingBatchLength())
		if err := cur.All(ctx, &orders); err != nil {
			return nil, err
		}
		return orders, nil
	}
	return nil, err
}
