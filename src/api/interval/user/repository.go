package user

import (
	"discordrm/api/interval/models"
	"discordrm/api/pkg/databases"
	"time"
)

type Repository interface {
	Create(user models.User) (models.User, error)
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
		return models.User{}, res.Error
	}

	return user, nil
}

func (r repository) GetByUsername(username string) (models.User, error) {
	user := models.User{}

	res := r.db.PgSql.Where(&models.User{Username: username}).First(&user)
	if res.Error != nil {
		return models.User{}, res.Error
	}

	return user, nil
}

func (r repository) GetByEmail(email string) (models.User, error) {
	user := models.User{}

	res := r.db.PgSql.Where(&models.User{Email: email}).First(&user)
	if res.Error != nil {
		return models.User{}, res.Error
	}

	return user, nil
}
