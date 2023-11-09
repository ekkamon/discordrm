package server

import (
	"discordrm/api/pkg/langs"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) MapHandlers() error {
	s.Fiber.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  http.StatusNotFound,
			"message": langs.ErrNotFound,
		})
	})

	return nil
}
