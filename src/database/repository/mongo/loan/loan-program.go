package loan

import (
	"context"
	"time"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProgramRepository struct {
	coll *mongo.Collection
}

func NewProgramRepository(coll *mongo.Collection) *ProgramRepository {
	return &ProgramRepository{coll: coll}
}

func (l ProgramRepository) LoanProgram(ctx context.Context, code string) (models.LoanProgram, error) {
	var loanPr models.LoanProgram

	filter := bson.D{
		{"code", code},
		{"is_enabled", true},
	}

	err := l.coll.FindOne(ctx, filter).Decode(&loanPr)
	switch err {
	case mongo.ErrNoDocuments:
		return loanPr, drivers.ErrDoesNotExist
	case nil:
		return loanPr, nil
	default:
		return loanPr, err
	}
}

func (l ProgramRepository) LoanPrograms(ctx context.Context, search selector.Paging) ([]models.LoanProgram, int64, error) {
	filter := bson.D{}

	opts := &options.FindOptions{
		Skip: &search.Skip,
		Sort: bson.D{
			{Key: search.SortKey, Value: search.SortVal},
		},
		Limit: &search.Limit,
	}
	total, err := l.coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	cur, err := l.coll.Find(ctx, filter, opts)
	switch err {
	case mongo.ErrNoDocuments:
		return nil, 0, drivers.ErrDoesNotExist
	case nil:
		progs := make([]models.LoanProgram, cur.RemainingBatchLength())
		if err := cur.All(ctx, &progs); err != nil {
			return nil, 0, err
		}
		return progs, total, nil
	default:
		return nil, 0, err
	}
}

func (l ProgramRepository) CreateLoanProgram(ctx context.Context, program models.LoanProgram) (string, error) {
	program.ID = primitive.NewObjectID().Hex()
	program.IsEnabled = true
	program.CreatedAt = time.Now().UTC()
	_, err := l.coll.InsertOne(ctx, program)
	if err != nil {
		return "", err
	}
	return program.ID, nil
}

func (l ProgramRepository) UpdateLoanProgram(ctx context.Context, code string, program models.LoanProgram) error {
	filter := bson.D{
		{"code", code},
	}
	update := bson.D{
		{"$set", bson.D{
			{"name", program.Name},
			{"note", program.Note},
			{"type", program.Type},
			{"partner_code", program.PartnerCode},
			{"is_enabled", program.IsEnabled},
			{"term", program.Term},
			{"rate", program.Rate},
			{"min_amount", program.MinAmount},
			{"max_amount", program.MaxAmount},
			{"updated_at", time.Now().UTC()},
		}},
	}
	_, err := l.coll.UpdateOne(ctx, filter, update)
	switch err {
	case mongo.ErrNoDocuments:
		return drivers.ErrDoesNotExist
	case nil:
		return nil
	default:
		return err
	}
}
