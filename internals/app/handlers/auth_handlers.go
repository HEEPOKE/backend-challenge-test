package handlers

import (
	"github.com/HEEPOKE/backend-challenge-test/internals/app/services"
	"github.com/HEEPOKE/backend-challenge-test/internals/domain/models"
	"github.com/HEEPOKE/backend-challenge-test/internals/domain/models/requests"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authServices services.AuthServices
}

func NewAuthHandler(authServices services.AuthServices) *AuthHandler {
	return &AuthHandler{
		authServices: authServices,
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var dataRequest requests.AuthRequest
	if err := c.BodyParser(&dataRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request body",
		})
	}
	user := models.User{
		Name:     dataRequest.Name,
		Email:    dataRequest.Email,
		Password: dataRequest.Password,
	}

	_, err := h.authServices.Register(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to register user",
		})
	}

	token, err := h.authServices.Authenticate(dataRequest.Email, dataRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Failed to authenticate user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"token": token,
			"user": fiber.Map{
				"name":  user.Name,
				"email": user.Email,
				"id":    user.ID,
			},
		},
	})
}
