package main

import (
	"log"
	"os"

	Cli "github.com/blackshark537/mantenimiento/src/cli"
	"github.com/blackshark537/mantenimiento/src/core"
	"github.com/blackshark537/mantenimiento/src/database/postgres"
)

func main() {
	db := postgres.DB()
	core.InjectDatabase(db)
	app := Cli.New()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
