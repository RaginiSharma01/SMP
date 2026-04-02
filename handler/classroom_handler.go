package handler

import (
	"smp/models"
	"smp/service"

	"github.com/gofiber/fiber/v3"
)

type ClassroomHandler struct {
	Service *service.ClassroomService
}

func NewClassroomHandler(service *service.ClassroomService) *ClassroomHandler {
	return &ClassroomHandler{
		Service: service,
	}
}

func (h *ClassroomHandler) CreateClassroom(c fiber.Ctx) error {

	var classroom models.Classroom

	if err := c.Bind().Body(&classroom); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	id, err := h.Service.CreateClassroom(c.Context(), classroom)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "classroom created",
		"id":      id,
	})
}
