package routes

import (
	"fmt"

	"github.com/blackshark537/mantenimiento/src/api/handlers"
	"github.com/blackshark537/mantenimiento/src/config"

	"github.com/gofiber/fiber/v2"
)

func ForRoot(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("%s - %s", config.AppName, config.AppDesc))
	})

	v1 := app.Group("/api/v1/mantenimiento")

	areas := v1.Group("/areas")
	areas.Post("/", handlers.CreateArea)
	areas.Get("/", handlers.ReadArea)
	areas.Put("/:id", handlers.UpdateArea)
	areas.Delete("/:id", handlers.DeleteArea)

	areasType := v1.Group("/areas/type")
	areasType.Post("/", handlers.CreateAreaType)
	areasType.Get("/", handlers.ReadAreaType)
	areasType.Put("/:id", handlers.UpdateAreaType)
	areasType.Delete("/:id", handlers.DeleteAreaType)

	equipos := v1.Group("/equipos")
	equipos.Post("/", handlers.CreateEquipo)
	equipos.Get("/", handlers.ReadEquipo)
	equipos.Put("/:id", handlers.UpdateEquipo)
	equipos.Delete("/:id", handlers.DeleteEquipo)

	equiposType := v1.Group("/equipos/type")
	equiposType.Post("/", handlers.CreateEquipoType)
	equiposType.Get("/", handlers.ReadEquipoType)
	equiposType.Put("/:id", handlers.UpdateEquipoType)
	equiposType.Delete("/:id", handlers.DeleteEquipoType)

	componentes := v1.Group("/componentes")
	componentes.Post("/", handlers.CreateComponente)
	componentes.Get("/", handlers.ReadComponente)
	componentes.Put("/:id", handlers.UpdateComponente)
	componentes.Delete("/:id", handlers.DeleteComponente)

	componentesType := v1.Group("/componentes/type")
	componentesType.Post("/", handlers.CreateComponenteType)
	componentesType.Get("/", handlers.ReadComponenteType)
	componentesType.Put("/:id", handlers.UpdateComponenteType)
	componentesType.Delete("/:id", handlers.DeleteComponenteType)

	suplidores := v1.Group("suplidores")
	suplidores.Post("/", handlers.CreateSuplidor)
	suplidores.Get("/", handlers.ReadSuplidor)
	suplidores.Put("/:id", handlers.UpdateSuplidor)
	suplidores.Delete("/:id", handlers.DeleteSuplidor)

	contactos := v1.Group("contactos")
	contactos.Post("/", handlers.CreateContacto)
	contactos.Get("/", handlers.ReadContacto)
	contactos.Put("/:id", handlers.UpdateContacto)
	contactos.Delete("/:id", handlers.DeleteContacto)
}
