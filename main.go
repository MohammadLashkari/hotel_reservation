package main

import (
	"context"
	"hotel-reservation/db"
	"hotel-reservation/handeler"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbUri = "mongodb://localhost:27017"

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(
			map[string]string{"error": err.Error()},
		)
	},
}

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatal(err)
	}

	userStore := db.NewMongoUserStore(client)
	userHandler := handeler.NewUserHandler(userStore)

	app := fiber.New(config)
	app.Get("/user/:id", userHandler.HandleGetUser)
	app.Get("/users", userHandler.HandleGetUsers)
	app.Post("/user", userHandler.HandlePostUser)
	app.Delete("/user/:id", userHandler.HandleDeleteUser)
	app.Listen(":8080")

}
