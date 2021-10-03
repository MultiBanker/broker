package main

import (
	"log"

	"github.com/MultiBanker/broker/src/database/repository"
)

func (a *application) repository() {
	repo, err := repository.NewRepository(a.ds)
	if err != nil {
		log.Fatal("[FATAL] FUCK YOU WHERE IS DATASTORE")
	}
	a.repo = repo
}
