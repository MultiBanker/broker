package main

import (
	"github.com/MultiBanker/broker/src/database/repository"
)

func (a *application) repository() error {
	repo, err := repository.NewRepository(a.ds)
	if err != nil {
		return err
	}
	a.repo = repo
	return nil
}
