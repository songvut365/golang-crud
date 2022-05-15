package handler

import (
	"echo-crud-postgresql/config"
	"echo-crud-postgresql/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetEmployees(c echo.Context) error {
	db := config.DB

	employees := []model.Employee{}

	sqlStatement := "SELECT * FROM employees ORDER BY id"
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}
	defer rows.Close()

	for rows.Next() {
		employee := model.Employee{}
		rows.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.Salary, &employee.Age)
		employees = append(employees, employee)
	}

	result := model.Response{
		Status:  "Success",
		Message: "Get employees sucess",
		Data:    employees,
	}

	return c.JSON(http.StatusOK, result)
}

func GetEmployee(c echo.Context) error {
	id := c.Param("id")

	return c.String(http.StatusOK, "Get Employee"+id)
}

func CreateEmployee(c echo.Context) error {
	return c.String(http.StatusOK, "Create Employee")
}

func UpdateEmployee(c echo.Context) error {
	id := c.Param("id")

	return c.String(http.StatusOK, "Update Employee"+id)
}

func DeleteEmployee(c echo.Context) error {
	id := c.Param("id")

	return c.String(http.StatusOK, "Delete Employee"+id)
}
