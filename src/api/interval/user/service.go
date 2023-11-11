package user

import (
	"discordrm/api/interval/models"
	"discordrm/api/pkg/langs"
	"errors"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	Register(user models.User) (models.User, int, error)
	Me(uid int) (models.User, int, error)
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
		logrus.WithFields(logrus.Fields{
			"file":    "user/service.go",
			"service": "Register",
			"func":    "GenerateFromPassword",
		}).Error(err)

		return models.User{}, http.StatusInternalServerError, errors.New(langs.ErrInternalServer)
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
			logrus.WithFields(logrus.Fields{
				"file":    "user/service.go",
				"service": "Register",
				"func":    "Create",
			}).Error(err)

			status = http.StatusInternalServerError
			err = errors.New(langs.ErrInternalServer)
		}

		return models.User{}, status, err
	}

	return user, http.StatusOK, nil
}

func (s service) Me(uid int) (models.User, int, error) {
	user, err := s.repo.GetByUID(uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, http.StatusBadRequest, errors.New(langs.ErrBadRequest)
		} else {
			return models.User{}, http.StatusInternalServerError, errors.New(langs.ErrInternalServer)
		}
	}

	return user, http.StatusOK, nil
}
