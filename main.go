package main

import (
	"context"
	"hotel-reservation/db"
	"hotel-reservation/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(
			map[string]string{"error": err.Error()},
		)
	},
}

func main() {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	// userStore := db.NewMongoUserStore(client, db.DBNAME)
	// userHandler := handeler.NewUserHandler(userStore)

	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)
	hotelStore.InsertHotel(
		context.Background(),
		&models.Hotel{
			Name:     "gho",
			Location: "kerman",
		},
	)
	// hotelHandler := handeler.NewUserHandler(userStore)

	// app := fiber.New(config)
	// app.Get("/user/:id", userHandler.HandleGetUser)
	// app.Get("/users", userHandler.HandleGetUsers)
	// app.Post("/user", userHandler.HandlePostUser)
	// app.Put("/user/:id", userHandler.HandlePutUser)
	// app.Delete("/user/:id", userHandler.HandleDeleteUser)
	// app.Listen(":8080")

}
