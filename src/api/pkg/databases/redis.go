package databases

import (
	"context"
	"discordrm/api/config"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type Redis struct {
	Client *redis.Client
	Ctx    context.Context
}

func NewRedisClient(cfg *config.Config) *Redis {
	rdb := &Redis{
		Client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
			Username: cfg.Redis.Username,
			Password: cfg.Redis.Password,
			DB:       0,
		}),
		Ctx: context.Background(),
	}

	_, err := rdb.Client.Ping(rdb.Ctx).Result()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "Redis",
		}).Fatal(err)
	}

	return rdb
}

func (r *Redis) SetModel(key string, value interface{}, expiredIn time.Duration) {
	// to json
	data, err := json.Marshal(value)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"file": "database/redis.go",
			"func": "SetStruct",
		}).Error(err)
	}

	r.Client.Set(r.Ctx, key, string(data), expiredIn)
}

func (r *Redis) GetModel(key string, dest interface{}) error {
	data, err := r.Client.Get(r.Ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			logrus.WithFields(logrus.Fields{
				"file": "database/redis.go",
				"func": "GetStruct",
			}).Error(err)
		}

		return err
	}

	return json.Unmarshal([]byte(data), dest)
}
