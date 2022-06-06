package drivers

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Datastore interface {
	Name() string
	Connect() error
	Database() interface{}
	Ping() error
	Close(ctx context.Context) error

	WithTransaction() func(ctx context.Context, sessionFunc func(session mongo.SessionContext) error) error
}

type TxCallback func(error) error

type TxStarter interface {
	StartSession(ctx context.Context) (context.Context, TxCallback, error)
}
