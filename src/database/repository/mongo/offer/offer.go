package offer

import (
	"context"
	"errors"
	"time"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	coll *mongo.Collection
}

func NewRepository(coll *mongo.Collection) *Repository {
	return &Repository{coll: coll}
}

func (r Repository) CreateOffer(ctx context.Context, offer models.Offer) (string, error) {
	offer.CreatedAt = time.Now().UTC()
	_, err := r.coll.InsertOne(ctx, offer)
	return offer.ID, err
}

func (r Repository) UpdateOffer(ctx context.Context, offer models.Offer) (models.Offer, error) {
	var newoffer models.Offer

	filter := bson.D{
		{"_id", offer.ID},
	}

	update := bson.D{
		{"$set", bson.D{
			{"name", offer.Name},
			{"partner_code", offer.PartnerCode},
			{"payment_type_group_code", offer.PaymentTypeGroupCode},
			{"min_order_sum", offer.MinOrderSum},
			{"max_order_sum", offer.MaxOrderSum},
			{"updated_at", time.Now().UTC()}},
		},
	}

	returnDoc := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &returnDoc,
	}

	err := r.coll.FindOneAndUpdate(ctx, filter, update, &opts).Decode(&newoffer)
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return newoffer, drivers.ErrDoesNotExist
	case errors.Is(err, nil):
		return newoffer, nil
	}
	return newoffer, err
}

func (r Repository) OfferByCode(ctx context.Context, code string) (models.Offer, error) {
	var offer models.Offer

	filter := bson.D{
		{"partner_code", code},
	}
	err := r.coll.FindOne(ctx, filter).Decode(&offer)
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return offer, nil
	case errors.Is(err, nil):
		return offer, nil
	}
	return offer, err
}

func (r Repository) Offers(ctx context.Context, paging selector.Paging) ([]models.Offer, int64, error) {
	filter := bson.D{}

	opts := options.FindOptions{
		Skip: &paging.Skip,
		Sort: bson.D{
			{Key: paging.SortKey, Value: paging.SortVal},
		},
		Limit: &paging.Limit,
	}

	total, err := r.coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	res, err := r.coll.Find(ctx, filter, &opts)
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return nil, 0, drivers.ErrDoesNotExist
	case errors.Is(err, nil):
		offers := make([]models.Offer, res.RemainingBatchLength())
		err = res.All(ctx, &offers)
		if err != nil {
			return nil, 0, err
		}

		return offers, total, nil
	}
	return nil, 0, err
}

func (r Repository) OffersByTotalSum(ctx context.Context, total int) ([]*models.Offer, error) {
	filter := bson.D{
		{"min_order_sum", bson.D{
			{"$lte", total},
		}},
		{"max_order_sum", bson.D{
			{"$gte",total},
		}},
	}

	cur, err := r.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	offers := make([]*models.Offer, cur.RemainingBatchLength())
	err = cur.All(ctx, &offers)

	return offers, err
}
