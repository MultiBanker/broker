package main

import (
	"log"

	"github.com/MultiBanker/broker/src/database/drivers"
)

func (a *application) datastore(fn func(opts drivers.DataStoreConfig) (drivers.Datastore, error)) {
	ds, err := fn(a.opts.Database.ToDataStore())
	if err != nil {
		log.Fatal("[ERROR] cannot create datastore")
	}

	if err := ds.Connect(); err != nil {
		log.Fatal("[ERROR] cannot connect to datastore")
	}

	a.ds = ds
	log.Printf("[INFO] Connected to Storage %s", a.opts.Database.DSName)
}
