package server

import (
	"discordrm/api/config"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Fiber *fiber.App
	Cfg   *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		Fiber: fiber.New(),
		Cfg:   cfg,
	}
}

func (s *Server) Start() error {
	if err := s.MapHandlers(); err != nil {
		return err
	}

	fiberURL := fmt.Sprintf("%s:%s", s.Cfg.Host, s.Cfg.Port)

	if err := s.Fiber.Listen(fiberURL); err != nil {
		return err
	}

	log.Fatalf("[Server] had been startd on %s", fiberURL)

	return nil
}
