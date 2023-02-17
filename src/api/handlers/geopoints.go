package handlers

import (
	Core "github.com/blackshark537/mantenimiento/src/core"
	"github.com/blackshark537/mantenimiento/src/core/actions"
	"github.com/gofiber/fiber/v2"
)

func CreateGeopoint(c *fiber.Ctx) error { return nil }
func ReadGeopoint(c *fiber.Ctx) error {
	filter := c.Query("filter")
	action := Core.Action{
		OfType:  actions.GetAllGeo,
		Payload: []byte(filter),
	}
	res, err := action.Query()
	if err != nil {
		return err
	}
	return c.JSON(res)
}
func UpdateGeopoint(c *fiber.Ctx) error { return nil }
func DeleteGeopoint(c *fiber.Ctx) error { return nil }
