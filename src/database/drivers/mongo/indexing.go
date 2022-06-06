package mongo

import (
	"context"

	"github.com/MultiBanker/broker/src/database/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m Mongo) LoanIndexes(ctx context.Context) error {
	col := m.DB.Collection(repository.LoanPrograms)

	var err error
	var exists bool

	indexes := make([]mongo.IndexModel, 0)

	exists, err = m.indexExistsByName(ctx, col, "code_idx")
	if err != nil {
		return err
	}
	if !exists {
		indexes = append(indexes,
			mongo.IndexModel{
				Keys: bson.D{
					bson.E{Key: "code", Value: 1},
					bson.E{Key: "is_enabled", Value: 1},
				},
				Options: options.Index().SetName("code_idx"),
			})
	}
	opts := options.CreateIndexes().SetMaxTime(m.ensureIdxTimeout)

	_, err = col.Indexes().CreateMany(ctx, indexes, opts)
	return err
}

func (m Mongo) MarketIndexes(ctx context.Context) error {
	col := m.DB.Collection(repository.Market)

	var err error
	var exists bool

	indexes := make([]mongo.IndexModel, 0)

	exists, err = m.indexExistsByName(ctx, col, "market_code_idx")
	if err != nil {
		return err
	}
	if !exists {
		indexes = append(indexes,
			mongo.IndexModel{
				Keys: bson.D{
					bson.E{Key: "code", Value: 1},
				},
				Options: options.Index().SetName("market_code_idx"),
			})
	}

	exists, err = m.indexExistsByName(ctx, col, "market_username_idx")
	if err != nil {
		return err
	}
	if !exists {
		indexes = append(indexes,
			mongo.IndexModel{
				Keys: bson.D{
					bson.E{Key: "username", Value: 1},
				},
				Options: options.Index().SetName("market_username_idx"),
			})
	}

	opts := options.CreateIndexes().SetMaxTime(m.ensureIdxTimeout)

	_, err = col.Indexes().CreateMany(ctx, indexes, opts)
	return err
}

func (m Mongo) OrderIndexes(ctx context.Context) error {
	col := m.DB.Collection(repository.Order)

	var err error
	var exists bool

	indexes := make([]mongo.IndexModel, 0)

	exists, err = m.indexExistsByName(ctx, col, "reference_id_idx")
	if err != nil {
		return err
	}
	if !exists {
		indexes = append(indexes,
			mongo.IndexModel{
				Keys: bson.D{
					bson.E{Key: "reference_id", Value: 1},
				},
				Options: options.Index().SetName("reference_id_idx"),
			})
	}

	opts := options.CreateIndexes().SetMaxTime(m.ensureIdxTimeout)

	_, err = col.Indexes().CreateMany(ctx, indexes, opts)
	return err
}

func (m Mongo) PartnerOrderIndexes(ctx context.Context) error {
	col := m.DB.Collection(repository.PartnerOrder)

	var err error
	var exists bool

	col.Indexes().DropAll(ctx)
	indexes := make([]mongo.IndexModel, 0)

	exists, err = m.indexExistsByName(ctx, col, "update_order_idx")
	if err != nil {
		return err
	}
	if !exists {
		indexes = append(indexes,
			mongo.IndexModel{
				Keys: bson.D{
					{"partner_code", 1},
					{"reference_id", 1},
				},
				Options: options.Index().SetName("update_order_idx"),
			})
	}

	exists, err = m.indexExistsByName(ctx, col, "status_to_idx")
	if err != nil {
		return err
	}
	if !exists {
		indexes = append(indexes,
			mongo.IndexModel{
				Keys: bson.D{
					{"state", 1},
					{"created_at", 1},
				},
				Options: options.Index().SetName("status_to_idx"),
			})
	}

	exists, err = m.indexExistsByName(ctx, col, "order_refer_market_idx")
	if err != nil {
		return err
	}
	if !exists {
		indexes = append(indexes,
			mongo.IndexModel{
				Keys: bson.D{
					{"market_code", 1},
					{"reference_id", 1},
				},
				Options: options.Index().SetName("order_refer_market_idx"),
			})
	}

	exists, err = m.indexExistsByName(ctx, col, "order_refer_partner_idx")
	if err != nil {
		return err
	}
	if !exists {
		indexes = append(indexes,
			mongo.IndexModel{
				Keys: bson.D{
					{"reference_id", 1},
					{"partner_code", 1},
				},
				Options: options.Index().SetName("order_refer_partner_idx"),
			})
	}


	opts := options.CreateIndexes().SetMaxTime(m.ensureIdxTimeout)

	_, err = col.Indexes().CreateMany(ctx, indexes, opts)
	return err
}


func (m Mongo) PartnerIndexes(ctx context.Context) error {
	col := m.DB.Collection(repository.Partner)

	var err error
	var exists bool


	indexes := make([]mongo.IndexModel, 0)

	exists, err = m.indexExistsByName(ctx, col, "partner_code_idx")
	if err != nil {
		return err
	}
	if !exists {
		indexes = append(indexes,
			mongo.IndexModel{
				Keys: bson.D{
					bson.E{Key: "code", Value: 1},
				},
				Options: options.Index().SetName("partner_code_idx"),
			})
	}

	exists, err = m.indexExistsByName(ctx, col, "partner_username_idx")
	if err != nil {
		return err
	}
	if !exists {
		indexes = append(indexes,
			mongo.IndexModel{
				Keys: bson.D{
					bson.E{Key: "username", Value: 1},
				},
				Options: options.Index().SetName("partner_username_idx"),
			})
	}


	opts := options.CreateIndexes().SetMaxTime(m.ensureIdxTimeout)

	_, err = col.Indexes().CreateMany(ctx, indexes, opts)
	return err
}

func (m Mongo) SequenceIndexes(ctx context.Context) error {
	col := m.DB.Collection(repository.Sequence)

	var err error
	var exists bool

	indexes := make([]mongo.IndexModel, 0)

	exists, err = m.indexExistsByName(ctx, col, "val_seq_idx")
	if err != nil {
		return err
	}
	if !exists {
		indexes = append(indexes,
			mongo.IndexModel{
				Keys: bson.D{
					bson.E{Key: "value", Value: 1},
				},
				Options: options.Index().SetName("val_seq_idx"),
			})
	}


	opts := options.CreateIndexes().SetMaxTime(m.ensureIdxTimeout)

	_, err = col.Indexes().CreateMany(ctx, indexes, opts)
	return err
}