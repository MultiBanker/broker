package repository

import (
	"context"

	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
)

type UsersRepository interface {
	Create(ctx context.Context, user models.User) (string, error)
	Get(ctx context.Context, paging *selector.SearchQuery) ([]models.User, error)
	GetOrCreateUserByPhone(ctx context.Context, phone string) (string, error)
	GetByID(ctx context.Context, userID string) (models.User, error)
	GetByIDs(ctx context.Context, userIDs []string) ([]models.User, error)
	Update(ctx context.Context, userID string, user models.User) error
	UpdatePassword(ctx context.Context, userID, password string) error
	UpdatePhone(ctx context.Context, userID, phone string) error
	Delete(ctx context.Context, userID string) error
	Count(ctx context.Context, paging *selector.SearchQuery) (int64, error)
}

type VerificationRepository interface {
	Create(ctx context.Context, verify models.Verification) error
	GetByChannel(ctx context.Context, channel string, destination string) (models.Verification, error)
	GetByID(ctx context.Context, verifyID string) (models.Verification, error)
	GetNewVerification(ctx context.Context) (models.Verification, error)
	RollbackVerification(ctx context.Context, verifyID string) error
	FinishVerification(ctx context.Context, verifyID string) error
	Update(ctx context.Context, verify models.Verification) error
	Delete(ctx context.Context, verifyID string) error
}

type RecoveryRepository interface {
	Create(ctx context.Context, recovery models.Recovery) error
	GetByChannel(ctx context.Context, channel string, destination string) (models.Recovery, error)
	GetByID(ctx context.Context, recoveryID string) (models.Recovery, error)
	GetNewRecovery(ctx context.Context) (models.Recovery, error)
	RollbackRecovery(ctx context.Context, recoveryID string) error
	FinishRecovery(ctx context.Context, recoveryID string) error
	Update(ctx context.Context, recovery models.Recovery) error
	Delete(ctx context.Context, recoveryID string) error
}
type UserApplicationRepository interface {
	Create(ctx context.Context, application models.UserApplication) (string, error)
	Get(ctx context.Context, sku string) (models.UserApplication, error)
	Lock(ctx context.Context, id string) (models.UserApplication, error)
}
