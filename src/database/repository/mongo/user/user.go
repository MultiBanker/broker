package user

import (
	"context"
	"log"
	"time"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/database/repository/mongo/transaction"
	"github.com/MultiBanker/broker/src/models"
	"github.com/MultiBanker/broker/src/models/selector"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UsersRepositoryImpl struct {
	coll        *mongo.Collection
	transaction transaction.Func
}

func NewUsersRepositoryImpl(coll *mongo.Collection, transaction transaction.Func) *UsersRepositoryImpl {
	return &UsersRepositoryImpl{
		coll:        coll,
		transaction: transaction,
	}
}

func (u *UsersRepositoryImpl) Create(ctx context.Context, user models.User) (string, error) {
	result, err := u.coll.InsertOne(ctx, user)
	if mongo.IsDuplicateKeyError(err) {
		return "", drivers.ErrAlreadyExists
	}
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), err
}

func (u *UsersRepositoryImpl) Get(ctx context.Context, query *selector.SearchQuery) ([]models.User, error) {
	users := make([]models.User, 0, query.Pagination.Limit)
	opts := new(options.FindOptions)

	opts.SetLimit(query.Pagination.Limit)
	opts.SetSkip(query.Pagination.Page * query.Pagination.Limit)

	if query.HasSorting() {
		opts.SetSort(bson.D{
			{Key: query.Sorting.Key, Value: query.Sorting.Direction},
			{Key: "_id", Value: query.Sorting.Direction},
		})
	} else {
		opts.SetSort(bson.D{
			{Key: "last_name", Value: 1},
			{Key: "_id", Value: 1},
		})
	}

	cur, err := u.coll.Find(ctx, u.searchFilters(query), opts)
	if err != nil {
		return users, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var user models.User
		err = cur.Decode(&user)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *UsersRepositoryImpl) GetOrCreateUserByPhone(ctx context.Context, phone string) (string, error) {

	user := struct {
		ID string `bson:"_id"`
	}{}

	var err error

	filter := bson.M{"phone": phone}
	opts := options.FindOne().SetProjection(bson.M{"_id": 1})

	err = u.coll.FindOne(ctx, filter, opts).Decode(&user)
	switch err {
	case mongo.ErrNoDocuments:
		result, err := u.coll.InsertOne(ctx, &models.User{
			Phone:     phone,
			IsEnabled: true,
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		})
		if err != nil {
			return "", err
		}
		return result.InsertedID.(primitive.ObjectID).Hex(), nil
	case nil:
		return user.ID, nil
	default:
		return "", err
	}
}

func (u *UsersRepositoryImpl) GetByID(ctx context.Context, userID string) (models.User, error) {
	var user models.User
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return user, drivers.ErrBadID
	}

	filter := bson.D{{Key: "_id", Value: oid}}
	err = u.coll.FindOne(ctx, filter).Decode(&user)
	switch err {
	case mongo.ErrNoDocuments:
		return user, drivers.ErrDoesNotExist
	default:
		return user, err
	}
}

func (u *UsersRepositoryImpl) GetByIDs(ctx context.Context, userIDs []string) ([]models.User, error) {
	const op = "UsersRepositoryImpl.GetByIDs"

	if len(userIDs) == 0 {
		return []models.User{}, nil
	}

	oids := make([]primitive.ObjectID, 0, len(userIDs))
	for _, userID := range userIDs {
		oid, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			log.Printf("[ERROR] %s:%s", op, userID)
			continue
		}

		oids = append(oids, oid)
	}

	filter := bson.D{{Key: "_id", Value: bson.D{
		{Key: "$in", Value: oids},
	}}}

	cur, err := u.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	res := make([]models.User, 0, cur.RemainingBatchLength())
	if err = cur.All(ctx, &res); err != nil {
		return nil, err
	}

	return res, err
}

func (u *UsersRepositoryImpl) Update(ctx context.Context, userID string, user models.User) error {
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return drivers.ErrBadID
	}

	filter := bson.D{{Key: "_id", Value: oid}}
	update := bson.D{
		{
			Key: "$set", Value: bson.D{
				{Key: "first_name", Value: user.FirstName},
				{Key: "last_name", Value: user.LastName},
				{Key: "patronymic", Value: user.Patronymic},
				{Key: "phone", Value: user.Phone},
				{Key: "email", Value: user.Email},
				{Key: "password", Value: user.Password},
				{Key: "created_at", Value: user.CreatedAt},
				{Key: "updated_at", Value: time.Now().UTC()},
				{Key: "is_enabled", Value: user.IsEnabled},
				{Key: "is_phone_verified", Value: user.IsPhoneVerified},
				{Key: "is_email_verified", Value: user.IsEmailVerified},
			},
		},
	}

	result, err := u.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return drivers.ErrDoesNotExist
	}

	return nil
}

func (u *UsersRepositoryImpl) UpdatePassword(ctx context.Context, userID, password string) error {
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return drivers.ErrBadID
	}

	filter := bson.D{{Key: "_id", Value: oid}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "password", Value: password},
	}}}

	result, err := u.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return drivers.ErrDoesNotExist
	}
	return nil
}

func (u *UsersRepositoryImpl) UpdatePhone(ctx context.Context, userID, phone string) error {
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return drivers.ErrBadID
	}

	filter := bson.D{{Key: "_id", Value: oid}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "phone", Value: phone},
	}}}

	result, err := u.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return drivers.ErrDoesNotExist
	}
	return nil
}

func (u *UsersRepositoryImpl) Delete(ctx context.Context, userID string) error {
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return drivers.ErrBadID
	}

	filter := bson.D{{Key: "_id", Value: oid}}

	_, err = u.coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (u *UsersRepositoryImpl) Count(ctx context.Context, query *selector.SearchQuery) (int64, error) {
	result, err := u.coll.CountDocuments(ctx, u.searchFilters(query))
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (u *UsersRepositoryImpl) searchFilters(query *selector.SearchQuery) bson.D {
	searchFilter := make(bson.D, 0)
	if query == nil {
		return searchFilter
	}

	if query.UserID != "" {
		searchFilter = append(searchFilter, bson.E{Key: "_id", Value: query.UserID})
	}
	if query.FirstName != "" {
		searchFilter = append(searchFilter, bson.E{Key: "first_name", Value: primitive.Regex{
			Pattern: query.FirstName,
			Options: "i",
		}})
	}
	if query.LastName != "" {
		searchFilter = append(searchFilter, bson.E{Key: "last_name", Value: primitive.Regex{
			Pattern: query.LastName,
			Options: "i",
		}})
	}
	if query.Patronymic != "" {
		searchFilter = append(searchFilter, bson.E{Key: "patronymic", Value: primitive.Regex{
			Pattern: query.Patronymic,
			Options: "i",
		}})
	}
	if query.Phone != "" {
		searchFilter = append(searchFilter, bson.E{Key: "phone", Value: query.Phone})
	}
	if query.Email != "" {
		searchFilter = append(searchFilter, bson.E{Key: "email", Value: query.Email})
	}
	if query.IsEmailVerified != nil {
		searchFilter = append(searchFilter, bson.E{Key: "is_email_verified", Value: query.IsEmailVerified})
	}
	if query.IsPhoneVerified != nil {
		searchFilter = append(searchFilter, bson.E{Key: "is_phone_verified", Value: query.IsPhoneVerified})
	}
	return searchFilter
}
