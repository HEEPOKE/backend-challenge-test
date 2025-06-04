package routes

import (
	"github.com/HEEPOKE/backend-challenge-test/internals/app/handlers"
	"github.com/HEEPOKE/backend-challenge-test/internals/app/services"
	"github.com/HEEPOKE/backend-challenge-test/internals/domain/repositories"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutesUsers(app *fiber.App, db *mongo.Client) {
	usersHandler := handlers.NewUsersHandler(*services.NewUsersServices(repositories.NewUsersRepository(db)))

	app.Post("/users", usersHandler.CreateUser)
	app.Get("/users/:id", usersHandler.GetUserByID)
	app.Get("/users", usersHandler.ListUsers)
	app.Put("/users/:id", usersHandler.UpdateUser)
	app.Delete("/users/:id", usersHandler.DeleteUser)
}
