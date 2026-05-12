// internal/handlers/User.go: Implement HTTP requests and responses handling
package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/kauanpecanha/odsquiz-auth/internal/models"
	"github.com/kauanpecanha/odsquiz-auth/internal/services"
)

// UserHandler defines the connection between handler and service
type UserHandler struct {
	Service *services.UserService
}

func (h *UserHandler) CreateUser(c fiber.Ctx) error {
	user := new(models.User)

	if err := c.Bind().Body(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	createdUser, err := h.Service.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(createdUser)
}

func (h *UserHandler) LoginUser(c fiber.Ctx) error {
	user := new(models.LoginUserRequest)
	if err := c.Bind().Body(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	token, err := h.Service.LoginUser(user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{"error": err.Error()},
		)
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"token": token,
	})
}

func (h *UserHandler) GetAllUsers(c fiber.Ctx) error {
	users, err := h.Service.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (h *UserHandler) GetUserByID(c fiber.Ctx) error {
	id := c.Params("id")

	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	user, err := h.Service.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *UserHandler) UpdateUser(c fiber.Ctx) error {
	id := c.Params("id")

	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	user := new(models.User)

	if err := c.Bind().Body(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	// force route param ID
	user.ID = id

	updatedUser, err := h.Service.UpdateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(updatedUser)
}

func (h *UserHandler) DeleteUser(c fiber.Ctx) error {
	id := c.Params("id")

	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("invalid id")
	}

	err = h.Service.DeleteUser(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON("deleted successfully")
}