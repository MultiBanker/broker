package repository

import (
	"context"

	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
)

func (r Repository) AgreeSpecifications() AgreeSpecifications {
	return r.AgreeSpec
}

func (r Repository) AgreeSignatures() AgreeSignatures {
	return r.AgreeSignature
}

type AgreeSpecifications interface {
	List(ctx context.Context, pgn selector.Paging) ([]models.Specification, int64, error)
	Get(ctx context.Context, id string) (models.Specification, error)
	GetByCode(ctx context.Context, code string) (models.Specification, error)
	Insert(ctx context.Context, spec models.Specification) (string, error)
	Update(ctx context.Context, spec models.Specification) error
}

type AgreeSignatures interface {
	List(ctx context.Context, pgn selector.Paging) ([]models.Signature, int64, error)
	Get(ctx context.Context, id string) (models.Signature, error)
	Insert(ctx context.Context, signature models.Signature) (string, error)
	NewVerificationStatus(ctx context.Context) (models.Signature, error)
	RetrySign(ctx context.Context, signID, token string) (string, error)
	UpdateVerification(ctx context.Context, signature models.Signature) error
	GetVerifiedByCode(ctx context.Context, tdid, code string) (models.Signature, error)
}
