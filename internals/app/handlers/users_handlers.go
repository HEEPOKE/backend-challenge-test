package handlers

import (
	"github.com/HEEPOKE/backend-challenge-test/internals/app/services"
	"github.com/HEEPOKE/backend-challenge-test/internals/domain/models"
	"github.com/gofiber/fiber/v2"
)

type UsersHandler struct {
	userServices services.UsersServices
}

func NewUsersHandler(userServices services.UsersServices) *UsersHandler {
	return &UsersHandler{
		userServices: userServices,
	}
}

func (h *UsersHandler) CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	result, err := h.userServices.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(result)
}

func (h *UsersHandler) GetUserByID(c *fiber.Ctx) error {
	userID := c.Params("id")
	user, err := h.userServices.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *UsersHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.userServices.ListUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(users)
}

func (h *UsersHandler) UpdateUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	result, err := h.userServices.UpdateUser(userID, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (h *UsersHandler) DeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	result, err := h.userServices.DeleteUser(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}
