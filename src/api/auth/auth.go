package auth

import "github.com/gofiber/fiber/v2"

func Authorize(c *fiber.Ctx) error {
	return c.Next()
}
