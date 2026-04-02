package handler

import (
	"smp/models"
	"smp/service"

	"github.com/gofiber/fiber/v3"
)

type StudentHandler struct {
	Service *service.StudentService
}

func NewStudentHandler(service *service.StudentService) *StudentHandler {
	return &StudentHandler{
		Service: service,
	}
}
func (h *StudentHandler) EnterStudentDetails(c fiber.Ctx) error {

	var student models.Student

	if err := c.Bind().Body(&student); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	id, err := h.Service.EnterStudentDetails(c.Context(), student)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "student created",
		"id":      id,
	})
}
