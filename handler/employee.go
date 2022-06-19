package handler

import (
	"echo-crud-postgresql/config"
	"echo-crud-postgresql/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetEmployees(c echo.Context) error {
	db := config.DB

	// Find employees
	employees := []model.Employee{}

	sqlStatement := "SELECT * FROM employees ORDER BY id"
	rows, err := db.Query(sqlStatement)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.Response{
			Status:  "Error",
			Message: err.Error(),
		})
	}
	defer rows.Close()

	// Push employee to array
	for rows.Next() {
		employee := model.Employee{}
		rows.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.Salary, &employee.Age)
		employees = append(employees, employee)
	}

	// Success
	result := model.Response{
		Status:  "Success",
		Message: "Get employees sucess",
		Data:    employees,
	}

	return c.JSON(http.StatusOK, result)
}

func GetEmployee(c echo.Context) error {
	db := config.DB

	// Find employee
	id := c.Param("id")
	var employee model.Employee

	sqlStatement := "SELECT * FROM employees WHERE id = $1"
	err := db.QueryRow(sqlStatement, id).Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.Salary, &employee.Age)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  "Error",
			Message: "Employee not found",
		})
	}

	// Success
	result := model.Response{
		Status:  "Success",
		Message: "Get employee sucess",
		Data:    employee,
	}

	return c.JSON(http.StatusOK, result)
}

func CreateEmployee(c echo.Context) error {
	db := config.DB

	// Parser
	var employee model.Employee

	err := c.Bind(&employee)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	// Create employee
	sqlStatement := "INSERT INTO employees (first_name, last_name, salary, age) VALUES ($1, $2, $3, $4) RETURNING id"
	err = db.QueryRow(sqlStatement, employee.FirstName, employee.LastName, employee.Salary, employee.Age).Scan(&employee.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	// Success
	result := model.Response{
		Status:  "Success",
		Message: "Create employee success",
		Data:    employee,
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateEmployee(c echo.Context) error {
	db := config.DB

	// Parser
	var employee model.Employee

	err := c.Bind(&employee)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	// Update employee
	id := c.Param("id")
	employee.ID = &id

	sqlStatement := "UPDATE employees SET first_name = $1, last_name = $2, salary = $3, age = $4 WHERE id = $5"
	_, err = db.Query(sqlStatement, employee.FirstName, employee.LastName, employee.Salary, employee.Age, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	// Success
	result := model.Response{
		Status:  "Success",
		Message: "Update employee success",
		Data:    employee,
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteEmployee(c echo.Context) error {
	db := config.DB

	// Delete employee
	id := c.Param("id")

	sqlStatement := "DELETE FROM employees WHERE id = $1"
	_, err := db.Query(sqlStatement, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Status:  "Error",
			Message: err.Error(),
		})
	}

	// Success
	result := model.Response{
		Status:  "Success",
		Message: "Delete employee success",
	}

	return c.JSON(http.StatusOK, result)
}
