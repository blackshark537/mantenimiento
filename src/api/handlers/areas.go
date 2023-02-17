package handlers

import (
	Core "github.com/blackshark537/mantenimiento/src/core"
	"github.com/blackshark537/mantenimiento/src/core/actions"
	"github.com/gofiber/fiber/v2"
)

func CreateArea(c *fiber.Ctx) error { return nil }

func ReadArea(c *fiber.Ctx) error {
	filter := c.Query("filter")
	action := Core.Action{
		OfType:  actions.GetAllAreas,
		Payload: []byte(filter),
	}
	res, err := action.Query()
	if err != nil {
		return err
	}
	return c.JSON(res)
}

func UpdateArea(c *fiber.Ctx) error { return nil }
func DeleteArea(c *fiber.Ctx) error { return nil }

func CreateAreaType(c *fiber.Ctx) error { return nil }
func ReadAreaType(c *fiber.Ctx) error {
	filter := c.Query("filter")
	action := Core.Action{
		OfType:  actions.GetAllAreaType,
		Payload: []byte(filter),
	}
	res, err := action.Query()
	if err != nil {
		return err
	}
	return c.JSON(res)
}
func UpdateAreaType(c *fiber.Ctx) error { return nil }
func DeleteAreaType(c *fiber.Ctx) error { return nil }
