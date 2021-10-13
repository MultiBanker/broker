package agree

import (
	"context"

	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
)

type Specification interface {
	Get(ctx context.Context, id string) (models.Specification, error)
	GetByCode(ctx context.Context, code string) (models.Specification, error)
	List(ctx context.Context, pgn selector.Paging) ([]models.Specification, int64, error)
	Create(ctx context.Context, spec models.Specification) (string, error)
	Update(ctx context.Context, spec models.Specification) error
}

type SpecImpl struct {
	specsRepo repository.AgreeSpecifications
}

func NewManager(repos repository.Repositories) *SpecImpl {
	return &SpecImpl{
		specsRepo: repos.AgreeSpecifications(),
	}
}

func (s SpecImpl) Get(ctx context.Context, id string) (models.Specification, error) {
	return s.specsRepo.Get(ctx, id)
}

func (s SpecImpl) GetByCode(ctx context.Context, code string) (models.Specification, error) {
	return s.specsRepo.GetByCode(ctx, code)
}

func (s SpecImpl) List(ctx context.Context, pgn selector.Paging) ([]models.Specification, int64, error) {
	return s.specsRepo.List(ctx, pgn)
}

func (s SpecImpl) Create(ctx context.Context, spec models.Specification) (string, error) {
	return s.specsRepo.Insert(ctx, spec)
}

func (s SpecImpl) Update(ctx context.Context, spec models.Specification) error {
	return s.specsRepo.Update(ctx, spec)
}