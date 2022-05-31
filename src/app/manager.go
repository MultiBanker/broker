package main

import (
	"github.com/MultiBanker/broker/src/manager"
)

func (a *application) managers() {
	a.man = manager.NewWrapper(a.ds, a.repo, a.opts, a.metric)
}