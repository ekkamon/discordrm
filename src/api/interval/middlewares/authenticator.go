package middlewares

import (
	"discordrm/api/pkg/databases"
	"discordrm/api/pkg/langs"
	"discordrm/api/pkg/utils"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func IsAuthenticated(db *databases.Conn) fiber.Handler {
	return func(c *fiber.Ctx) error {
		schema := "Bearer "
		reqToken := string(c.Request().Header.Peek("Authorization"))
		if len(reqToken) > len(schema) {
			reqToken = reqToken[len(schema):]
		}

		data, err := utils.ValidateToken(reqToken)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"status":  http.StatusUnauthorized,
				"message": langs.ErrUnAuthorized,
			})
		}
		uid := int(data["uid"].(float64))

		prefix := fmt.Sprintf("auth:%d", uid)
		token, err := db.Redis.Client.Get(db.Redis.Ctx, prefix).Result()
		if err != nil {
			status := http.StatusUnauthorized
			message := langs.ErrUnAuthorized

			if err != redis.Nil {
				status = http.StatusInternalServerError
				message = langs.ErrInternalServer

				logrus.WithFields(logrus.Fields{
					"file": "middlewares/authenticator.go",
					"func": "db.Redis.Client.Get",
				}).Error(err)
			}

			return c.Status(status).JSON(fiber.Map{
				"status":  status,
				"message": message,
			})
		}

		if token != reqToken {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"status":  http.StatusUnauthorized,
				"message": langs.ErrUnAuthorized,
			})
		}

		// past uid value
		c.Locals("uid", uid)

		return c.Next()
	}
}
