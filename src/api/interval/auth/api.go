package auth

import (
	"discordrm/api/interval/entities"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type routes struct {
	service Service
}

func (r *routes) loginWithPassword(c *fiber.Ctx) error {
	// payload validate
	user, status, err := LoginWithPasswordValidate(c)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":  status,
			"message": err.Error(),
		})
	}

	res := entities.TokenData{}
	status, err = r.service.LoginWithPassword(user, &res)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":  status,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": http.StatusOK,
		"data":   res,
	})

}
