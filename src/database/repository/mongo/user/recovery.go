package user

import (
	"context"

	"github.com/MultiBanker/broker/src/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecoveryRepositoryImpl struct {
	coll *mongo.Collection
}

func NewRecoveryRepositoryImpl(coll *mongo.Collection) *RecoveryRepositoryImpl {
	return &RecoveryRepositoryImpl{
		coll: coll,
	}
}

func (r RecoveryRepositoryImpl) Create(ctx context.Context, recovery models.Recovery) error {
	panic("implement me")
}

func (r RecoveryRepositoryImpl) GetByChannel(ctx context.Context, channel string, destination string) (models.Recovery, error) {
	panic("implement me")
}

func (r RecoveryRepositoryImpl) GetByID(ctx context.Context, recoveryID string) (models.Recovery, error) {
	panic("implement me")
}

func (r RecoveryRepositoryImpl) GetNewRecovery(ctx context.Context) (models.Recovery, error) {
	panic("implement me")
}

func (r RecoveryRepositoryImpl) RollbackRecovery(ctx context.Context, recoveryID string) error {
	panic("implement me")
}

func (r RecoveryRepositoryImpl) FinishRecovery(ctx context.Context, recoveryID string) error {
	panic("implement me")
}

func (r RecoveryRepositoryImpl) Update(ctx context.Context, recovery models.Recovery) error {
	panic("implement me")
}

func (r RecoveryRepositoryImpl) Delete(ctx context.Context, recoveryID string) error {
	panic("implement me")
}