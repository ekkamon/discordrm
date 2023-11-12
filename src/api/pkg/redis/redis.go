package redis

import (
	"context"
	"discordrm/api/config"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type Redis struct {
	ctx context.Context
	db  redis.Client
}

func NewClient(cfg *config.Config) *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Username: cfg.Redis.Username,
		Password: cfg.Redis.Password,
		DB:       0,
	})

	ctx := context.Background()
	cli := &Redis{
		ctx: ctx,
		db:  *rdb,
	}

	_, err := cli.db.Ping(ctx).Result()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "Redis",
		}).Fatal(err)
	}

	return cli
}
