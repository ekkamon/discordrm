package user

import (
	"discordrm/api/interval/entities"
	"discordrm/api/pkg/langs"
	"discordrm/api/pkg/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func RegisterValidate(c *fiber.Ctx) error {
	payload := entities.RegisterRequest{}

	// body parser
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": langs.ErrBadRequest,
		})
	}

	// validate
	if err := utils.Validate().Struct(payload); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": langs.ErrBadRequest,
		})
	}

	return c.Next()
}
