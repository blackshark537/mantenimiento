package handlers

import (
	"fmt"

	Core "github.com/blackshark537/mantenimiento/src/core"
	"github.com/blackshark537/mantenimiento/src/core/actions"
	"github.com/gofiber/fiber/v2"
)

func CreateArea(c *fiber.Ctx) error {
	a := Core.Action{
		OfType:  actions.CreateArea,
		Payload: c.Body(),
		Uid:     fmt.Sprintf(`%s`, c.Locals("uid")),
	}
	err := a.Exec()
	if err != nil {
		return sendMsg(c, 500, fmt.Sprintf("[Error] Creating Area - %v\n", err))
	}

	return sendMsg(c, 200, "done!")
}

func ReadArea(c *fiber.Ctx) error {
	filter := c.Query("filter")
	action := Core.Action{
		OfType:  actions.GetAllAreas,
		Payload: []byte(filter),
		Uid:     fmt.Sprintf(`%s`, c.Locals("uid")),
	}
	res, err := action.Query()
	if err != nil {
		return sendMsg(c, 500, fmt.Sprintf("[Error] Reading Areas - %v\n", err))
	}
	return c.JSON(res)
}

func UpdateArea(c *fiber.Ctx) error { return nil }
func DeleteArea(c *fiber.Ctx) error { return nil }

func CreateAreaType(c *fiber.Ctx) error {
	a := Core.Action{
		OfType:  actions.CreateAreaType,
		Payload: c.Body(),
		Uid:     fmt.Sprintf(`%s`, c.Locals("uid")),
	}
	err := a.Exec()
	if err != nil {
		return sendMsg(c, 500, fmt.Sprintf("[Error] Creating AreaType - %v\n", err))
	}

	return sendMsg(c, 200, "done!")
}

func ReadAreaType(c *fiber.Ctx) error {
	filter := c.Query("filter")
	action := Core.Action{
		OfType:  actions.GetAllAreaType,
		Payload: []byte(filter),
		Uid:     fmt.Sprintf(`%s`, c.Locals("uid")),
	}
	res, err := action.Query()
	if err != nil {
		return sendMsg(c, 500, fmt.Sprintf("[Error] Reading AreaTypes - %v\n", err))
	}
	return c.JSON(res)
}
func UpdateAreaType(c *fiber.Ctx) error { return nil }
func DeleteAreaType(c *fiber.Ctx) error { return nil }
