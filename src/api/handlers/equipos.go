package handlers

import (
	"fmt"

	Core "github.com/blackshark537/mantenimiento/src/core"
	"github.com/blackshark537/mantenimiento/src/core/actions"
	"github.com/gofiber/fiber/v2"
)

func CreateEquipo(c *fiber.Ctx) error { return nil }
func ReadEquipo(c *fiber.Ctx) error {
	filter := c.Query("filter")
	action := Core.Action{
		OfType:  actions.GetAllEquipos,
		Payload: []byte(filter),
		Uid:     fmt.Sprintf(`%s`, c.Locals("uid")),
	}
	res, err := action.Query()
	if err != nil {
		return sendMsg(c, 500, err)
	}
	return c.JSON(res)
}
func UpdateEquipo(c *fiber.Ctx) error { return nil }
func DeleteEquipo(c *fiber.Ctx) error { return nil }

func CreateEquipoType(c *fiber.Ctx) error { return nil }

func ReadEquipoType(c *fiber.Ctx) error {
	filter := c.Query("filter")
	action := Core.Action{
		OfType:  actions.GetAllEquipoType,
		Payload: []byte(filter),
		Uid:     fmt.Sprintf(`%s`, c.Locals("uid")),
	}
	res, err := action.Query()
	if err != nil {
		return sendMsg(c, 500, err)
	}
	return c.JSON(res)
}
func UpdateEquipoType(c *fiber.Ctx) error { return nil }
func DeleteEquipoType(c *fiber.Ctx) error { return nil }
