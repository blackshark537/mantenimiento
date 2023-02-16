package fiberApi

import (
	"github.com/blackshark537/mantenimiento/src/api/routes"
	"github.com/gofiber/fiber/v2"
)

type api struct{}

func NewApi() *api {
	return new(api)
}

func (api *api) ForRoot(port string) error {
	app := fiber.New()
	routes.ForRoot(app)
	return app.Listen(":" + port)
}
