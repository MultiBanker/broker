package user

import (
	"context"
	"errors"
	"time"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/database/repository"
)

const (
	VerifyOtpTTL          = 5 * time.Minute  // Время жизни ОТП на верификацию
	VerifyLongSpamPenalty = 60 * time.Minute // Время длительного штрафа при попытках спама
	VerifySpamPenalty     = 60 * time.Second // Время короткого штрафа при попытках спама
	VerifyTriesLimit      = 3                // Количество попыток при вводе СМС
)

type VerificationManager interface {
	SendOTP(ctx context.Context, channel, destination string) error
	ValidateOTP(ctx context.Context, channel, destination, otp string) (string, error)
}

type VerificationManagerImpl struct {
	isTesting  bool
	verifyRepo repository.VerificationRepository
	usersRepo  repository.UsersRepository
}

func NewVerificationManagerImpl(
	isTesting bool,
	verifyRepo repository.VerificationRepository,
	usersRepo repository.UsersRepository) *VerificationManagerImpl {
	return &VerificationManagerImpl{
		isTesting:  isTesting,
		verifyRepo: verifyRepo,
		usersRepo:  usersRepo,
	}
}

func (v VerificationManagerImpl) SendOTP(ctx context.Context, channel, destination string) error {
	user, err := v.usersRepo.GetByPhone(ctx, destination)
	if err != nil {
		return err
	}

	if !user.IsEnabled {
		return ErrUserDisabled
	}

	verify, err := v.verifyRepo.GetByChannel(ctx, channel, destination)
	switch {
	case errors.Is(err, drivers.ErrDoesNotExist):
		return err
	case err == nil:
		if time.Now().UTC().Before(verify.NextAttemptAt.UTC()) {
			return TooManyRequestsError{
				NextAttemptAt: verify.NextAttemptAt.UTC().Format(time.RFC3339),
			}
		}

		verify.Status = "new"
		verify.Send = false
		verify.Count += 1
		verify.ExpiredAt = time.Now().Add(VerifyOtpTTL)
		verify.NextAttemptAt = time.Now().Add(VerifySpamPenalty)
		verify.UpdatedAt = time.Now().UTC()
		return v.verifyRepo.Update(ctx, verify)
	}

	return err
}

func (v VerificationManagerImpl) ValidateOTP(ctx context.Context, channel, destination, otp string) (string, error) {
	verify, err := v.verifyRepo.GetByChannel(ctx, channel, destination)
	if err != nil {
		return "", err
	}
	user, err := v.usersRepo.GetByPhone(ctx, destination)
	if err != nil {
		return "", err
	}

	// Проверка на корректность введённого пользователем OTP
	if verify.OTP != otp {
		// Проверка на количество введённого пользователем OTP
		if verify.Tries >= VerifyTriesLimit {

			verify.OTP = ""
			verify.Tries = 0
			verify.Count = 0
			verify.NextAttemptAt = time.Now().UTC().Add(VerifyLongSpamPenalty) // Увеличиваем время штрафа до 60 мин
			verify.UpdatedAt = time.Now().UTC()

			err = v.verifyRepo.Update(ctx, verify)
			if err != nil {
				return verify.UserID, err
			}

			return verify.UserID, ErrTriesLimitIsOver
		}

		verify.Tries += 1 // Увеличиваем количество попыток на единицу
		err = v.verifyRepo.Update(ctx, verify)
		if err != nil {
			return verify.UserID, err
		}

		return verify.UserID, ErrInvalidOTP
	}

	// Проверяем, не истек ли время жизни OTP
	if time.Now().UTC().After(verify.ExpiredAt.UTC()) {
		return verify.UserID, ErrExpiredOTP
	}

	err = v.verifyRepo.Delete(ctx, verify.ID)
	if err != nil {
		return verify.UserID, err
	}

	if !user.IsPhoneVerified {
		user.IsPhoneVerified = true
		err = v.usersRepo.Update(ctx, user.ID, user)
		if err != nil {
			return user.ID, err
		}
	}

	return user.ID, nil
}
