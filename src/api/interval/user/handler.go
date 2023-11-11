package user

import (
	"discordrm/api/pkg/databases"

	"github.com/gofiber/fiber/v2"
)

func Handler(fiber fiber.Router, db *databases.Conn) {
	res := routes{
		service: NewService(NewRepository(db)),
	}

	r := fiber.Group("/user")

	r.Post("/register", res.register)
}
