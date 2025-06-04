package interfaces

import (
	"context"

	"github.com/HEEPOKE/backend-challenge-test/internals/domain/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepository interface {
	CreateUser(ctx context.Context, user models.User) (*mongo.InsertOneResult, error)
	GetUserByID(ctx context.Context, userID string) (models.User, error)
	ListUsers(ctx context.Context) ([]models.User, error)
	UpdateUser(ctx context.Context, user models.User) (*mongo.UpdateResult, error)
	DeleteUser(ctx context.Context, userID string) (*mongo.DeleteResult, error)
}
