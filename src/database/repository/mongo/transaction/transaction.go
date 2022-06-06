package transaction

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Func func(ctx context.Context, sessionFunc func(session mongo.SessionContext) error) error
