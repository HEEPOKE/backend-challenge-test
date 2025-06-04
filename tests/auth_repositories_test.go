package repositories

import (
	"context"
	"testing"

	"github.com/HEEPOKE/backend-challenge-test/internals/core/common"
	"github.com/HEEPOKE/backend-challenge-test/internals/domain/models"
	"github.com/HEEPOKE/backend-challenge-test/internals/domain/repositories"
	"github.com/HEEPOKE/backend-challenge-test/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestRegisterRepository(t *testing.T) {
	mockClient := new(mocks.MockClient)
	mockCollection := new(mocks.MockCollection)

	mockDatabase := new(mocks.MockDatabase)

	mockClient.On("Database", "test").Return(mockDatabase)

	mockDatabase.On("Collection", "users").Return(mockCollection)

	mockCollection.On("InsertOne", mock.Anything, mock.Anything).Return(&mongo.InsertOneResult{InsertedID: "12345"}, nil)

	repo := repositories.NewAuthRepository(mockClient)

	user := models.User{
		Name:     "Jane Doe",
		Email:    "jane.doe@example.com",
		Password: "password123",
	}

	result, err := repo.Register(context.Background(), user)

	assert.Nil(t, err)
	assert.Equal(t, "12345", result.InsertedID)

	mockClient.AssertExpectations(t)
	mockCollection.AssertExpectations(t)
}

func TestAuthenticateRepository(t *testing.T) {
	mockClient := new(mocks.MockClient)
	mockCollection := new(mocks.MockCollection)

	mockDatabase := new(mocks.MockDatabase)

	mockClient.On("Database", "test").Return(mockDatabase)

	mockDatabase.On("Collection", "users").Return(mockCollection)

	mockCollection.On("FindOne", mock.Anything, mock.Anything).Return(&mongo.SingleResult{})

	repo := repositories.NewAuthRepository(mockClient)

	user := models.User{
		ID:       12345,
		Name:     "Jane Doe",
		Email:    "jane.doe@example.com",
		Password: "$2a$12$C0uwG9l5aFKNXKp2jSh3t.JWzrtQ5m3nmti/cWsTRD1ZyfMzVbXS6",
	}

	mockDatabase.On("Collection", "users").Return(mockCollection)

	common.CompareHashAndPassword = func(hashedPassword, password string) error {
		return nil
	}
	common.GenerateJWT = func(email string) (string, error) {
		return "mockedJWTToken", nil
	}

	token, err := repo.Authenticate(context.Background(), "jane.doe@example.com", "password123")

	assert.Nil(t, err)
	assert.Equal(t, "mockedJWTToken", token)

	mockClient.AssertExpectations(t)
	mockCollection.AssertExpectations(t)
}
