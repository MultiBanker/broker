package sequence

import (
	"context"
	"errors"

	"github.com/MultiBanker/broker/src/database/repository/mongo/transaction"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ErrEmptySequenceID = errors.New("empty sequence ID")

type sequence struct {
	Value int `bson:"value"`
}

type Repository struct {
	collection  *mongo.Collection
	transaction transaction.Func
}

func NewRepository(collection *mongo.Collection, transaction transaction.Func) Repository {
	return Repository{collection: collection, transaction: transaction}
}

// NextSequenceValue атомарно инкрементирует счетчик с именем sequenceName.
func (sr Repository) NextSequenceValue(ctx context.Context, sequenceName string) (int, error) {
	if sequenceName == "" {
		return -1, ErrEmptySequenceID
	}

	filter := bson.D{{Key: "_id", Value: sequenceName}}
	update := bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "value", Value: 1},
		}},
	}

	var seq sequence
	err := sr.collection.FindOneAndUpdate(ctx, filter, update).Decode(&seq)
	if err != nil {
		// это новый sequence
		return 0, sr.NewSequence(ctx, sequenceName)
	}

	return seq.Value, nil
}

// NewSequence создает новый инкремент с указанным идентификатором.
func (sr Repository) NewSequence(ctx context.Context, sequenceName string) error {
	if sequenceName == "" {
		return ErrEmptySequenceID
	}

	insert := bson.D{
		{Key: "_id", Value: sequenceName},
		{Key: "value", Value: 0},
	}

	if _, err := sr.collection.InsertOne(ctx, insert); err != nil {
		return err
	}

	return nil
}
