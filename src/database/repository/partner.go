package repository

import (
	"context"

	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
)

type Partnerer interface {
	NewPartner(ctx context.Context, partner *models.Partner) (string, error)
	UpdatePartner(ctx context.Context, partner *models.Partner) (string, error)
	PartnerByID(ctx context.Context, id string) (models.Partner, error)
	PartnerByCode(ctx context.Context, id string) (models.Partner, error)
	PartnerByUsername(ctx context.Context, id string) (models.Partner, error)
	Partners(ctx context.Context, paging *selector.Paging) ([]models.Partner, int64, error)
}
