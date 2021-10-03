package drivers

import (
	"context"
)

type Datastore interface {
	Name() string
	Connect() error
	Database() interface{}
	Ping() error
	Close(ctx context.Context) error
}
