package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/kauanpecanha/odsquiz-auth/internal/apperrors"
	"github.com/kauanpecanha/odsquiz-auth/internal/models"
	"github.com/kauanpecanha/odsquiz-auth/internal/services"
)

type Handler struct {
	Service *services.Service
}

func (h *Handler) CreateOne(c fiber.Ctx) error {
	one := new(models.User)

	if err := c.Bind().Body(one); err != nil {
		return respondError(c, apperrors.BadRequest(
			apperrors.CodeInvalidRequest,
			err,
		))
	}

	createdOne, err := h.Service.CreateOne(one)
	if err != nil {
		return respondError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(createdOne)
}

func (h *Handler) Login(c fiber.Ctx) error {
	one := new(models.LoginRequest)
	if err := c.Bind().Body(one); err != nil {
		return respondError(c, apperrors.BadRequest(
			apperrors.CodeInvalidRequest,
			err,
		))
	}

	token, err := h.Service.Login(one)
	if err != nil {
		return respondError(c, err)
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"token": token,
	})
}

func (h *Handler) GetAllOnes(c fiber.Ctx) error {
	ones, err := h.Service.GetAllOnes()
	if err != nil {
		return respondError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(ones)
}

func (h *Handler) GetOneByID(c fiber.Ctx) error {
	id := c.Params("id")

	_, err := uuid.Parse(id)
	if err != nil {
		return respondError(c, apperrors.BadRequest(
			apperrors.CodeInvalidRequest,
			err,
		))
	}

	one, err := h.Service.GetOneByID(id)
	if err != nil {
		return respondError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(one)
}

func (h *Handler) UpdateOne(c fiber.Ctx) error {
	id := c.Params("id")

	_, err := uuid.Parse(id)
	if err != nil {
		return respondError(c, apperrors.BadRequest(
			apperrors.CodeInvalidRequest,
			err,
		))
	}

	one := new(models.User)

	if err := c.Bind().Body(one); err != nil {
		return respondError(c, apperrors.BadRequest(
			apperrors.CodeInvalidRequest,
			err,
		))
	}

	// force route param ID
	one.ID = id

	updatedOne, err := h.Service.UpdateOne(one)
	if err != nil {
		return respondError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(updatedOne)
}

func (h *Handler) DeleteOne(c fiber.Ctx) error {
	id := c.Params("id")

	_, err := uuid.Parse(id)
	if err != nil {
		return respondError(c, apperrors.BadRequest(
			apperrors.CodeInvalidRequest,
			err,
		))
	}

	err = h.Service.DeleteOne(id)
	if err != nil {
		return respondError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON("deleted successfully")
}

type errorResponse struct {
	Code apperrors.Code `json:"code"`
}

func respondError(c fiber.Ctx, err error) error {
	appErr, ok := apperrors.From(err)
	if !ok {
		appErr = apperrors.Internal(err)
	}

	return c.Status(statusFor(appErr.Kind)).JSON(errorResponse{
		Code: appErr.Code,
	})
}

func statusFor(kind apperrors.Kind) int {
	switch kind {
	case apperrors.KindBadRequest:
		return fiber.StatusBadRequest
	case apperrors.KindUnauthorized:
		return fiber.StatusUnauthorized
	case apperrors.KindNotFound:
		return fiber.StatusNotFound
	case apperrors.KindConflict:
		return fiber.StatusConflict
	default:
		return fiber.StatusInternalServerError
	}
}
