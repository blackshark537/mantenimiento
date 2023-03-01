package handlers

import (
	"fmt"

	Core "github.com/blackshark537/mantenimiento/src/core"
	actions "github.com/blackshark537/mantenimiento/src/core/actions"
	"github.com/gofiber/fiber/v2"
)

func CreateContacto(c *fiber.Ctx) error {
	action := Core.Action{
		OfType:  actions.CreateContacto,
		Payload: c.Body(),
		Uid:     fmt.Sprintf("%v", c.Locals("uid")),
	}
	err := action.Exec()
	if err != nil {
		return sendMsg(c, 500, err)
	}
	return sendMsg(c, 200, "Done!")
}

func ReadContacto(c *fiber.Ctx) error {
	filter := c.Query("filter")
	action := Core.Action{
		OfType:  actions.GetAllContactos,
		Payload: []byte(filter),
		Uid:     fmt.Sprintf("%v", c.Locals("uid")),
	}
	res, err := action.Query()
	if err != nil {
		return sendMsg(c, 500, err)
	}
	return c.JSON(res)
}

func UpdateContacto(c *fiber.Ctx) error { return nil }
func DeleteContacto(c *fiber.Ctx) error { return nil }
