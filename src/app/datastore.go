package main

import (
	"log"

	"github.com/MultiBanker/broker/src/database/drivers"
	"github.com/pkg/errors"
)

func (a *application) datastore(fn func(opts drivers.DataStoreConfig) (drivers.Datastore, error)) error {
	ds, err := fn(a.opts.Database.ToDataStore())
	if err != nil {
		return errors.Wrap(err, "DataStore Init")
	}

	if err := ds.Connect(); err != nil {
		return errors.Wrap(err, ds.Name() + "connection")
	}

	a.ds = ds
	log.Printf("[INFO] Connected to Storage %s", a.opts.Database.DSName)

	return nil
}
