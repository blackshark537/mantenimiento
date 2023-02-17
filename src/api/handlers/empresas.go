package handlers

import (
	Core "github.com/blackshark537/mantenimiento/src/core"
	"github.com/blackshark537/mantenimiento/src/core/actions"
	"github.com/gofiber/fiber/v2"
)

func CreateEmpresa(c *fiber.Ctx) error { return nil }

func ReadEmpresa(c *fiber.Ctx) error {
	filter := c.Query("filter")
	action := Core.Action{
		OfType:  actions.GetAllEmpresas,
		Payload: []byte(filter),
	}
	res, err := action.Query()
	if err != nil {
		return err
	}
	return c.JSON(res)
}

func UpdateEmpresa(c *fiber.Ctx) error { return nil }
func DeleteEmpresa(c *fiber.Ctx) error { return nil }
