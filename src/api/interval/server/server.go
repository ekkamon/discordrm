package server

import (
	"discordrm/api/config"
	"discordrm/api/pkg/databases"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Fiber *fiber.App
	Cfg   *config.Config
	DB    *databases.Conn
}

func NewServer(cfg *config.Config, dbs *databases.Conn) *Server {
	return &Server{
		Fiber: fiber.New(),
		Cfg:   cfg,
		DB:    dbs,
	}
}

func (s *Server) Start() {
	if err := s.MapHandlers(); err != nil {
		panic(err)
	}

	fiberURL := fmt.Sprintf("%s:%s", s.Cfg.Server.Host, s.Cfg.Server.Port)

	logrus.Infof("Server had been startd on %s", fiberURL)

	if err := s.Fiber.Listen(fiberURL); err != nil {
		panic(err)
	}
}
