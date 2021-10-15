package main

import (
	"fmt"
	"os"

	"github.com/MultiBanker/broker/src/database"
)

var version = ""

// @title Broker API
// @version 1.0
// @description Broker banker requester
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @host api.test.airba.dev
// @BasePath /broker/api/v1/broker
// @schemes https
// @query.collection.format multi
func main() {
	fmt.Printf("version %s", version)

	app, err := initApp(version)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err = app.datastore(database.New); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err = app.repository(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	app.managers()
	app.services()
	app.run()
}
