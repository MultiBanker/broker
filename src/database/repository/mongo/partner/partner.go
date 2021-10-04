package partner

import (
	"context"

	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/models/selector"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/MultiBanker/broker/src/models"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) repository.Partnerer {
	return &Repository{collection: collection}
}

func (p Repository) NewPartner(ctx context.Context, partner *models.Partner) (string, error) {
	panic("implement me")
}

func (p Repository) UpdatePartner(ctx context.Context, partner *models.Partner) (string, error) {
	panic("implement me")
}

func (p Repository) PartnerByID(ctx context.Context, id string) (models.Partner, error) {
	panic("implement me")
}

func (p Repository) PartnerByUsername(ctx context.Context, id string) (models.Partner, error) {
	panic("implement me")
}

func (p Repository) Partners(ctx context.Context, paging *selector.Paging) ([]models.Partner, int64, error) {
	panic("implement me")
}
