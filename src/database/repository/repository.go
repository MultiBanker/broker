package repository

import (
	"github.com/MultiBanker/broker/src/database/repository/mongo/auto"
	"github.com/MultiBanker/broker/src/database/repository/mongo/loan"
	"github.com/MultiBanker/broker/src/database/repository/mongo/market"
	"github.com/MultiBanker/broker/src/database/repository/mongo/offer"
	"github.com/MultiBanker/broker/src/database/repository/mongo/user"
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
	LoanPrograms = "loan-programs"
	User         = "user"
	Recovery     = "recovery"
	Verify       = "verify"
	AutoColl     = "auto"
)

type Repositories struct {
	Partner      Partnerer
	Sequence     Sequencer
	Order        Orderer
	Market       Marketer
	PartnerOrder PartnerOrderer
	Offer        Offer
	LoanProgram  LoanProgram
	User         UsersRepository
	Verify       VerificationRepository
	Recovery     RecoveryRepository
	Auto         Auto
}

func NewRepository(datastore drivers.Datastore) (Repositories, error) {
	if datastore.Name() == "mongo" {
		db := datastore.Database().(*mongo.Database)
		return Repositories{
			Sequence:     sequence.NewRepository(db.Collection(Sequence)),
			Partner:      partner.NewRepository(db.Collection(Partner)),
			Order:        order.NewRepository(db.Collection(Order)),
			Market:       market.NewRepository(db.Collection(Market)),
			PartnerOrder: order.NewPartnerOrderRepository(db.Collection(PartnerOrder)),
			Offer:        offer.NewRepository(db.Collection(OfferColl)),
			LoanProgram:  loan.NewProgramRepository(db.Collection(LoanPrograms)),
			User:         user.NewUsersRepositoryImpl(db.Collection(User)),
			Verify:       user.NewVerificationRepositoryImpl(db.Collection(Verify)),
			Recovery:     user.NewRecoveryRepositoryImpl(db.Collection(Recovery)),
			Auto:         auto.NewRepository(db.Collection(AutoColl)),
		}, nil
	}
	return Repositories{}, ErrDatastoreNotImplemented
}
