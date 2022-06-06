package order

import (
	"context"
	"errors"
	"time"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/database/repository/mongo/transaction"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	collection  *mongo.Collection
	transaction transaction.Func
}

func NewRepository(collection *mongo.Collection, transaction transaction.Func) Repository {
	return Repository{collection: collection, transaction: transaction}
}

func (or Repository) NewOrder(ctx context.Context, order *models.Order) (string, error) {
	order.CreatedAt = time.Now().UTC()
	_, err := or.collection.InsertOne(ctx, order)
	return order.ID, err
}

func (or Repository) UpdateOrder(ctx context.Context, order *models.Order) (string, error) {
	filter := bson.D{
		{"_id", order.ID},
	}
	update := bson.D{
		{"$set", bson.D{}},
		//{"reference_id", order.ReferenceID},
		//{"order_state", order.OrderState},
		//{"redirect_url", order.RedirectURL},
		//{"channel", order.Channel},
		//{"product_type", order.ProductType},
		//{"payment_method", order.PaymentMethod},
		//{"is_delivery", order.IsDelivery},
		//{"total_cost", order.Amount},
		//{"loan_length", order.LoanLength},
		//{"sales_price", order.SalesPlace},
		//{"verification_sms_code", order.VerificationSmsCode},
		//{"verification_sms_datetime", order.VerificationSmsDateTime},
		//{"customer", order.Customer},
		//{"address", order.Address},
		//{"goods", order.Goods},
		//{"bank_type", order.BankType},
		//{"updated_at", order.UpdatedAt},
	}
	_, err := or.collection.UpdateOne(ctx, filter, update)
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return "", drivers.ErrDoesNotExist
	case errors.Is(err, nil):
		return order.ID, nil
	}
	return "", err
}

//func (or Repository) UpdateOrderState(ctx context.Context, request dto.UpdateMarketOrderRequest) error {
//	filter := bson.D{
//		{"reference_id", request.ReferenceId},
//	}
//	update := bson.D{
//		{"state", request.State},
//		{"state_title", request.StateTitle},
//		{"reason", request.Reason},
//	}
//
//	err := or.collection.FindOneAndUpdate(ctx, filter, update).Err()
//	switch {
//	case errors.Is(err, mongo.ErrNoDocuments):
//		return drivers.ErrDoesNotExist
//	case errors.Is(err, nil):
//		return nil
//	}
//	return err
//}

func (or Repository) OrderByID(ctx context.Context, id string) (models.Order, error) {
	var order models.Order

	filter := bson.D{
		{"_id", id},
	}

	err := or.collection.FindOne(ctx, filter).Decode(&order)
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return order, drivers.ErrDoesNotExist
	case errors.Is(err, nil):
		return order, nil
	}
	return order, err
}

func (or Repository) Orders(ctx context.Context, paging *selector.Paging) ([]*models.Order, int64, error) {
	filter := bson.D{}

	opts := options.FindOptions{
		Skip: &paging.Skip,
		Sort: bson.D{
			{Key: paging.SortKey, Value: paging.SortVal},
		},
		Limit: &paging.Limit,
	}

	total, err := or.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	res, err := or.collection.Find(ctx, filter, &opts)
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return nil, 0, drivers.ErrDoesNotExist
	case errors.Is(err, nil):
		orders := make([]*models.Order, res.RemainingBatchLength())
		err = res.All(ctx, &orders)
		if err != nil {
			return nil, 0, err
		}

		return orders, total, nil
	}
	return nil, 0, err
}

func (or Repository) OrdersByReferenceID(ctx context.Context, referenceID string) ([]*models.Order, error) {
	filter := bson.D{
		{"reference_id", referenceID},
	}

	res, err := or.collection.Find(ctx, filter)
	switch err {
	case mongo.ErrNoDocuments:
		return nil, drivers.ErrDoesNotExist
	case nil:
		orders := make([]*models.Order, res.RemainingBatchLength())
		if err := res.All(ctx, orders); err != nil {
			return nil, err
		}
		return orders, nil
	default:
		return nil, err
	}
}
