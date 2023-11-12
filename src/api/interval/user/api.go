package user

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type routes struct {
	service Service
}

func (r routes) register(c *fiber.Ctx) error {
	user, status, err := RegisterValidate(c)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":  status,
			"message": err.Error(),
		})
	}

	_, status, err = r.service.Register(user)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":  status,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": http.StatusOK,
	})
}

func (r routes) me(c *fiber.Ctx) error {
	uid := c.Locals("uid").(int)
	user, status, err := r.service.Me(uid)
	if err != nil {
		return c.Status(status).JSON(fiber.Map{
			"status":  status,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": http.StatusOK,
		"data":   user.ToPublic(),
	})
}
