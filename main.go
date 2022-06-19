package main

import (
	"fiber-crud-mysql/configs"
	"fiber-crud-mysql/handlers"
	"fiber-crud-mysql/models"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupDotEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")

	book := v1.Group("/book")
	book.Get("/", handlers.GetBooks)
	book.Get("/:id", handlers.GetBook)
	book.Post("/", handlers.CreateBook)
	book.Put("/:id", handlers.UpdateBook)
	book.Delete("/:id", handlers.DeleteBook)
}

func initDatabase() {
	var err error

	DATABASE_URI := os.Getenv("DATABASE_URI")

	configs.Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	configs.Database.AutoMigrate(&models.Book{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	setupDotEnv()

	setupRoutes(app)

	initDatabase()

	log.Fatal(app.Listen(":3000"))
}
