package router

import (
	"echo-crud-postgresql/handler"

	"github.com/labstack/echo/v4"
)

func SetupEmployeeRoute(e *echo.Echo) {
	employee := e.Group("/employee")

	employee.GET("/", handler.GetEmployees)
	employee.GET("/:id", handler.GetEmployee)
	employee.POST("/", handler.CreateEmployee)
	employee.PUT("/:id", handler.UpdateEmployee)
	employee.DELETE("/:id", handler.DeleteEmployee)
}
