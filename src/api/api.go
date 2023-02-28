package fiberApi

import (
	"github.com/blackshark537/mantenimiento/src/api/middlewares"
	"github.com/blackshark537/mantenimiento/src/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type api struct{}

func NewApi() *api {
	return new(api)
}

func (api *api) ForRoot(port string) error {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New())

	file := middlewares.FileLoggerForRoot(app)
	defer file.Close()

	app.Use(middlewares.IsAuth)
	routes.ForRoot(app)
	return app.Listen(":" + port)
}
