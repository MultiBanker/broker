package partner

import (
	"context"
	"fmt"
	"strconv"

	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/manager/auth"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
)

type Partnerer interface {
	NewPartner(ctx context.Context, partner *models.Partner) (string, error)
	UpdatePartner(ctx context.Context, partner *models.Partner) (string, error)
	PartnerByID(ctx context.Context, id string) (models.Partner, error)
	Partners(ctx context.Context, paging *selector.Paging) ([]models.Partner, int64, error)
	PartnerByUsername(ctx context.Context, username, password string) (models.Partner, error)
}

type Partner struct {
	sequenceColl repository.Sequencer
	partnerColl  repository.Partnerer
}

func NewPartner(repos repository.Repositories) Partner {
	return Partner{
		sequenceColl: repos.SequenceRepo(),
		partnerColl:  repos.PartnerRepo(),
	}
}

func (p Partner) NewPartner(ctx context.Context, partner *models.Partner) (string, error) {
	bytePass, err := auth.HashPassword(partner.Password)
	if err != nil {
		return "", err
	}
	partner.Password = string(bytePass)

	idInt, err := p.sequenceColl.NextSequenceValue(ctx, models.PartnerSequences)
	if err != nil {
		return "", err
	}
	partner.ID = strconv.Itoa(idInt)
	return p.partnerColl.NewPartner(ctx, partner)
}

func (p Partner) UpdatePartner(ctx context.Context, partner *models.Partner) (string, error) {
	res, err := p.partnerColl.PartnerByCode(ctx, partner.ID)
	if err != nil {
		return "", err
	}
	if !auth.CheckPasswordHash(partner.Password, []byte(res.Password)) {
		return "", fmt.Errorf("[ERROR] Wrong Password")
	}
	return p.partnerColl.UpdatePartner(ctx, partner)
}

func (p Partner) PartnerByID(ctx context.Context, id string) (models.Partner, error) {
	return p.partnerColl.PartnerByCode(ctx, id)
}

func (p Partner) Partners(ctx context.Context, paging *selector.Paging) ([]models.Partner, int64, error) {
	return p.partnerColl.Partners(ctx, paging)
}

func (p Partner) PartnerByUsername(ctx context.Context, username, password string) (models.Partner, error) {
	var partner models.Partner
	partner, err := p.partnerColl.PartnerByUsername(ctx, username)
	if err != nil {
		return partner, err
	}

	if !auth.CheckPasswordHash(password, []byte(partner.Password)) {
		return partner, fmt.Errorf("[ERROR] Wrong Password")
	}
	return partner, nil
}
