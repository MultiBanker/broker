package repository

import (
	"github.com/MultiBanker/broker/src/database/repository/mongo/market"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/database/repository/mongo/order"
	"github.com/MultiBanker/broker/src/database/repository/mongo/partner"
	"github.com/MultiBanker/broker/src/database/repository/mongo/sequence"
)

const (
	Order    = "order"
	Sequence = "counters"
	Partner  = "partner"
	Market   = "market"
)

type Repositories interface {
	PartnerRepo() Partnerer
	SequenceRepo() Sequencer
	OrderRepo() Orderer
	MarketRepo() Marketer
}

type Repository struct {
	Partner  Partnerer
	Sequence Sequencer
	Order    Orderer
	Market   Marketer
}

func NewRepository(datastore drivers.Datastore) (Repositories, error) {
	if datastore.Name() == "mongo" {
		db := datastore.Database().(*mongo.Database)
		return &Repository{
			Sequence: sequence.NewRepository(db.Collection(Sequence)),
			Partner:  partner.NewRepository(db.Collection(Partner)),
			Order:    order.NewRepository(db.Collection(Order)),
			Market:   market.NewRepository(db.Collection(Market)),
		}, nil
	}
	return nil, ErrDatastoreNotImplemented
}
