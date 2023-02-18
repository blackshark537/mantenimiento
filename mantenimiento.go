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
	dbConfig := postgres.DBConfig{
		Host: os.Getenv("DB_HOST"),
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		Name: os.Getenv("DB_NAME"),
		Port: os.Getenv("DB_PORT"),
	}
	core.InjectDatabase(dbConfig.NewDB())

	api := fiberApi.NewApi()
	core.InjectApi(api)

	app := Cli.New()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
