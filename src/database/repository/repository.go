package repository

import (
	"github.com/MultiBanker/broker/src/database/repository/mongo/auto"
	"github.com/MultiBanker/broker/src/database/repository/mongo/loan"
	"github.com/MultiBanker/broker/src/database/repository/mongo/market"
	"github.com/MultiBanker/broker/src/database/repository/mongo/offer"
	"github.com/MultiBanker/broker/src/database/repository/mongo/user"
	"github.com/MultiBanker/broker/src/database/repository/mongo/userauto"
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
	UserAuto     = "user_auto"
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
	UserAuto     UserAutoRepository
}

func NewRepository(datastore drivers.Datastore) (Repositories, error) {
	if datastore.Name() == "mongo" {
		db := datastore.Database().(*mongo.Database)
		return Repositories{
			Sequence:     sequence.NewRepository(db.Collection(Sequence), datastore.WithTransaction()),
			Partner:      partner.NewRepository(db.Collection(Partner), datastore.WithTransaction()),
			Order:        order.NewRepository(db.Collection(Order), datastore.WithTransaction()),
			Market:       market.NewRepository(db.Collection(Market), datastore.WithTransaction()),
			PartnerOrder: order.NewPartnerOrderRepository(db.Collection(PartnerOrder), datastore.WithTransaction()),
			Offer:        offer.NewRepository(db.Collection(OfferColl), datastore.WithTransaction()),
			LoanProgram:  loan.NewProgramRepository(db.Collection(LoanPrograms), datastore.WithTransaction()),
			User:         user.NewUsersRepositoryImpl(db.Collection(User), datastore.WithTransaction()),
			Verify:       user.NewVerificationRepositoryImpl(db.Collection(Verify), datastore.WithTransaction()),
			Recovery:     user.NewRecoveryRepositoryImpl(db.Collection(Recovery), datastore.WithTransaction()),
			Auto:         auto.NewRepository(db.Collection(AutoColl), datastore.WithTransaction()),
			UserAuto:     userauto.NewUserAutoRepository(db.Collection(UserAuto), datastore.WithTransaction()),
		}, nil
	}
	return Repositories{}, ErrDatastoreNotImplemented
}
