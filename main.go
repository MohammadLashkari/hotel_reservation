package main

import (
	"hotel-reservation/db"
	"hotel-reservation/handler"

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

	client := db.InitMongo()
	app := fiber.New(config)

	var (
		userStore   = db.NewMongoUserStore(client)
		hotelStrore = db.NewMongoHotelStore(client)
		roomStrore  = db.NewMongoRoomStore(client, hotelStrore)
		store       = &db.Store{
			RoomStore:  roomStrore,
			HotelStore: hotelStrore,
			UserStore:  userStore,
		}
		// handlers
		userHandler  = handler.NewUserHandler(store)
		hotelHandler = handler.NewHotelHandler(store)
	)

	// user
	app.Get("/user/:id", userHandler.HandleGetUser)
	app.Get("/user", userHandler.HandleGetUsers)
	app.Post("/user", userHandler.HandlePostUser)
	app.Put("/user/:id", userHandler.HandlePutUser)
	app.Delete("/user/:id", userHandler.HandleDeleteUser)
	// hotel
	app.Get("/hotel", hotelHandler.HandleGetHotels)
	app.Get("/hotel/:id", hotelHandler.HandleGetHotel)
	app.Get("/hotel/:id/rooms", hotelHandler.HandleGetHotelRooms)
	app.Listen(":8080")

}
