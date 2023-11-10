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
