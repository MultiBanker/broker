package main

import (
	"github.com/MultiBanker/broker/src/manager"
)

func (a *application) managers() {
	a.clients()
	a.man = manager.NewAbstract(a.ds, a.repo, a.opts)
}

func (a *application) clients() {

}
