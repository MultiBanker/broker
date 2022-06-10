package auto

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/database/repository/mongo/transaction"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	coll        *mongo.Collection
	transaction transaction.Func
}

func NewRepository(coll *mongo.Collection, transaction transaction.Func) *Repository {
	return &Repository{coll: coll, transaction: transaction}
}

func (r Repository) Get(ctx context.Context, sku string) (models.Auto, error) {
	var auto models.Auto

	filter := bson.D{
		{Key: "sku", Value: sku},
	}

	if err := r.coll.FindOne(ctx, filter).Decode(&auto); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return auto, drivers.ErrDoesNotExist
		}
		return auto, err
	}

	return auto, nil
}

func (r Repository) Create(ctx context.Context, auto models.Auto) (string, error) {
	auto.ID = primitive.NewObjectID().Hex()
	now := time.Now().UTC()
	auto.CreatedAt = now
	auto.UpdatedAt = now

	if _, err := r.coll.InsertOne(ctx, auto); err != nil {
		return "", fmt.Errorf("create auto: %w", err)
	}
	return auto.ID, nil
}

func (r Repository) Update(ctx context.Context, auto models.Auto) (string, error) {
	filter := bson.D{
		{Key: "sku", Value: auto.SKU},
	}
	update := bson.D{
		{Key: "title", Value: auto.Title},
		{Key: "brand", Value: auto.Brand},
		{Key: "color", Value: auto.Color},
		{Key: "media", Value: auto.Media},
		{Key: "price", Value: auto.Price},
		{Key: "about", Value: auto.About},
		{Key: "updated_at", Value: time.Now().UTC()},
	}

	_, err := r.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return "", fmt.Errorf("update auto: %w", err)
	}
	return auto.SKU, nil
}

func (r Repository) List(ctx context.Context, query selector.SearchQuery) ([]models.Auto, int64, error) {
	filter := bson.D{}

	opts := &options.FindOptions{
		Skip: &query.Pagination.Page,
		Sort: bson.D{
			{Key: "created_at", Value: -1},
		},
		Limit: &query.Pagination.Limit,
	}
	total, err := r.coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	cur, err := r.coll.Find(ctx, filter, opts)
	switch err {
	case mongo.ErrNoDocuments:
		return nil, 0, drivers.ErrDoesNotExist
	case nil:
		autos := make([]models.Auto, cur.RemainingBatchLength())
		if err := cur.All(ctx, &autos); err != nil {
			return nil, 0, err
		}
		return autos, total, nil
	default:
		return nil, 0, err
	}
}

func (r Repository) Delete(ctx context.Context, sku string) error {
	filter := bson.D{
		{Key: "sku", Value: sku},
	}

	_, err := r.coll.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}
