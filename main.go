package main

import (
	"fiber-crud-mysql/book"
	"fiber-crud-mysql/config"
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

	b := api.Group("/book")
	b.Get("/", book.GetBooks)
	b.Get("/:id", book.GetBook)
	b.Post("/", book.CreateBook)
	b.Put("/:id", book.UpdateBook)
	b.Delete("/:id", book.DeleteBook)
}

func initDatabase() {
	var err error

	DATABASE_URI := os.Getenv("DATABASE_URI")

	config.Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	config.Database.AutoMigrate(&models.Book{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()

	setupDotEnv()

	setupRoutes(app)

	initDatabase()

	log.Fatal(app.Listen(":3000"))
}
