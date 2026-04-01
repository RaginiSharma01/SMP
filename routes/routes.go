package routes

import (
	"smp/handler"

	"github.com/gofiber/fiber/v3"
)

func SetupUserRoutes(
	app *fiber.App,
	userHandler *handler.UserHandler,
) {

	app.Post("/signup", userHandler.OnboardUsers)
	app.Post("/verify-otp", userHandler.VerifyOTP)

}
