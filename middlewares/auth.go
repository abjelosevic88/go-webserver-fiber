package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func Auth(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_KEY")),
	})(ctx)
}
