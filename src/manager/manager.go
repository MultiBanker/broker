package manager

import (
	"github.com/MultiBanker/broker/src/config"
	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/manager/auth"
	"github.com/MultiBanker/broker/src/manager/order"
	"github.com/MultiBanker/broker/src/manager/partner"
)

type Abstractor interface {
	Auther() auth.Authenticator
	Partnerer() partner.Partnerer
	Orderer() order.Orderer
	Pinger() error
}

type Abstract struct {
	db         drivers.Datastore
	partnerMan partner.Partnerer
	authMan    auth.Authenticator
	orderMan   order.Orderer
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

func NewAbstract(db drivers.Datastore, repo repository.Repositories, opts *config.Config) Abstractor {
	return &Abstract{
		db:         db,
		partnerMan: partner.NewPartner(repo),
		authMan:    auth.NewAuthenticator(opts),
		orderMan:   order.NewOrder(repo),
	}
}
