package routes

import (
	"fiber-todo-mongo/controllers"

	"github.com/gofiber/fiber/v2"
)

func TodoRoutes(app fiber.Router) {
	todo := app.Group("/todo")

	todo.Post("/", controllers.CreateTodo)
	todo.Get("/", controllers.GetTodos)
	todo.Get("/:id", controllers.GetTodo)
	todo.Put("/:id", controllers.UpdateTodo)
	todo.Delete("/:id", controllers.DeleteTodo)
}
