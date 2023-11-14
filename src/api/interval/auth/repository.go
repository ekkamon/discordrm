package auth

import (
	"discordrm/api/pkg/databases"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type Repository interface {
	SetToken(uid int, token string, lifeTime int) error
}

type repository struct {
	db *databases.Conn
}

func NewRepository(db *databases.Conn) Repository {
	return repository{db}
}

func (r repository) SetToken(uid int, token string, lifeTime int) error {
	prefix := fmt.Sprintf("auth:%d", uid)
	time := time.Duration(lifeTime) * time.Second

	if err := r.db.Redis.Client.Set(r.db.Redis.Ctx, prefix, token, time).Err(); err != nil {
		logrus.WithFields(logrus.Fields{
			"file": "auth/repository.go",
			"func": "SetToken",
		}).Error(err)
		return err
	}

	return nil
}
