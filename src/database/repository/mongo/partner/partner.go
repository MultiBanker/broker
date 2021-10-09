package partner

import (
	"context"
	"errors"
	"time"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/models/selector"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/MultiBanker/broker/src/models"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) Repository {
	return Repository{collection: collection}
}

func (p Repository) NewPartner(ctx context.Context, partner *models.Partner) (string, error) {
	partner.CreatedAt = time.Now().UTC()
	_, err := p.collection.InsertOne(ctx, partner)
	return partner.ID, err
}

func (p Repository) UpdatePartner(ctx context.Context, partner *models.Partner) (string, error) {
	filter := bson.D{
		{"_id", partner.ID},
	}
	update := bson.D{
		{"$set", bson.D{
			{"company_name", partner.CompanyName},
			{"username", partner.Username},
			{"password", partner.Password},
			{"url", partner.URL},
			{"bin", partner.BIN},
			{"commission", partner.Commission},
			{"logo_url", partner.LogoURL},
			{"contact", partner.Contact},
			{"enabled", partner.Enabled},
			{"updated_at", time.Now().UTC()},
		}},
	}
	_, err := p.collection.UpdateOne(ctx, filter, update)
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return "", drivers.ErrDoesNotExist
	case errors.Is(err, nil):
		return partner.ID, nil
	}
	return "", err
}

func (p Repository) PartnerByID(ctx context.Context, id string) (models.Partner, error) {
	var partner models.Partner
	filter := bson.D{
		{"_id", id},
	}
	err := p.collection.FindOne(ctx, filter).Decode(&partner)
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return partner, drivers.ErrDoesNotExist
	case errors.Is(err, nil):
		return partner, nil
	}

	return partner, err
}

func (p Repository) PartnerByUsername(ctx context.Context, username string) (models.Partner, error) {
	var partner models.Partner
	filter := bson.D{
		{"username", username},
	}
	err := p.collection.FindOne(ctx, filter).Decode(&partner)
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return partner, drivers.ErrDoesNotExist
	case errors.Is(err, nil):
		return partner, nil
	}

	return partner, err
}

func (p Repository) Partners(ctx context.Context, paging *selector.Paging) ([]models.Partner, int64, error) {
	filter := bson.D{}

	opts := options.FindOptions{
		Skip: &paging.Skip,
		Sort: bson.D{
			{Key: paging.SortKey, Value: paging.SortVal},
		},
		Limit: &paging.Limit,
	}

	total, err := p.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	res, err := p.collection.Find(ctx, filter, &opts)
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return nil, 0, drivers.ErrDoesNotExist
	case errors.Is(err, nil):
		partners := make([]models.Partner, res.RemainingBatchLength())
		err = res.All(ctx, &partners)
		if err != nil {
			return nil, 0, err
		}

		return partners, total, nil
	}
	return nil, 0, err
}
