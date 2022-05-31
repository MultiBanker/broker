package manager

import (
	"github.com/MultiBanker/broker/pkg/auth"
	"github.com/MultiBanker/broker/src/config"
	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/manager/auto"
	"github.com/MultiBanker/broker/src/manager/loan"
	"github.com/MultiBanker/broker/src/manager/market"
	"github.com/MultiBanker/broker/src/manager/offer"
	"github.com/MultiBanker/broker/src/manager/order"
	"github.com/MultiBanker/broker/src/manager/partner"
	"github.com/MultiBanker/broker/src/manager/user"
	"github.com/VictoriaMetrics/metrics"
)

type Managers struct {
	DB                 drivers.Datastore
	PartnerMan         partner.Partnerer
	AuthMan            auth.Authenticator
	OrderMan           order.Orderer
	MarketMan          market.Marketer
	OfferMan           offer.Manager
	LoanMan            loan.Program
	UserMan            user.UsersManager
	UserApplicationMan user.ApplicationManager
	LoginMan           user.LogInManager
	RecoveryMan        user.RecoveryManager
	VerifyMan          user.VerificationManager
	AutoMan            auto.Auto
	MetricMan          *metrics.Set
}

func NewWrapper(db drivers.Datastore, repo repository.Repositories, opts *config.Config, metric *metrics.Set) Managers {
	return Managers{
		DB:          db,
		PartnerMan:  partner.NewPartner(repo),
		AuthMan:     auth.NewAuthenticator([]byte(opts.JWTKey), opts.AccessTokenTime, opts.RefreshTokenTime),
		OrderMan:    order.NewOrder(repo),
		MarketMan:   market.NewMarket(repo),
		OfferMan:    offer.NewOffer(repo),
		LoanMan:     loan.NewProgramManager(repo),
		UserMan:     user.NewUsersManagerImpl(repo.User),
		LoginMan:    user.NewLogInManagerImpl(repo.User),
		RecoveryMan: user.NewRecoveryManagerImpl(false, repo.Recovery, repo.User),
		VerifyMan:   user.NewVerificationManagerImpl(false, repo.Verify, repo.User),
		MetricMan:   metric,
	}
}
