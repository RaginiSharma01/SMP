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

func (h *UserHandler) ResendOTP(c fiber.Ctx) error {

	type request struct {
		Email string `json:"email"`
	}

	var req request

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	err := h.Service.ResendOTP(c.Context(), req.Email)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "otp sent successfully",
	})
}

func (h *UserHandler) VerifyOTP(c fiber.Ctx) error {

	type request struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}

	var req request

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	err := h.Service.VerifyOTP(c.Context(), req.Email, req.OTP)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "email verified successfully",
	})
}
