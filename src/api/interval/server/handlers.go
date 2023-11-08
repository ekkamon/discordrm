package server

import (
	"discordrm/api/pkg/errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) MapHandlers() error {
	s.Fiber.Use(func(ctx *fiber.Ctx) error {
		status := http.StatusNotFound
		return ctx.Status(status).JSON(errors.NewRespError(status, fiber.ErrNotFound.Message))
	})

	return nil
}
