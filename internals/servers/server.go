package server

import (
	"strings"

	"github.com/HEEPOKE/backend-challenge-test/internals/core/middlewares"
	"github.com/HEEPOKE/backend-challenge-test/internals/servers/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	fib *fiber.App
	db  *mongo.Client
}

func NewServer(db *mongo.Client) *Server {
	app := fiber.New(fiber.Config{
		ServerHeader:      "Fiber",
		CaseSensitive:     true,
		StrictRouting:     true,
		Prefork:           false,
		EnablePrintRoutes: true,
		BodyLimit:         50 * 1024 * 1024,
	})

	app.Use(middlewares.LogMiddleware())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPatch,
			fiber.MethodDelete,
		}, ","),
	}))
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	return &Server{
		fib: app,
		db:  db,
	}
}

func (s *Server) Init() *fiber.App {
	routes.SetupRoutesAuth(s.fib, s.db)
	routes.SetupRoutesUsers(s.fib, s.db)

	return s.fib
}
