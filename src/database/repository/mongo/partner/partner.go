package partner

import (
	"context"

	"github.com/MultiBanker/broker/src/models/selector"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/MultiBanker/broker/src/models"
)

type PartnerRepository struct {
	collection *mongo.Collection
}

func NewPartnerRepository(collection *mongo.Collection) *PartnerRepository {
	return &PartnerRepository{collection: collection}
}

func (p PartnerRepository) NewPartner(ctx context.Context, partner *models.Partner) (string, error) {
	panic("implement me")
}

func (p PartnerRepository) UpdatePartner(ctx context.Context, partner *models.Partner) (string, error) {
	panic("implement me")
}

func (p PartnerRepository) PartnerByID(ctx context.Context, id string) (models.Partner, error) {
	panic("implement me")
}

func (p PartnerRepository) PartnerByUsername(ctx context.Context, id string) (models.Partner, error) {
	panic("implement me")
}

func (p PartnerRepository) Partners(ctx context.Context, paging *selector.Paging) ([]models.Partner, error) {
	panic("implement me")
}

