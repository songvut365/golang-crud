package main

import (
	"fiber-todo-mongo/config"
	"fiber-todo-mongo/routes"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")

	routes.TodoRoutes(api)
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	//dot env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//config db
	config.ConnectDB()

	//setup routes
	setupRoutes(app)

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
