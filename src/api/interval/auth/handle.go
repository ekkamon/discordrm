package auth

import (
	"discordrm/api/interval/user"
	"discordrm/api/pkg/databases"

	"github.com/gofiber/fiber/v2"
)

func Handler(fiber fiber.Router, db *databases.Conn) {
	res := routes{
		service: NewService(NewRepository(db), user.NewRepository(db)),
	}
	r := fiber.Group("/auth")

	r.Post("/login-with-password", res.loginWithPassword)
}
