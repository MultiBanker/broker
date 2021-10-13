package agreement

import (
	"context"
	"errors"
	"time"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	baseColl *mongo.Collection
}

func NewSpecification(signaturesColl *mongo.Collection) *Repository {
	return &Repository{
		baseColl: signaturesColl,
	}
}

func (s Repository) List(ctx context.Context, pgn selector.Paging) ([]models.Specification, int64, error) {
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
		spec := make([]models.Specification, cur.RemainingBatchLength())
		if err := cur.All(ctx, &spec); err != nil {
			return nil, 0, err
		}
		return spec, count, nil
	default:
		return nil, 0, err
	}
}

func (s Repository) Get(ctx context.Context, id string) (models.Specification, error) {
	var spec models.Specification

	filter := bson.D{
		{"_id", id},
	}

	err := s.baseColl.FindOne(ctx, filter).Decode(&spec)
	switch err {
	case mongo.ErrNoDocuments:
		return spec, drivers.ErrDoesNotExist
	case nil:
		return spec, nil
	default:
		return spec, err
	}
}

func (s Repository) GetByCode(ctx context.Context, code string) (models.Specification, error) {
	var spec models.Specification

	filter := bson.D{
		{"code", code},
	}

	err := s.baseColl.FindOne(ctx, filter).Decode(&spec)
	switch err {
	case mongo.ErrNoDocuments:
		return spec, drivers.ErrDoesNotExist
	case nil:
		return spec, nil
	default:
		return spec, err
	}
}

func (s Repository) Insert(ctx context.Context, spec models.Specification) (string, error) {
	spec.ID = primitive.NewObjectID().Hex()
	spec.CreatedAt = time.Now().UTC()

	_, err := s.baseColl.InsertOne(ctx, spec)
	return "", err
}

func (s Repository) Update(ctx context.Context, spec models.Specification) error {
	filter := bson.D{
		{"_id", spec.ID},
	}
	update := bson.D{
		{"agreement", spec.Agreement},
		{"additional_agreements", spec.AdditionalAgreements},
		{"updated_at", time.Now().UTC()},

	}

	_, err := s.baseColl.UpdateOne(ctx, filter, update)
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return err
	case errors.Is(err, nil):
		return nil
	default:
		return err
	}
}

