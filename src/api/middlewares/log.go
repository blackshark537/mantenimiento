package middlewares

import "github.com/gofiber/fiber/v2"

func Logger(c *fiber.Ctx) error {
	return c.Next()
}
