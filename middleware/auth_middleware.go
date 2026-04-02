package middleware

import (
	"smp/utils"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func AuthMiddleware(c fiber.Ctx) error {

	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{
			"error": "missing authorization header",
		})
	}

	// Expect: Bearer token
	parts := strings.Split(authHeader, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(401).JSON(fiber.Map{
			"error": "invalid authorization format",
		})
	}

	token := parts[1]

	claims, err := utils.VerifyJWT(token)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "invalid token",
		})
	}

	// store user info in request context
	c.Locals("userId", claims["empId"])
	c.Locals("email", claims["email"])

	return c.Next()
}