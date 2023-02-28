package routes

import (
	"github.com/blackshark537/mantenimiento/src/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func ForRoot(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to DataProd - Preventive Maintanence")
	})

	v1 := app.Group("/api/v1")

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

	/* alimentos := v1.Group("/alimentos")
	alimentos.Post("/", handlers.CreateAlimneto)
	alimentos.Get("/", handlers.ReadAlimneto)
	alimentos.Put("/:id", handlers.UpdateAlimneto)
	alimentos.Delete("/:id", handlers.DeleteAlimneto)

	ingredientes := v1.Group("/ingredientes")
	ingredientes.Post("/", handlers.CreateIngrediente)
	ingredientes.Get("/", handlers.ReadIngrediente)
	ingredientes.Put("/:id", handlers.UpdateIngrediente)
	ingredientes.Delete("/:id", handlers.DeleteIngrediente) */
}
