# Fiber CRUD MySQL

## Create Project

Create project name `fiber-crud-mysql` and initial as follows:

```cmd
mkdir fiber-crud-mysql
cd fiber-crud-mysql
go mod init fiber-crud-mysql
```

## Folder Structure

```text
FIBER-CRUD-MYSQL
│ .env
│ go.mod
│ go.sum
│ main.go
│
├───configs
│     database.go
│
├───handlers
│     book.go
│
└───models
      book.go
```

## Installation

Install dependencies, which are as follows:

```cmd
go get github.com/gofiber/fiber/v2
go get gorm.io/gorm
go get gorm.io/driver/mysql
go get github.com/joho/godotenv
```

## Environment

Add your application configuration to `.env` file in the root of project:

- Database username: `root`
- Database password: `1234`
- Database hostname: `localhost`
- Database port: `3306`
- Database schema: `songvut` 

```text
DATABASE_URI=root:1234@tcp(localhost:3306)/songvut?charset=utf8mb4&parseTime=True&loc=Local
```

## APIs

- POST /api/book
- GET /api/books
- GET /api/book/:id
- PUT /api/book/:id
- DELETE /api/book/:id

## Referrence

- [Examples for Fiber](https://github.com/gofiber/recipes)
- [How to Build REST API using Go Fiber and Gorm ORM](https://dev.to/franciscomendes10866/how-to-build-rest-api-using-go-fiber-and-gorm-orm-2jbe)
- [GORM with MySQL](https://gorm.io/docs/connecting_to_the_database.html#MySQL)
