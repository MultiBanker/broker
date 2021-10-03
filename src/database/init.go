package database

import (
	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/MultiBanker/broker/src/database/drivers/mongo"
)

func New(db drivers.DataStoreConfig) (drivers.Datastore, error) {
	if db.Engine == "mongo" {
		return mongo.New(db)
	}
	return nil, ErrDatastoreNotImplemented
}