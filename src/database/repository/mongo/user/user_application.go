package user

import (
	"context"

	"github.com/MultiBanker/broker/src/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApplicationRepository struct {
	coll *mongo.Collection
}

func NewApplicationRepository(coll *mongo.Collection) *ApplicationRepository {
	return &ApplicationRepository{coll: coll}
}

func (a ApplicationRepository) Create(ctx context.Context, application models.UserApplication) (string, error) {
	panic("implement me")
}

func (a ApplicationRepository) Get(ctx context.Context, id string) (models.UserApplication, error) {
	panic("implement me")
}

func (a ApplicationRepository) Lock(ctx context.Context, id string) (models.UserApplication, error) {
	panic("implement me")
}