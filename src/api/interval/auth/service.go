package auth

import (
	"discordrm/api/interval/entities"
	"discordrm/api/interval/models"
	"discordrm/api/interval/user"
	"discordrm/api/pkg/langs"
	"discordrm/api/pkg/utils"
	"errors"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	LoginWithPassword(user models.User, tokenData *entities.TokenData) (int, error)
}

type service struct {
	repo     Repository
	userRepo user.Repository
}

func NewService(repo Repository, userRepo user.Repository) Service {
	return service{repo, userRepo}
}

func (s service) LoginWithPassword(user models.User, tokenData *entities.TokenData) (int, error) {
	// database check
	userData, err := s.userRepo.GetByUsername(user.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusBadRequest, errors.New(langs.ErrLoginWithPasswordFailure)
		} else {
			return http.StatusInternalServerError, errors.New(langs.ErrInternalServer)
		}
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password)); err != nil {
		return http.StatusBadRequest, errors.New(langs.ErrLoginWithPasswordFailure)
	}

	// create token
	expiredIn := 30 * 86400
	addedTime := time.Now().Add(time.Duration(expiredIn) * time.Second).Unix()
	token, err := utils.CreateToken(userData, addedTime)
	if err != nil {
		return http.StatusInternalServerError, errors.New(langs.ErrInternalServer)
	}

	// save token in redis
	if err := s.repo.SetToken(userData.UID, token, expiredIn); err != nil {
		return http.StatusInternalServerError, errors.New(langs.ErrInternalServer)
	}

	*tokenData = entities.TokenData{
		Token:   token,
		Expired: addedTime,
	}

	return http.StatusOK, nil
}
