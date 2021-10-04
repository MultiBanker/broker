package market

import (
	"context"
	"errors"
	"time"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	market.CreatedAt = time.Now().UTC()
	_, err := r.collection.InsertOne(ctx, market)
	return market.ID, err
}

func (r Repository) MarketByID(ctx context.Context, id string) (models.Market, error) {
	var market models.Market

	filter := bson.D{
		{"_id", id},
	}
	err := r.collection.FindOne(ctx, filter).Decode(&market)
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return market, drivers.ErrDoesNotExist
	case errors.Is(err, nil):
		return market, nil
	}
	return market, err
}

func (r Repository) Markets(ctx context.Context, paging selector.Paging) ([]models.Market, int64, error) {
	filter := bson.D{}

	opts := options.FindOptions{
		Skip: &paging.Skip,
		Sort: bson.D{
			{Key: paging.SortKey, Value: paging.SortVal},
		},
		Limit: &paging.Limit,
	}

	total, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	res, err := r.collection.Find(ctx, filter, &opts)
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return nil, 0, drivers.ErrDoesNotExist
	case errors.Is(err, nil):
		markets := make([]models.Market, res.RemainingBatchLength())
		err = res.All(ctx, &markets)
		if err != nil {
			return nil, 0, err
		}

		return markets, total, nil
	}
	return nil, 0, err
}

func (r Repository) UpdateMarket(ctx context.Context, market models.Market) (string, error) {
	filter := bson.D{
		{"_id", market.ID},
	}

	update := bson.D{
		{"title", market.Title},
		{"logo_url", market.LogoURL},
		{"location", market.Location},
		{"web_address", market.WebAddress},
		{"bin", market.BIN},
		{"update_order_url", market.UpdateOrderURL},
		{"contact", market.Contact},
		{"enabled", market.Enabled},
		{"updated_at", time.Now().UTC()},
	}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return "", drivers.ErrDoesNotExist
	case errors.Is(err, nil):
		return market.ID, nil
	}
	return "", err
}
