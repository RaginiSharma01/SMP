package handler

import (
	"smp/models"
	"smp/service"

	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) OnboardUsers(c fiber.Ctx) error {

	var user models.User

	if err := c.Bind().Body(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid details",
			"data":  fiber.Map{},
		})
	}

	id, err := h.Service.OnboardUsers(c.Context(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			"data":  fiber.Map{},
		})
	}

	return c.JSON(fiber.Map{
		"message": "Signed up successfully",
		"data": fiber.Map{
			"id": id,
		},
	})
}