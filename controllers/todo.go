package controllers

import (
	"os"
	"time"

	"fiber-todo-mongo/config"
	"fiber-todo-mongo/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTodos(c *fiber.Ctx) error {
	todoCollection := config.MI.DB.Collection(os.Getenv("TODO_COLLECTION"))

	//Query to filter
	query := bson.D{{}}

	cursor, err := todoCollection.Find(c.Context(), query)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}

	var todos []models.Todo = make([]models.Todo, 0)

	err = cursor.All(c.Context(), &todos)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todos": todos,
		},
	})
}

func GetTodo(c *fiber.Ctx) error {
	todoCollection := config.MI.DB.Collection(os.Getenv("TODO_COLLECTION"))

	paramID := c.Params("id")

	id, err := primitive.ObjectIDFromHex(paramID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse ID",
			"error":   err.Error(),
		})
	}

	todo := &models.Todo{}

	query := bson.D{{Key: "_id", Value: id}}

	err = todoCollection.FindOne(c.Context(), query).Decode(todo)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Todo Not Found",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": todo,
		},
	})
}

func CreateTodo(c *fiber.Ctx) error {
	todoCollection := config.MI.DB.Collection(os.Getenv("TODO_COLLECTION"))

	data := new(models.Todo)

	err := c.BodyParser(&data)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}

	data.ID = nil
	f := false
	data.Completed = &f
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	result, err := todoCollection.InsertOne(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Cannot insert todo",
			"error":   err.Error(),
		})
	}

	todo := &models.Todo{}
	query := bson.D{{Key: "_id", Value: result.InsertedID}}

	todoCollection.FindOne(c.Context(), query).Decode(todo)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": todo,
		},
	})

}

func UpdateTodo(c *fiber.Ctx) error {
	todoCollection := config.MI.DB.Collection(os.Getenv("TODO_COLLECTION"))

	paramID := c.Params("id")

	id, err := primitive.ObjectIDFromHex(paramID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
			"error":   err,
		})
	}

	data := new(models.Todo)
	err = c.BodyParser(&data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	query := bson.D{{Key: "_id", Value: id}}

	var dataToUpdate bson.D

	if data.Title != nil {
		dataToUpdate = append(dataToUpdate, bson.E{Key: "title", Value: data.Title})
	}

	if data.Completed != nil {
		dataToUpdate = append(dataToUpdate, bson.E{Key: "completed", Value: data.Completed})
	}

	dataToUpdate = append(dataToUpdate, bson.E{Key: "updatedAt", Value: time.Now()})

	update := bson.D{
		{Key: "$set", Value: dataToUpdate},
	}

	err = todoCollection.FindOneAndUpdate(c.Context(), query, update).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Todo Not Found",
				"error":   err,
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot update tdo",
			"error":   err,
		})
	}

	todo := &models.Todo{}

	todoCollection.FindOne(c.Context(), query).Decode(todo)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"todo": todo,
		},
	})
}

func DeleteTodo(c *fiber.Ctx) error {
	todoCollection := config.MI.DB.Collection(os.Getenv("TODO_COLLECTION"))

	paramID := c.Params("id")

	id, err := primitive.ObjectIDFromHex(paramID)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse id",
			"error":   err,
		})
	}

	query := bson.D{{Key: "_id", Value: id}}

	err = todoCollection.FindOneAndDelete(c.Context(), query).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "Todo Not Found",
				"error":   err,
			})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot delete tdo",
			"error":   err,
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
