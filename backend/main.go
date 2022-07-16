
package main

import (
	"github.com/RobertOchmanek/ebiznes_go/database"
	"github.com/RobertOchmanek/ebiznes_go/routing"
)

func main() {

	//Initialize DB connection
	database.Init()

	//Initialize and start server
	e := routing.Init()
	e.Logger.Fatal(e.Start(":8080"))
}