package offer

import (
	"context"
	"strconv"

	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
)

type Offer struct {
	seqColl     repository.Sequencer
	offerColl   repository.Offer
	partnerColl repository.Partnerer
}

func NewOffer(repo repository.Repositories) *Offer {
	return &Offer{
		seqColl:     repo.SequenceRepo(),
		offerColl:   repo.OfferRepo(),
		partnerColl: repo.PartnerRepo(),
	}
}

var _ Manager = (*Offer)(nil)

type Manager interface {
	CreateOffer(ctx context.Context, offer models.Offer) (string, error)
	UpdateOffer(ctx context.Context, offer models.Offer) (models.Offer, error)
	OfferByCode(ctx context.Context, code string) (models.Offer, error)
	Offers(ctx context.Context, paging selector.Paging) ([]models.Offer, int64, error)
}

func (o Offer) CreateOffer(ctx context.Context, offer models.Offer) (string, error) {
	_, err := o.partnerColl.PartnerByCode(ctx, offer.PartnerCode)
	if err != nil {
		return "", err
	}
	idInt, err := o.seqColl.NextSequenceValue(ctx, models.OfferSequences)
	if err != nil {
		return "", err
	}
	offer.ID = strconv.Itoa(idInt)
	return o.offerColl.CreateOffer(ctx, offer)
}

func (o Offer) UpdateOffer(ctx context.Context, offer models.Offer) (models.Offer, error) {
	_, err := o.partnerColl.PartnerByCode(ctx, offer.PartnerCode)
	if err != nil {
		return models.Offer{}, err
	}
	return o.offerColl.UpdateOffer(ctx, offer)
}

func (o Offer) OfferByCode(ctx context.Context, code string) (models.Offer, error) {
	return o.offerColl.OfferByCode(ctx, code)
}

func (o Offer) Offers(ctx context.Context, paging selector.Paging) ([]models.Offer, int64, error) {
	return o.offerColl.Offers(ctx, paging)
}
