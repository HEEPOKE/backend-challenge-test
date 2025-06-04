package middlewares

import (
	"fmt"

	"github.com/HEEPOKE/backend-challenge-test/pkg/configs"
	jwtWare "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func JwtMiddleware() fiber.Handler {
	return jwtWare.New(jwtWare.Config{
		SigningKey: jwtWare.SigningKey{Key: []byte(configs.Cfg.JWT_SECRET_KEY)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": fmt.Sprintf("Unauthorized: %s", err.Error()),
			})
		},
	})
}
