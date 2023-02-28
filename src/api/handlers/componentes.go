package handlers

import (
	"fmt"

	Core "github.com/blackshark537/mantenimiento/src/core"
	"github.com/blackshark537/mantenimiento/src/core/actions"
	"github.com/gofiber/fiber/v2"
)

func CreateComponente(c *fiber.Ctx) error { return nil }
func ReadComponente(c *fiber.Ctx) error {
	filter := c.Query("filter")
	action := Core.Action{
		OfType:  actions.GetAllComponentes,
		Payload: []byte(filter),
		Uid:     fmt.Sprintf(`%s`, c.Locals("uid")),
	}
	res, err := action.Query()
	if err != nil {
		return err
	}
	return c.JSON(res)
}
func UpdateComponente(c *fiber.Ctx) error { return nil }
func DeleteComponente(c *fiber.Ctx) error { return nil }

func CreateComponenteType(c *fiber.Ctx) error { return nil }
func ReadComponenteType(c *fiber.Ctx) error {
	filter := c.Query("filter")
	action := Core.Action{
		OfType:  actions.GetAllComponentesType,
		Payload: []byte(filter),
		Uid:     fmt.Sprintf(`%s`, c.Locals("uid")),
	}
	res, err := action.Query()
	if err != nil {
		return err
	}
	return c.JSON(res)
}
func UpdateComponenteType(c *fiber.Ctx) error { return nil }
func DeleteComponenteType(c *fiber.Ctx) error { return nil }
