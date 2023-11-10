package user

import (
	"discordrm/api/interval/models"
	"discordrm/api/pkg/langs"
	"errors"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(user models.User) (models.User, int, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return service{repo}
}

func (s service) Register(user models.User) (models.User, int, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		return models.User{}, http.StatusInternalServerError, err
	}

	user.Password = string(bytes)

	user, err = s.repo.Create(user)
	if err != nil {
		status := http.StatusBadRequest

		if strings.Contains(err.Error(), "email_key") {
			err = errors.New(langs.ErrEmailAlreadyExists)
		} else if strings.Contains(err.Error(), "username_key") {
			err = errors.New(langs.ErrUsernameAlreadyExists)
		} else {
			status = http.StatusInternalServerError
			err = errors.New(langs.ErrInternalServer)
		}

		return models.User{}, status, err
	}

	return user, http.StatusOK, nil
}
