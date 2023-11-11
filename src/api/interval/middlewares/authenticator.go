package middlewares

import (
	"discordrm/api/pkg/langs"
	"discordrm/api/pkg/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated() fiber.Handler {
	return func(c *fiber.Ctx) error {
		schema := "Bearer "
		token := string(c.Request().Header.Peek("Authorization"))
		if len(token) > len(schema) {
			token = token[len(schema):]
		}

		data, err := utils.ValidateToken(token)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"status":  http.StatusUnauthorized,
				"message": langs.ErrUnAuthorized,
			})
		}

		c.Locals("uid", data["uid"])

		return c.Next()
	}
}
