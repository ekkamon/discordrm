package auth

import (
	"discordrm/api/interval/entities"
	"discordrm/api/interval/models"
	"discordrm/api/pkg/langs"
	"discordrm/api/pkg/utils"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func LoginWithPasswordValidate(c *fiber.Ctx) (models.User, int, error) {
	payload := entities.LoginWithPasswordRequest{}

	// body parser
	if err := c.BodyParser(&payload); err != nil {
		return models.User{}, http.StatusBadRequest, errors.New(langs.ErrBadRequest)
	}

	// validate
	if err := utils.Validate().Struct(payload); err != nil {
		return models.User{}, http.StatusBadRequest, errors.New(langs.ErrBadRequest)
	}

	// body parser payload -> gorm model
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return models.User{}, http.StatusInternalServerError, errors.New(langs.ErrInternalServer)
	}

	return user, http.StatusOK, nil
}
