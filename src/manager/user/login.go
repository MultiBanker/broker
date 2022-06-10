package user

import (
	"context"

	"github.com/MultiBanker/broker/pkg/auth"
	"github.com/MultiBanker/broker/src/database/repository"
)

type LogInManager interface {
	SignInByPhone(ctx context.Context, phone, password string) (string, error)
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
	user, err := l.usersRepo.GetByPhone(ctx, phone)
	if err != nil {
		return "", err
	}

	//if !user.IsEnabled {
	//	return "", ErrUserDisabled
	//}

	if !user.IsPhoneVerified {
		return "", ErrUserPhoneIsNotVerified
	}

	if !auth.CheckPasswordHash(password, user.Password) {
		return "", ErrInvalidLoginOrPassword
	}

	return user.ID, nil
}
