package repositories

import (
	"context"
	"fmt"

	"github.com/HEEPOKE/backend-challenge-test/internals/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepository struct {
	conn *mongo.Client
}

func NewUsersRepository(conn *mongo.Client) *UsersRepository {
	return &UsersRepository{
		conn: conn,
	}
}

func (repo *UsersRepository) CreateUser(ctx context.Context, user models.User) (*mongo.InsertOneResult, error) {
	collection := repo.conn.Database("test").Collection("users")
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %v", err)
	}
	return result, nil
}

func (repo *UsersRepository) GetUserByID(ctx context.Context, userID string) (models.User, error) {
	var user models.User
	collection := repo.conn.Database("test").Collection("users")
	err := collection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.User{}, fmt.Errorf("user not found")
		}
		return models.User{}, fmt.Errorf("failed to fetch user: %v", err)
	}
	return user, nil
}

func (repo *UsersRepository) ListUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	collection := repo.conn.Database("test").Collection("users")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %v", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, fmt.Errorf("failed to decode user: %v", err)
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}
	return users, nil
}

func (repo *UsersRepository) UpdateUser(ctx context.Context, user models.User) (*mongo.UpdateResult, error) {
	collection := repo.conn.Database("test").Collection("users")
	update := bson.M{}
	if user.Name != "" {
		update["name"] = user.Name
	}
	if user.Email != "" {
		update["email"] = user.Email
	}

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": user.ID},
		bson.M{"$set": update},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %v", err)
	}
	return result, nil
}

func (repo *UsersRepository) DeleteUser(ctx context.Context, userID string) (*mongo.DeleteResult, error) {
	collection := repo.conn.Database("test").Collection("users")
	result, err := collection.DeleteOne(ctx, bson.M{"_id": userID})
	if err != nil {
		return nil, fmt.Errorf("failed to delete user: %v", err)
	}
	return result, nil
}
