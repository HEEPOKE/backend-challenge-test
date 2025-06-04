package services

import (
	"context"

	"github.com/HEEPOKE/backend-challenge-test/internals/core/interfaces"
	"github.com/HEEPOKE/backend-challenge-test/internals/domain/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersServices struct {
	usersInterface interfaces.UsersRepository
}

func NewUsersServices(usersInterface interfaces.UsersRepository) *UsersServices {
	return &UsersServices{usersInterface: usersInterface}
}

func (s *UsersServices) CreateUser(user models.User) (*mongo.InsertOneResult, error) {
	result, err := s.usersInterface.CreateUser(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *UsersServices) GetUserByID(userID string) (models.User, error) {
	user, err := s.usersInterface.GetUserByID(context.Background(), userID)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *UsersServices) ListUsers() ([]models.User, error) {
	users, err := s.usersInterface.ListUsers(context.Background())
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UsersServices) UpdateUser(userID string, user models.User) (*mongo.UpdateResult, error) {
	result, err := s.usersInterface.UpdateUser(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *UsersServices) DeleteUser(userID string) (*mongo.DeleteResult, error) {
	result, err := s.usersInterface.DeleteUser(context.Background(), userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
