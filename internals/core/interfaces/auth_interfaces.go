package interfaces

import (
	"context"

	"github.com/HEEPOKE/backend-challenge-test/internals/domain/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository interface {
	Register(ctx context.Context, user models.User) (*mongo.InsertOneResult, error)
	Authenticate(ctx context.Context, email, password string) (string, error)
}
