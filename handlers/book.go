package handlers

import (
	"fiber-crud-mysql/configs"
	"fiber-crud-mysql/models"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	var books []models.Book

	configs.Database.Find(&books)

	return c.Status(200).JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	result := configs.Database.Find(&book, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&book)
}

func CreateBook(c *fiber.Ctx) error {
	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	configs.Database.Create(&book)
	return c.Status(201).JSON(book)
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	book := new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	configs.Database.Where("id = ?", id).Updates(&book)
	return c.Status(200).JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book models.Book

	result := configs.Database.Delete(&book, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
