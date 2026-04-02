package routes

import (
	"smp/handler"

	"github.com/gofiber/fiber/v3"
)

func SetupUserRoutes(
	app *fiber.App,
	userHandler *handler.UserHandler,
	studentHandler *handler.StudentHandler,
	classHandler *handler.ClassroomHandler,
) {

	app.Post("/signup", userHandler.OnboardUsers)
	app.Post("/resend-otp", userHandler.ResendOTP)
	app.Post("/verify-otp", userHandler.VerifyOTP)
	app.Post("/login", userHandler.Login)

	app.Post("/student-details", studentHandler.EnterStudentDetails)

	app.Post("/classrooms", classHandler.CreateClassroom)
}
