# Echo CRUD PostgreSQL

## How to run

```cmd
$ go mod download
$ go run main.go
```

## Database Setup

This project doesn't use GORM or auto migration, so databases and tables have to be created manually.

```cmd

```

## Environment

```cmd
DB_SERVER=postgres
DB_USER=postgres
DB_PASSWORD=1234
DB_NAME=company
```

## APIs

- GET /employee/
- GET /employee/:id
- POST /employee/
- PUT /employee/
- DELETE /employee/

## Reference

[Echo](https://echo.labstack.com/)