package main

import (
	"github.com/MultiBanker/broker/src/manager"
)

func (a *application) managers() {
	a.clients()
	a.man = manager.NewAbstract(a.ds, a.repo, a.opts, a.metric)
}

func (a *application) clients() {

}
