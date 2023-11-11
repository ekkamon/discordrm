package user

import (
	"discordrm/api/interval/middlewares"
	"discordrm/api/pkg/databases"

	"github.com/gofiber/fiber/v2"
)

func Handler(fiber fiber.Router, db *databases.Conn) {
	res := routes{
		service: NewService(NewRepository(db)),
	}

	r := fiber.Group("/user")

	r.Post("/register", res.register)

	// protected routes
	r.Use(middlewares.IsAuthenticated())
	r.Get("/me", res.me)
}
