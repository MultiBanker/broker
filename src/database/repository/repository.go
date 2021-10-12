package repository

import (
	"github.com/MultiBanker/broker/src/database/repository/mongo/market"
	"github.com/MultiBanker/broker/src/database/repository/mongo/offer"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/database/repository/mongo/order"
	"github.com/MultiBanker/broker/src/database/repository/mongo/partner"
	"github.com/MultiBanker/broker/src/database/repository/mongo/sequence"
)

const (
	Order        = "order"
	PartnerOrder = "partner_order"
	Sequence     = "counters"
	Partner      = "partner"
	Market       = "market"
	OfferColl    = "offer"
)

type Repositories interface {
	PartnerRepo() Partnerer
	SequenceRepo() Sequencer
	OrderRepo() Orderer
	MarketRepo() Marketer
	PartnerOrderRepo() PartnerOrderer
	OfferRepo() Offer
}

type Repository struct {
	Partner      Partnerer
	Sequence     Sequencer
	Order        Orderer
	Market       Marketer
	PartnerOrder PartnerOrderer
	Offer        Offer
}

func NewRepository(datastore drivers.Datastore) (Repositories, error) {
	if datastore.Name() == "mongo" {
		db := datastore.Database().(*mongo.Database)
		return &Repository{
			Sequence:     sequence.NewRepository(db.Collection(Sequence)),
			Partner:      partner.NewRepository(db.Collection(Partner)),
			Order:        order.NewRepository(db.Collection(Order)),
			Market:       market.NewRepository(db.Collection(Market)),
			PartnerOrder: order.NewPartnerOrderRepository(db.Collection(PartnerOrder)),
			Offer:        offer.NewRepository(db.Collection(OfferColl)),
		}, nil
	}
	return nil, ErrDatastoreNotImplemented
}
