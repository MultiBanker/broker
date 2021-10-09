package repository

import (
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
)

type Repositories interface {
	PartnerRepo() Partnerer
	SequenceRepo() Sequencer
	OrderRepo() Orderer
	PartnerOrderRepo() PartnerOrderer
}

type Repository struct {
	Partner      Partnerer
	Sequence     Sequencer
	Order        Orderer
	PartnerOrder PartnerOrderer
}

func NewRepository(datastore drivers.Datastore) (Repositories, error) {
	if datastore.Name() == "mongo" {
		db := datastore.Database().(*mongo.Database)
		return &Repository{
			Sequence:     sequence.NewRepository(db.Collection(Sequence)),
			Partner:      partner.NewRepository(db.Collection(Partner)),
			Order:        order.NewRepository(db.Collection(Order)),
			PartnerOrder: order.NewPartnerOrderRepository(db.Collection(PartnerOrder)),
		}, nil
	}
	return nil, ErrDatastoreNotImplemented
}
