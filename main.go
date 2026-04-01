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

	userRepo := repository.NewUserRepo(database.Pool)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	routes.SetupUserRoutes(app, userHandler)

	log.Fatal(app.Listen(cfg.ServerPort))
}
