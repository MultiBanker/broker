package user

import (
	"context"
	"errors"
	"time"

	"github.com/MultiBanker/broker/pkg/auth"
	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
)

const (
	RecoveryOtpTTL          = 2 * time.Minute  // Время жизни ОТП на восстановление
	RecoveryLongSpamPenalty = 60 * time.Minute // Время длительного штрафа при попытках спама
	RecoverySpamPenalty     = 60 * time.Second // Время короткого штрафа при попытках спама
	RecoveryTriesLimit      = 3                // Количество попыток при вводе СМС
)

type RecoveryManager interface {
	SendOTP(ctx context.Context, channel, destination string) error
	ValidateOTP(ctx context.Context, channel, phone, otp, password string) error
}

type RecoveryManagerImpl struct {
	isTesting    bool
	usersRepo    repository.UsersRepository
	recoveryRepo repository.RecoveryRepository
}

func NewRecoveryManagerImpl(
	isTesting bool,
	recoveryRepo repository.RecoveryRepository,
	usersRepo repository.UsersRepository) *RecoveryManagerImpl {
	return &RecoveryManagerImpl{
		isTesting:    isTesting,
		usersRepo:    usersRepo,
		recoveryRepo: recoveryRepo,
	}
}

func (man *RecoveryManagerImpl) SendOTP(ctx context.Context, channel, destination string) error {
	user, err := man.usersRepo.Get(ctx, &selector.SearchQuery{})
	if err != nil {
		return err
	}

	recovery, err := man.recoveryRepo.GetByChannel(ctx, channel, destination)
	switch {
	case errors.Is(err, drivers.ErrDoesNotExist):

		newRecovery := models.Recovery{
			UserID:        user[0].ID,
			Channel:       channel,
			Destination:   destination,
			Status:        "NEW",
			Send:          false,
			Count:         1,
			Tries:         0,
			ExpiredAt:     time.Now().UTC().Add(RecoveryOtpTTL),
			NextAttemptAt: time.Now().UTC().Add(RecoverySpamPenalty),
		}

		return man.recoveryRepo.Create(ctx, newRecovery)
	case err == nil:
		if time.Now().UTC().Before(recovery.NextAttemptAt.UTC()) {
			return TooManyRequestsError{
				NextAttemptAt: recovery.NextAttemptAt.UTC().Format(time.RFC3339),
			}
		}

		recovery.Status = "NEW"
		recovery.Send = false
		recovery.Count += 1
		recovery.ExpiredAt = time.Now().Add(RecoveryOtpTTL)
		recovery.NextAttemptAt = time.Now().Add(RecoverySpamPenalty)

		return man.recoveryRepo.Update(ctx, recovery)
	}

	return err
}

func (man *RecoveryManagerImpl) ValidateOTP(ctx context.Context, channel, destination, otp, password string) error {
	recovery, err := man.recoveryRepo.GetByChannel(ctx, channel, destination)
	if err != nil {
		return err
	}

	// Проверка на корректность введённого пользователем OTP
	if recovery.OTP != otp {
		// Проверка на количество введённого пользователем OTP
		if recovery.Tries >= RecoveryTriesLimit {

			recovery.OTP = ""
			recovery.Tries = 0
			recovery.Count = 0
			recovery.NextAttemptAt = time.Now().Add(RecoveryLongSpamPenalty) // Увеличиваем время штрафа до 60 мин

			err = man.recoveryRepo.Update(ctx, recovery)
			if err != nil {
				return err
			}

			return ErrTriesLimitIsOver
		}

		recovery.Tries += 1 // Увеличиваем количество попыток на единицу
		err = man.recoveryRepo.Update(ctx, recovery)
		if err != nil {
			return err
		}

		return ErrInvalidOTP
	}

	// Проверяем, не истек ли время жизни OTP
	if time.Now().UTC().After(recovery.ExpiredAt.UTC()) {
		return ErrExpiredOTP
	}

	hashedPass, err := auth.HashPassword(password)
	if err != nil {
		return ErrAuthorization
	}
	err = man.usersRepo.UpdatePassword(ctx, recovery.UserID, string(hashedPass))
	if err != nil {
		return err
	}

	return man.recoveryRepo.Delete(ctx, recovery.ID)
}
