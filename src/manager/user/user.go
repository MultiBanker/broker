package user

import (
	"context"

	"github.com/MultiBanker/broker/pkg/auth"
	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
)

type UsersManager interface {
	Create(ctx context.Context, user models.User) (string, error)
	Get(ctx context.Context, query *selector.SearchQuery) ([]models.UserInfo, error)
	Count(ctx context.Context, query *selector.SearchQuery) (int64, error)
	GetOrCreateUserByPhone(ctx context.Context, phone string) (string, error)
	GetByUserID(ctx context.Context, userID string) (models.User, error)
	GetByUserIDs(ctx context.Context, userIDs []string) ([]models.User, error)
	Update(ctx context.Context, userID string, updatedUser models.User) (models.UserInfo, error)
	UpdatePassword(ctx context.Context, userID, password string) error
	UpdatePhone(ctx context.Context, userID, phone string) error
	Delete(ctx context.Context, userID string) error
	EnableUser(ctx context.Context, userID string) error
	DisableUser(ctx context.Context, userID string) error
}

type UsersManagerImpl struct {
	usersRepo repository.UsersRepository
}

func NewUsersManagerImpl(usersRepo repository.UsersRepository) *UsersManagerImpl {
	return &UsersManagerImpl{
		usersRepo: usersRepo,
	}
}

func (man *UsersManagerImpl) Create(ctx context.Context, user models.User) (string, error) {
	return man.usersRepo.Create(ctx, user)
}

func (man *UsersManagerImpl) Get(ctx context.Context, query *selector.SearchQuery) ([]models.UserInfo, error) {
	usersInfo := make([]models.UserInfo, 0)
	users, err := man.usersRepo.Get(ctx, query)
	if err != nil {
		return usersInfo, err
	}

	for _, user := range users {
		userInfo := models.UserInfo{
			ID:              user.ID,
			FirstName:       user.FirstName,
			LastName:        user.LastName,
			Patronymic:      user.Patronymic,
			Phone:           user.Phone,
			Email:           user.Email,
			CreatedAt:       user.CreatedAt,
			UpdatedAt:       user.UpdatedAt,
			IsEnabled:       user.IsEnabled,
			IsEmailVerified: user.IsEmailVerified,
			IsPhoneVerified: user.IsPhoneVerified,
		}

		usersInfo = append(usersInfo, userInfo)
	}

	return usersInfo, nil
}

func (man *UsersManagerImpl) Count(ctx context.Context, query *selector.SearchQuery) (int64, error) {
	return man.usersRepo.Count(ctx, query)
}

func (man *UsersManagerImpl) GetOrCreateUserByPhone(ctx context.Context, phone string) (string, error) {
	return man.usersRepo.GetOrCreateUserByPhone(ctx, phone)
}

func (man *UsersManagerImpl) GetByUserID(ctx context.Context, userID string) (models.User, error) {
	return man.usersRepo.GetByID(ctx, userID)
}

func (man *UsersManagerImpl) GetByUserIDs(ctx context.Context, userIDs []string) ([]models.User, error) {
	return man.usersRepo.GetByIDs(ctx, userIDs)
}

func (man *UsersManagerImpl) Update(ctx context.Context, userID string, user models.User) (models.UserInfo, error) {
	if err := man.usersRepo.Update(ctx, userID, user); err != nil {
		return models.UserInfo{}, err
	}
	return models.UserInfo{}, nil
}

func (man *UsersManagerImpl) UpdatePassword(ctx context.Context, userID, password string) error {
	hashedPass, err := auth.HashPassword(password)
	if err != nil {
		return ErrAuthorization
	}
	return man.usersRepo.UpdatePassword(ctx, userID, string(hashedPass))
}

func (man *UsersManagerImpl) UpdatePhone(ctx context.Context, userID, phone string) error {
	return man.usersRepo.UpdatePhone(ctx, userID, phone)
}

func (man *UsersManagerImpl) Delete(ctx context.Context, userID string) error {
	return man.usersRepo.Delete(ctx, userID)
}

func (man *UsersManagerImpl) EnableUser(ctx context.Context, userID string) error {
	user, err := man.usersRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	user.IsEnabled = true
	err = man.usersRepo.Update(ctx, userID, user)
	if err != nil {
		return err
	}

	return nil
}

func (man *UsersManagerImpl) DisableUser(ctx context.Context, userID string) error {
	user, err := man.usersRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	user.IsEnabled = false
	err = man.usersRepo.Update(ctx, userID, user)
	if err != nil {
		return err
	}

	return nil
}
