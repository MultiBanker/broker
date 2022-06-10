package market

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/database/repository/mongo/transaction"
	"github.com/MultiBanker/broker/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AutoRepository struct {
	coll        *mongo.Collection
	transaction transaction.Func
}

func NewAutoRepository(coll *mongo.Collection, transaction transaction.Func) *AutoRepository {
	return &AutoRepository{coll: coll, transaction: transaction}
}

func (a AutoRepository) Create(ctx context.Context, auto models.MarketAuto) (string, error) {
	auto.ID = primitive.NewObjectID().Hex()
	now := time.Now().UTC()
	auto.CreatedAt = now
	auto.UpdatedAt = now

	if _, err := a.coll.InsertOne(ctx, auto); err != nil {
		return "", fmt.Errorf("create auto: %w", err)
	}
	return auto.ID, nil
}

func (a AutoRepository) Get(ctx context.Context, sku string) (models.MarketAuto, error) {
	var auto models.MarketAuto

	filter := bson.D{
		{Key: "sku", Value: sku},
	}

	if err := a.coll.FindOne(ctx, filter).Decode(&auto); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return auto, drivers.ErrDoesNotExist
		}
		return auto, err
	}

	return auto, nil
}

func (a AutoRepository) Lock(ctx context.Context, id string) (models.MarketAuto, error) {
	panic("implement me")
}
