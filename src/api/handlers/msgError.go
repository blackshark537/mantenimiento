package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func sendMsg(c *fiber.Ctx, code int, msg interface{}) error {
	resp := map[string]any{
		"code":    code,
		"message": msg,
	}
	fmt.Printf("[%v][Error]: %s\n", time.Now(), msg)
	c.SendStatus(code)
	return c.JSON(resp)
}
