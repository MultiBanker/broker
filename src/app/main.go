package main

import (
	"fmt"

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

	app := initApp(version)
	app.datastore(database.New)
	app.repository()
	app.managers()
	app.services()
	app.run()
}
