package user

import (
	"discordrm/api/interval/models"
	"discordrm/api/pkg/databases"
	"errors"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user models.User) (models.User, error)
	GetByUID(uid int) (models.User, error)
	GetByUsername(username string) (models.User, error)
}

type repository struct {
	db *databases.Conn
}

func NewRepository(db *databases.Conn) Repository {
	return repository{db}
}

func (r repository) Create(user models.User) (models.User, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if res := r.db.PgSql.Omit("ID").Create(&user); res.Error != nil {
		if !strings.Contains(res.Error.Error(), "duplicate key") {
			logrus.WithFields(logrus.Fields{
				"file": "user/repository.go",
				"func": "Create",
			}).Error(res.Error)
		}

		return models.User{}, res.Error
	}

	return user, nil
}

func (r repository) GetByUID(uid int) (models.User, error) {
	user := models.User{}

	res := r.db.PgSql.Where(&models.User{UID: uid}).First(&user)
	if res.Error != nil {
		if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
			logrus.WithFields(logrus.Fields{
				"file": "user/repository.go",
				"func": "GetByUID",
			}).Error(res.Error)
		}

		return models.User{}, res.Error
	}

	return user, nil
}

func (r repository) GetByUsername(username string) (models.User, error) {
	user := models.User{}

	res := r.db.PgSql.Where(&models.User{Username: username}).First(&user)
	if res.Error != nil {
		if !errors.Is(res.Error, gorm.ErrRecordNotFound) {
			logrus.WithFields(logrus.Fields{
				"file": "user/repository.go",
				"func": "GetByUsername",
			}).Error(res.Error)
		}

		return models.User{}, res.Error
	}

	return user, nil
}
