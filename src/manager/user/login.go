package user

import (
	"context"

	"github.com/MultiBanker/broker/src/database/repository"
)

type LogInManager interface {
	SignInByPhone(ctx context.Context, phone, password string) (string, error)
	SignInByEmail(ctx context.Context, email, password string) (string, error)
}

type LogInManagerImpl struct {
	usersRepo repository.UsersRepository
}

func NewLogInManagerImpl(usersRepo repository.UsersRepository) *LogInManagerImpl {
	return &LogInManagerImpl{
		usersRepo: usersRepo,
	}
}

func (l LogInManagerImpl) SignInByPhone(ctx context.Context, phone, password string) (string, error) {
	panic("implement me")
}

func (l LogInManagerImpl) SignInByEmail(ctx context.Context, email, password string) (string, error) {
	panic("implement me")
}