package signature

import (
	"context"
	"math/rand"
	"strings"
	"time"

	"github.com/MultiBanker/broker/src/database/repository"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/agree"
	"github.com/MultiBanker/broker/src/models/dto"
	"github.com/MultiBanker/broker/src/models/selector"
)

const (
	tokenLength              = 4
	hoursInDay               = 24
	PhoneTokenLen            = 4
	TriesLimit               = 2
	DefaultVerifySpamPenalty = time.Second * 60
)

type Signature interface {
	List(ctx context.Context, pgn selector.Paging) ([]models.Signature, int64, error)
	Get(ctx context.Context, id string) (models.Signature, error)
	Create(ctx context.Context, sign dto.CreateSignatureReq) (string, error)
	CheckVerification(ctx context.Context, signID, token string) (models.Signature, error)
	Update(ctx context.Context, signID string) (string, error)
}

type SignImpl struct {
	signaturesRepo    repository.AgreeSignatures
	specificationRepo repository.AgreeSpecifications
	isTesting         bool
}

func NewManager(repos repository.Repositories, isTesting bool) *SignImpl {
	return &SignImpl{
		signaturesRepo:    repos.AgreeSignatures(),
		specificationRepo: repos.AgreeSpecifications(),
		isTesting:         isTesting,
	}
}

func (s *SignImpl) List(ctx context.Context, pgn selector.Paging) ([]models.Signature, int64, error) {
	return s.signaturesRepo.List(ctx, pgn)
}

func (s *SignImpl) Get(ctx context.Context, id string) (models.Signature, error) {
	return s.signaturesRepo.Get(ctx, id)
}

func (s SignImpl) Update(ctx context.Context, signID string) (string, error) {
	token, err := GenerateRandomNumbers(PhoneTokenLen)
	if err != nil {
		return "", err
	}
	return s.signaturesRepo.RetrySign(ctx, signID, token)
}

func (s *SignImpl) Create(ctx context.Context, sign dto.CreateSignatureReq) (string, error) {
	spec, err := s.specificationRepo.GetByCode(ctx, sign.SpecCode)
	if err != nil {
		return "", err
	}

	now := time.Now()

	token, err := GenerateRandomNumbers(PhoneTokenLen)
	if err != nil {
		return "", err
	}

	signature := models.Signature{
		Agree:           spec,
		AdditionalCodes: sign.AdditionalCodes,
		Verification: models.Verification{
			Token:         token,
			Phone:         sign.Phone,
			Status:        agree.NEW.String(),
			Tries:         TriesLimit,
			NextAttemptAt: now.In(time.UTC).Add(DefaultVerifySpamPenalty),
		},
		CreatedAt: now,
		UpdatedAt: now,
	}

	return s.signaturesRepo.Insert(ctx, signature)
}

func (s *SignImpl) CheckVerification(ctx context.Context, signID, token string) (models.Signature, error) {
	// Todo: If FindOneAndUpdate is Success put status done, if not return Resource Not Found
	return s.signaturesRepo.GetVerifiedByCode(ctx, signID, token)
}

// generateToken refactor
func (s *SignImpl) generateToken() string {
	if s.isTesting {
		return strings.Repeat("1", tokenLength)
	}

	nums := []byte("0123456789")
	b := make([]byte, tokenLength)
	for i := range b {
		b[i] = nums[rand.Intn(len(b))]
	}
	return string(b)
}
