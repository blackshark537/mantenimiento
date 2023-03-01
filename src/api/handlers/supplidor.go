package handlers

import (
	"fmt"

	Core "github.com/blackshark537/mantenimiento/src/core"
	actions "github.com/blackshark537/mantenimiento/src/core/actions"
	"github.com/gofiber/fiber/v2"
)

func CreateSuplidor(c *fiber.Ctx) error {
	action := Core.Action{
		OfType:  actions.CreateSuplidor,
		Payload: c.Body(),
		Uid:     fmt.Sprintf("%v", c.Locals("uid")),
	}
	err := action.Exec()
	if err != nil {
		return sendMsg(c, 500, err)
	}
	return sendMsg(c, 200, "Done!")
}

func ReadSuplidor(c *fiber.Ctx) error {
	filter := c.Query("filter")
	action := Core.Action{
		OfType:  actions.GetAllSuplidores,
		Payload: []byte(filter),
		Uid:     fmt.Sprintf("%v", c.Locals("uid")),
	}
	res, err := action.Query()
	if err != nil {
		return err
	}
	return c.JSON(res)
}

func UpdateSuplidor(c *fiber.Ctx) error { return nil }
func DeleteSuplidor(c *fiber.Ctx) error { return nil }
