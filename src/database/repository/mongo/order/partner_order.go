package order

import (
	"context"
	"errors"
	"time"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/models"
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

func (p PartnerOrderRepository) NewOrder(ctx context.Context, order models.PartnerOrder) (string, error) {
	order.ID = primitive.NewObjectID().Hex()
	order.CreatedAt = time.Now().UTC()
	_, err := p.coll.InsertOne(ctx, order)
	return order.ID, err
}

func (p PartnerOrderRepository) UpdateOrder(ctx context.Context, porder models.PartnerOrder) (string, error) {
	var order models.PartnerOrder

	filter := bson.D{
		{"partner_code", porder.PartnerCode},
		{"reference_id", porder.ReferenceID},
	}

	update := bson.D{
		{"state", porder.State},
		{"state_title", porder.StateTitle},
		{"customer", porder.Customer},
		{"offers", porder.Offers},
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

func (p PartnerOrderRepository) UpdateInitStatusByTimeOut(ctx context.Context) error {
	now := time.Now().UTC()

	filter := bson.D{
		{"state", models.INIT.Status()},
		{"created_at", bson.D{
			{Key: "$lte", Value: now.Add(-3 * time.Minute)},
		}},
	}
	update := bson.D{
		{"$set", bson.D{
			{"state", models.CANCELLED.Status()},
			{"statetitle", models.CANCELLED.Title()},
		}},
	}

	err := p.coll.FindOneAndUpdate(ctx, filter, update).Err()
	switch {
	case errors.Is(err, nil), errors.Is(err, mongo.ErrNoDocuments):
		return nil
	default:
		return err
	}
}

func (p PartnerOrderRepository) OrdersByReferenceID(ctx context.Context, marketCode, referenceID string) ([]*models.PartnerOrder, error) {
	filter := bson.D{
		{"market_code", marketCode},
		{"reference_id", referenceID},
	}

	cur, err := p.coll.Find(ctx, filter)
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return nil, drivers.ErrDoesNotExist
	case errors.Is(err, nil):
		orders := make([]*models.PartnerOrder, cur.RemainingBatchLength())
		if err := cur.All(ctx, &orders); err != nil {
			return nil, err
		}
		return orders, nil
	}
	return nil, err
}

func (p PartnerOrderRepository) OrderPartner(ctx context.Context, referenceID, partnerCode string) (models.PartnerOrder, error) {
	var order models.PartnerOrder

	filter := bson.D{
		{"reference_id", referenceID},
		{"partner_code", partnerCode},
	}

	err := p.coll.FindOne(ctx, filter).Decode(&order)
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return order, drivers.ErrDoesNotExist
	case errors.Is(err, nil):
		return order, nil
	}
	return order, err
}
