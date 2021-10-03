package repository

import (
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
)

type Repositories interface {
	PartnerRepo() PartnerRepository
	SequenceRepo() SequencesRepository
	OrderRepo() OrderRepository
}

type Repository struct {
	Partner  PartnerRepository
	Sequence SequencesRepository
	Order    OrderRepository
}

func NewRepository(datastore drivers.Datastore) (Repositories, error) {
	if datastore.Name() == "mongo" {
		db := datastore.Database().(*mongo.Database)
		return &Repository{
			Sequence: sequence.NewSequencesRepository(db.Collection(Sequence)),
			Partner:  partner.NewPartnerRepository(db.Collection(Partner)),
			Order:    order.NewOrderRepository(db.Collection(Order)),
		}, nil
	}
	return nil, ErrDatastoreNotImplemented
}
