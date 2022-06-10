package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/database/repository/mongo/transaction"
	"github.com/MultiBanker/broker/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApplicationRepository struct {
	coll        *mongo.Collection
	transaction transaction.Func
}

func NewApplicationRepository(coll *mongo.Collection, transaction transaction.Func) *ApplicationRepository {
	return &ApplicationRepository{coll: coll, transaction: transaction}
}

func (a ApplicationRepository) Create(ctx context.Context, application models.UserApplication) (string, error) {
	now := time.Now().UTC()
	application.CreatedAt = now
	application.UpdatedAt = now

	if _, err := a.coll.InsertOne(ctx, application); err != nil {
		return "", fmt.Errorf("create auto: %w", err)
	}
	return application.ApplicationID, nil
}

func (a ApplicationRepository) Get(ctx context.Context, id string) (models.UserApplication, error) {
	var userAppl models.UserApplication

	filter := bson.D{
		{Key: "application_id", Value: id},
	}

	if err := a.coll.FindOne(ctx, filter).Decode(&userAppl); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return userAppl, drivers.ErrDoesNotExist
		}
		return userAppl, err
	}

	return userAppl, nil
}

func (a ApplicationRepository) Lock(ctx context.Context, id string) (models.UserApplication, error) {
	panic("implement me")
}
