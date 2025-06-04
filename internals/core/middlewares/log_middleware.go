package middlewares

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LogMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		duration := time.Since(start)
		log.Printf("Method: %s, Path: %s, Execution Time: %v", c.Method(), c.Path(), duration)

		return err
	}
}
