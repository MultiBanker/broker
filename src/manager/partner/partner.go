package partner

import (
	"context"
	"fmt"

	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
)

type Partnerer interface {
	NewPartner(ctx context.Context, partner *models.Partner) (string, error)
	UpdatePartner(ctx context.Context, partner *models.Partner) (string, error)
	PartnerByID(ctx context.Context, id string) (models.Partner, error)
	Partners(ctx context.Context, paging *selector.Paging) ([]models.Partner, error)
	PartnerByUsername(ctx context.Context, username string) (models.Partner, error)
}

type Partner struct {
	sequenceColl repository.SequencesRepository
	partnerColl  repository.PartnerRepository
}

func NewPartner(repos repository.Repositories) Partner {
	return Partner{
		sequenceColl: repos.SequenceRepo(),
		partnerColl:  repos.PartnerRepo(),
	}
}

func (p Partner) NewPartner(ctx context.Context, partner *models.Partner) (string, error) {
	bytePass, err := HashPassword(partner.Password)
	if err != nil {
		return "", err
	}
	partner.Password = string(bytePass)
	return p.partnerColl.NewPartner(ctx, partner)
}

func (p Partner) UpdatePartner(ctx context.Context, partner *models.Partner) (string, error) {
	res, err := p.partnerColl.PartnerByID(ctx, partner.ID)
	if err != nil {
		return "", err
	}
	if !CheckPasswordHash(partner.Password, []byte(res.Password)) {
		return "", fmt.Errorf("[ERROR] Wrong Password")
	}
	return p.partnerColl.UpdatePartner(ctx, partner)
}

func (p Partner) PartnerByID(ctx context.Context, id string) (models.Partner, error) {
	return p.partnerColl.PartnerByID(ctx, id)
}

func (p Partner) Partners(ctx context.Context, paging *selector.Paging) ([]models.Partner, error) {
	return p.partnerColl.Partners(ctx, paging)
}

func (p Partner) PartnerByUsername(ctx context.Context, username string) (models.Partner, error) {
	return p.partnerColl.PartnerByUsername(ctx, username)
}
