package server

import (
	"discordrm/api/interval/user"
	"discordrm/api/pkg/langs"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) MapHandlers() error {
	v1 := s.Fiber.Group("/v1")

	user.Handler(v1, s.DB)

	// catch not found end-points
	s.Fiber.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  http.StatusNotFound,
			"message": langs.ErrNotFound,
		})
	})

	return nil
}
