package handlers

import (
	Core "github.com/blackshark537/mantenimiento/src/core"
	"github.com/blackshark537/mantenimiento/src/core/actions"
	"github.com/gofiber/fiber/v2"
)

func CreateContato(c *fiber.Ctx) error { return nil }
func ReadContato(c *fiber.Ctx) error {
	filter := c.Query("filter")
	action := Core.Action{
		OfType:  actions.GetAllContactos,
		Payload: []byte(filter),
	}
	res, err := action.Query()
	if err != nil {
		return err
	}
	return c.JSON(res)
}
func UpdateContato(c *fiber.Ctx) error { return nil }
func DeleteContato(c *fiber.Ctx) error { return nil }
