package server

import (
	"discordrm/api/config"
	"discordrm/api/pkg/databases"
	"discordrm/api/pkg/redis"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Fiber *fiber.App
	Cfg   *config.Config
	DB    *databases.Conn
	Redis *redis.Redis
}

func NewServer(cfg *config.Config, dbs *databases.Conn, redis *redis.Redis) *Server {
	return &Server{
		Fiber: fiber.New(),
		Cfg:   cfg,
		DB:    dbs,
		Redis: redis,
	}
}

func (s *Server) Start() {
	if err := s.MapHandlers(); err != nil {
		logrus.WithFields(logrus.Fields{
			"file": "server/server.go",
			"func": "MapHandlers",
		}).Fatal(err)
	}

	fiberURL := fmt.Sprintf("%s:%s", s.Cfg.Server.Host, s.Cfg.Server.Port)

	logrus.Infof("Server had been startd on %s", fiberURL)

	if err := s.Fiber.Listen(fiberURL); err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "Fiber",
		}).Fatal(err)
	}
}
