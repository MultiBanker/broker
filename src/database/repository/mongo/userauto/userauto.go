package userauto

import (
	"context"
	"errors"
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

type userAutoRepository struct {
	coll        *mongo.Collection
	transaction transaction.Func
}

func NewUserAutoRepository(coll *mongo.Collection, transaction transaction.Func) *userAutoRepository {
	return &userAutoRepository{coll: coll, transaction: transaction}
}

func (u userAutoRepository) Get(ctx context.Context, userID string) (models.UserAuto, error) {
	var userAuto models.UserAuto

	filter := bson.D{
		{Key: "user_id", Value: userID},
	}

	if err := u.coll.FindOne(ctx, filter).Decode(&userAuto); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return userAuto, drivers.ErrDoesNotExist
		}
		return userAuto, err
	}

	return userAuto, nil
}

func (u userAutoRepository) List(ctx context.Context, search selector.SearchQuery) ([]models.UserAuto, int64, error) {
	var opts *options.FindOptions

	if search.HasSorting() {
		opts.SetSort(bson.D{{Key: "created_at", Value: -1}}).
			SetSkip(search.Pagination.Page).
			SetLimit(search.Pagination.Limit)
	}
	count, err := u.coll.CountDocuments(ctx, bson.D{})
	if err != nil {
		return nil, 0, err
	}
	cur, err := u.coll.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, 0, err
	}
	userAutos := make([]models.UserAuto, 0, search.Pagination.Limit)
	err = cur.All(ctx, &userAutos)
	if err != nil {
		return nil, 0, err
	}

	return userAutos, count, nil
}

func (u userAutoRepository) Create(ctx context.Context, auto models.UserAuto) (string, error) {
	auto.ID = primitive.NewObjectID().Hex()
	auto.CreatedAt = time.Now().UTC()

	_, err := u.coll.InsertOne(ctx, auto)
	if err != nil {
		return "", err
	}
	return auto.ID, nil
}
