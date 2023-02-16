package main

import (
	"log"
	"os"

	fiberApi "github.com/blackshark537/mantenimiento/src/api"
	Cli "github.com/blackshark537/mantenimiento/src/cli"
	"github.com/blackshark537/mantenimiento/src/core"
	"github.com/blackshark537/mantenimiento/src/database/postgres"
)

func main() {
	db := postgres.NewDB()
	core.InjectDatabase(db)

	api := fiberApi.NewApi()
	core.InjectApi(api)

	app := Cli.New()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
