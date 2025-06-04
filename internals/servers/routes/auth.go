package routes

import (
	"github.com/HEEPOKE/backend-challenge-test/internals/app/handlers"
	"github.com/HEEPOKE/backend-challenge-test/internals/app/services"
	"github.com/HEEPOKE/backend-challenge-test/internals/domain/repositories"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutesAuth(app *fiber.App, db *mongo.Client) {
	handlers := handlers.NewAuthHandler(*services.NewAuthServices(repositories.NewAuthRepository(db)))

	auth := app.Group("/apis/auth")

	auth.Post("/register", handlers.Register)
}
