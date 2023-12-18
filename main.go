package main

import (
	"hotel-reservation/db"
	"hotel-reservation/handeler"

	"github.com/gofiber/fiber/v2"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(
			map[string]string{"error": err.Error()},
		)
	},
}

func main() {

	client := db.Init()

	userStore := db.NewMongoUserStore(client)
	userHandler := handeler.NewUserHandler(userStore)

	app := fiber.New(config)
	app.Get("/user/:id", userHandler.HandleGetUser)
	app.Get("/users", userHandler.HandleGetUsers)
	app.Post("/user", userHandler.HandlePostUser)
	app.Put("/user/:id", userHandler.HandlePutUser)
	app.Delete("/user/:id", userHandler.HandleDeleteUser)
	app.Listen(":8080")

}
