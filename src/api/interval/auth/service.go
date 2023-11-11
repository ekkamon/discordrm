package auth

import (
	"discordrm/api/interval/models"
	"discordrm/api/interval/user"
	"discordrm/api/pkg/langs"
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	LoginWithPassword(user models.User) (models.User, int, error)
}

type service struct {
	repo user.Repository
}

func NewService(repo user.Repository) Service {
	return service{repo}
}

func (s service) LoginWithPassword(user models.User) (models.User, int, error) {
	// database check
	userData, err := s.repo.GetByUsername(user.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, http.StatusBadRequest, errors.New(langs.ErrLoginWithPasswordFailure)
		} else {
			return models.User{}, http.StatusInternalServerError, errors.New(langs.ErrInternalServer)
		}
	}

	// check password
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if err != nil {
		return models.User{}, http.StatusBadRequest, errors.New(langs.ErrLoginWithPasswordFailure)
	}

	return userData, http.StatusOK, nil
}
