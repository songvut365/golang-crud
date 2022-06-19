package main

import (
	"echo-crud-postgresql/config"
	"echo-crud-postgresql/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${method} ${uri} \t status:${status}\n",
	}))

	config.SetupDatabase()
	router.SetupEmployeeRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}
