package user

import (
	"context"
	"time"

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
	panic("implement me")
}

func (v VerificationManagerImpl) ValidateOTP(ctx context.Context, channel, destination, otp string) (string, error) {
	panic("implement me")
}
