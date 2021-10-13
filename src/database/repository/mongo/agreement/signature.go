package agreement

import (
	"context"
	"errors"
	"time"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/agree"
	"github.com/MultiBanker/broker/src/models/selector"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignatureRepository struct {
	baseColl *mongo.Collection
}

func NewSignature(signaturesColl *mongo.Collection) *SignatureRepository {
	return &SignatureRepository{
		baseColl: signaturesColl,
	}
}

func (s SignatureRepository) List(ctx context.Context, pgn selector.Paging) ([]models.Signature, int64, error) {
	filter := bson.D{}

	opts := &options.FindOptions{
		Skip: &pgn.Skip,
		Sort: bson.D{
			{Key: pgn.SortKey, Value: pgn.SortVal},
		},
		Limit: &pgn.Limit,
	}

	count, err := s.baseColl.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	cur, err := s.baseColl.Find(ctx, filter, opts)
	switch err {
	case mongo.ErrNoDocuments:
		return nil, 0, drivers.ErrDoesNotExist
	case nil:
		spec := make([]models.Signature, cur.RemainingBatchLength())
		if err := cur.All(ctx, &spec); err != nil {
			return nil, 0, err
		}
		return spec, count, nil
	default:
		return nil, 0, err
	}
}

func (s SignatureRepository) Get(ctx context.Context, id string) (models.Signature, error) {
	var sign models.Signature

	filter := bson.D{
		{"_id", id},
	}

	err := s.baseColl.FindOne(ctx, filter).Decode(&sign)
	switch err {
	case mongo.ErrNoDocuments:
		return sign, drivers.ErrDoesNotExist
	case nil:
		return sign, nil
	default:
		return sign, err
	}
}

func (s SignatureRepository) Insert(ctx context.Context, signature models.Signature) (string, error) {
	signature.ID = primitive.NewObjectID().Hex()
	signature.CreatedAt = time.Now().UTC()

	_, err := s.baseColl.InsertOne(ctx, signature)
	if err != nil {
		return "", err
	}
	return "", nil
}

func (s SignatureRepository) UpdateVerification(ctx context.Context, signature models.Signature) error {
	filter := bson.D{
		{"_id", signature.ID},
	}
	update := bson.D{
		{"verification_id", signature.VerificationID},
		{"verification", signature.Verification},
		{"agree", signature.Agree},
		{"user_id", signature.UserID},
		{"additional_codes", signature.AdditionalCodes},
		{"updated_at", time.Now().UTC()},
	}

	_, err := s.baseColl.UpdateOne(ctx, filter, update)
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return drivers.ErrDoesNotExist
	case errors.Is(err, nil):
		return nil
	}
	return err
}

func (s SignatureRepository) GetVerifiedByCode(ctx context.Context, userID, token string) (models.Signature, error) {
	var sign models.Signature

	filter := bson.D{
		{"user_id", userID},
		{"verification.token", token},
	}

	update := bson.D{
		{"verification.status", agree.VERIFIED.String()},
	}

	err := s.baseColl.FindOneAndUpdate(ctx, filter, update).Decode(&sign)
	switch err {
	case mongo.ErrNoDocuments:
		return sign, drivers.ErrDoesNotExist
	case nil:
		return sign, nil
	default:
		return sign, err
	}
}

func (s SignatureRepository) NewVerificationStatus(ctx context.Context) (models.Signature, error) {
	var sign models.Signature

	filter := bson.D{
		{"verification.status", agree.NEW.String()},
	}
	update := bson.D{
		{"verification.status", agree.ONCHECK.String()},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err := s.baseColl.FindOneAndUpdate(ctx, filter, update, opts).Decode(&sign)
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return sign, drivers.ErrDoesNotExist
	case errors.Is(err, nil):
		return sign, nil
	}
	return sign, err
}

func (s SignatureRepository) RetrySign(ctx context.Context, signID, token string) (string, error) {
	filter := bson.D{
		{"_id", signID},
		{"verification.token", token},
	}
	update := bson.D{
		{"verification.status", agree.NEW.String()},
	}
	err := s.baseColl.FindOneAndUpdate(ctx, filter, update).Err()
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return "", drivers.ErrDoesNotExist
	case errors.Is(err, nil):
		return signID, nil
	}
	return "", err
}
