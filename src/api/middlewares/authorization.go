package middlewares

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/blackshark537/mantenimiento/src/api/grpc"
	"github.com/blackshark537/mantenimiento/src/auth"
)

func IsAuth(c *fiber.Ctx) error {
	conn, err := grpc.Conn()
	if err != nil {
		return err
	}
	headers := c.GetReqHeaders()
	AuthHeader := headers["Authorization"]
	if AuthHeader == "" {
		return sendMsg(c, 403, "Access Forbidden - Authorization header required!")
	}
	token := strings.Replace(AuthHeader, "Bearer ", "", 1)

	req := auth.GetAuthRequest{
		Auth: &auth.Auth{
			Token: token,
		},
	}
	client := auth.NewAuthsClient(conn)
	uid, err := client.IsAuth(c.Context(), &req)

	if uid != nil {
		c.Locals("uid", uid.IsAuth)
		return c.Next()
	}
	return sendMsg(c, 403, fmt.Sprintf("%v", err))
}

func sendMsg(c *fiber.Ctx, code int, msg string) error {
	resp := map[string]any{
		"code":    code,
		"message": msg,
	}
	fmt.Printf("[%v][Error]: %s\n", time.Now(), msg)
	c.SendStatus(code)
	return c.JSON(resp)
}
