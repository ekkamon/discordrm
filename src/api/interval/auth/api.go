package auth

import (
	"discordrm/api/pkg/langs"
	"discordrm/api/pkg/utils"
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

	// login check
	user, status, err = r.service.LoginWithPassword(user)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":  status,
			"message": err.Error(),
		})
	}

	// generate access token
	token, expired, err := utils.CreateToken(user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": langs.ErrInternalServer,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": http.StatusOK,
		"data": fiber.Map{
			"access_token": token,
			"expired":      expired,
		},
	})

}
