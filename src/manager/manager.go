package manager

import (
	"github.com/MultiBanker/broker/src/config"
	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/manager/agree"
	"github.com/MultiBanker/broker/src/manager/auth"
	"github.com/MultiBanker/broker/src/manager/market"
	"github.com/MultiBanker/broker/src/manager/offer"
	"github.com/MultiBanker/broker/src/manager/order"
	"github.com/MultiBanker/broker/src/manager/partner"
	"github.com/MultiBanker/broker/src/manager/signature"
)

type Abstractor interface {
	Auther() auth.Authenticator
	Partnerer() partner.Partnerer
	Orderer() order.Orderer
	Marketer() market.Marketer
	Offer() offer.Manager
	Pinger() error
	Agree() agree.Specification
	Signature() signature.Signature
}

type Abstract struct {
	db         drivers.Datastore
	partnerMan partner.Partnerer
	authMan    auth.Authenticator
	orderMan   order.Orderer
	marketMan  market.Marketer
	offerMan   offer.Manager
	agreeMan   agree.Specification
	signMan    signature.Signature
}

func (a Abstract) Agree() agree.Specification {
	return a.agreeMan
}

func (a Abstract) Signature() signature.Signature {
	return a.signMan
}

func (a Abstract) Pinger() error {
	return a.db.Ping()
}

func (a *Abstract) Partnerer() partner.Partnerer {
	return a.partnerMan
}

func (a *Abstract) Auther() auth.Authenticator {
	return a.authMan
}

func (a *Abstract) Orderer() order.Orderer {
	return a.orderMan
}

func (a *Abstract) Marketer() market.Marketer {
	return a.marketMan
}

func (a Abstract) Offer() offer.Manager {
	return a.offerMan
}

func NewAbstract(db drivers.Datastore, repo repository.Repositories, opts *config.Config) Abstractor {
	return &Abstract{
		db:         db,
		partnerMan: partner.NewPartner(repo),
		authMan:    auth.NewAuthenticator(opts),
		orderMan:   order.NewOrder(repo),
		marketMan:  market.NewMarket(repo),
		offerMan:   offer.NewOffer(repo),
	}
}
