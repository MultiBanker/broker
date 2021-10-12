package repository

import (
	"context"

	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
)

func (r Repository) OfferRepo() Offer {
	return r.Offer
}

type Offer interface {
	CreateOffer(ctx context.Context, offer models.Offer) (string, error)
	UpdateOffer(ctx context.Context, offer models.Offer) (models.Offer, error)
	OfferByCode(ctx context.Context, code string) (models.Offer, error)
	Offers(ctx context.Context, paging selector.Paging) ([]models.Offer, int64, error)
}
