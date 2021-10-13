package repository

import (
	"github.com/MultiBanker/broker/src/database/repository/mongo/agreement"
	"github.com/MultiBanker/broker/src/database/repository/mongo/market"
	"github.com/MultiBanker/broker/src/database/repository/mongo/offer"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/database/repository/mongo/order"
	"github.com/MultiBanker/broker/src/database/repository/mongo/partner"
	"github.com/MultiBanker/broker/src/database/repository/mongo/sequence"
)

const (
	Order              = "order"
	PartnerOrder       = "partner_order"
	Sequence           = "counters"
	Partner            = "partner"
	Market             = "market"
	OfferColl          = "offer"
	AgreeSpecColl      = "agree"
	SignatureAgreeColl = "agree_sign"
)

type Repositories interface {
	PartnerRepo() Partnerer
	SequenceRepo() Sequencer
	OrderRepo() Orderer
	MarketRepo() Marketer
	PartnerOrderRepo() PartnerOrderer
	OfferRepo() Offer
	AgreeSpecifications() AgreeSpecifications
	AgreeSignatures() AgreeSignatures
}

type Repository struct {
	Partner        Partnerer
	Sequence       Sequencer
	Order          Orderer
	Market         Marketer
	PartnerOrder   PartnerOrderer
	Offer          Offer
	AgreeSpec      AgreeSpecifications
	AgreeSignature AgreeSignatures
}

func NewRepository(datastore drivers.Datastore) (Repositories, error) {
	if datastore.Name() == "mongo" {
		db := datastore.Database().(*mongo.Database)
		return &Repository{
			Sequence:       sequence.NewRepository(db.Collection(Sequence)),
			Partner:        partner.NewRepository(db.Collection(Partner)),
			Order:          order.NewRepository(db.Collection(Order)),
			Market:         market.NewRepository(db.Collection(Market)),
			PartnerOrder:   order.NewPartnerOrderRepository(db.Collection(PartnerOrder)),
			Offer:          offer.NewRepository(db.Collection(OfferColl)),
			AgreeSpec:      agreement.NewSpecification(db.Collection(AgreeSpecColl)),
			AgreeSignature: agreement.NewSignature(db.Collection(SignatureAgreeColl)),
		}, nil
	}
	return nil, ErrDatastoreNotImplemented
}
