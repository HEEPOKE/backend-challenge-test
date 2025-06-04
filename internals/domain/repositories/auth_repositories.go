package repositories

import (
	"context"
	"fmt"

	"github.com/HEEPOKE/backend-challenge-test/internals/core/common"
	"github.com/HEEPOKE/backend-challenge-test/internals/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository struct {
	conn *mongo.Client
}

func NewAuthRepository(conn *mongo.Client) *AuthRepository {
	return &AuthRepository{
		conn: conn,
	}
}

func (repo *AuthRepository) Register(ctx context.Context, user models.User) (*mongo.InsertOneResult, error) {
	hashedPassword, err := common.HashPassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}
	user.Password = string(hashedPassword)

	collection := repo.conn.Database("test").Collection("users")
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %v", err)
	}

	return result, nil
}

func (repo *AuthRepository) Authenticate(ctx context.Context, email, password string) (string, error) {
	collection := repo.conn.Database("test").Collection("users")
	var user models.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", fmt.Errorf("user not found")
		}
		return "", fmt.Errorf("failed to retrieve user: %v", err)
	}

	err = common.CompareHashAndPassword(user.Password, password)
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	token, err := common.GenerateJWT(user.Email)
	if err != nil {
		return "", fmt.Errorf("error generating token: %v", err)
	}

	return token, nil
}
