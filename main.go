package main

import (
	"fmt"
	"log"
	"smp/config"
	"smp/db"
	"smp/handler"
	"smp/repository"
	"smp/routes"
	"smp/service"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("Welcome to Student management portal")

	godotenv.Load()
	cfg := config.LoadConfig()

	database, err := db.ConnectDb(cfg)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	defer database.Pool.Close()

	app := fiber.New()

	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("server is running")
	})

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	userRepo := repository.NewUserRepo(database.Pool)
	userService := service.NewUserService(userRepo, rdb)
	userHandler := handler.NewUserHandler(userService)

	studentRepo := repository.NewStudent(database.Pool)
	studentService := service.NewStudentService(studentRepo)
	studentHandler := handler.NewStudentHandler(studentService)

	classRepo := repository.NewClassroomRepository(database.Pool)
	classService := service.NewClassroomService(classRepo)
	classHandler := handler.NewClassroomHandler(classService)

	routes.SetupUserRoutes(app, userHandler, studentHandler, classHandler)

	log.Fatal(app.Listen(cfg.ServerPort))
}
