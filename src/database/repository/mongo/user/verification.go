package user

import (
	"context"

	"github.com/MultiBanker/broker/src/database/repository/mongo/transaction"
	"github.com/MultiBanker/broker/src/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type VerificationRepositoryImpl struct {
	coll        *mongo.Collection
	transaction transaction.Func
}

func NewVerificationRepositoryImpl(coll *mongo.Collection, transaction transaction.Func) *VerificationRepositoryImpl {
	return &VerificationRepositoryImpl{
		transaction: transaction,
		coll:        coll,
	}
}

func (v VerificationRepositoryImpl) Create(ctx context.Context, verify models.Verification) error {
	panic("implement me")
}

func (v VerificationRepositoryImpl) GetByChannel(ctx context.Context, channel string, destination string) (models.Verification, error) {
	panic("implement me")
}

func (v VerificationRepositoryImpl) GetByID(ctx context.Context, verifyID string) (models.Verification, error) {
	panic("implement me")
}

func (v VerificationRepositoryImpl) GetNewVerification(ctx context.Context) (models.Verification, error) {
	panic("implement me")
}

func (v VerificationRepositoryImpl) RollbackVerification(ctx context.Context, verifyID string) error {
	panic("implement me")
}

func (v VerificationRepositoryImpl) FinishVerification(ctx context.Context, verifyID string) error {
	panic("implement me")
}

func (v VerificationRepositoryImpl) Update(ctx context.Context, verify models.Verification) error {
	panic("implement me")
}

func (v VerificationRepositoryImpl) Delete(ctx context.Context, verifyID string) error {
	panic("implement me")
}
