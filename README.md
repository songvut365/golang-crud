# Fiber Todo MongoDB

## Create Project

Create project name `fiber-todo-mongo`and initial as follows:

```cmd
mkdir fiber-todo-mongo
cd fiber-todo-mongo
go mod init fiber-todo-mongo
```

## Installation

```cmd
go get github.com/gofiber/fiber/v2
go get go.mongodb.org/mongo-driver
go get go.mongodb.org/mongo-driver/bson
go get github.com/joho/godotenv
```

## Environment

Add your application configuration to `.env` file in the root of project:

```text
MONGO_URI=mongodb+srv://<username>:<password>@<clustername></clustername>.flgum.mongodb.net/myFirstDatabase?retryWrites=true&w=majority
DATABASE_NAME=fiber-todo-mongo
TODO_COLLECTION=todos
```

## APIs

- POST /api/todo
- GET /api/todo
- GET /api/todo/:id
- POST /api/todo/:id
- DELETE /api/todo/:id

## Reference

- [Golang for Web (Part-II): Gofiber REST API + Mongo DB Atlas](https://dev.to/devsmranjan/golang-for-web-part-ii-gofiber-rest-api-mongo-db-atlas-2d1i)
- [Examples for Fiber](https://github.com/gofiber/recipes/tree/master/mongodb)
- [MongoDB Go Driber](https://www.mongodb.com/docs/drivers/go/current/)