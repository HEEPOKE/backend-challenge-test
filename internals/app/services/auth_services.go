package services

import (
	"context"

	"github.com/HEEPOKE/backend-challenge-test/internals/core/interfaces"
	"github.com/HEEPOKE/backend-challenge-test/internals/domain/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthServices struct {
	authRepository interfaces.AuthRepository
}

func NewAuthServices(authRepository interfaces.AuthRepository) *AuthServices {
	return &AuthServices{authRepository: authRepository}
}

func (s *AuthServices) Register(user models.User) (*mongo.InsertOneResult, error) {
	result, err := s.authRepository.Register(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *AuthServices) Authenticate(email, password string) (string, error) {
	token, err := s.authRepository.Authenticate(context.Background(), email, password)
	if err != nil {
		return "", err
	}
	return token, nil
}
