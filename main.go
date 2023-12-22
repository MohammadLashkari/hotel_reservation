package main

import (
	"fmt"
	"hotel-reservation/api"
	"hotel-reservation/api/middleware"
	"hotel-reservation/db"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(
			map[string]string{"error": err.Error()},
		)
	},
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	client := db.InitMongo()
	app := fiber.New(config)

	var (
		userStore     = db.NewMongoUserStore(client)
		hotelStrore   = db.NewMongoHotelStore(client)
		roomStrore    = db.NewMongoRoomStore(client, hotelStrore)
		bookingStrore = db.NewMongoBookingStore(client)
		store         = &db.Store{
			RoomStore:    roomStrore,
			HotelStore:   hotelStrore,
			UserStore:    userStore,
			BookingStore: bookingStrore,
		}
		// handlers
		userHandler    = api.NewUserHandler(store)
		hotelHandler   = api.NewHotelHandler(store)
		roomHandler    = api.NewRoomHandler(store)
		authHandler    = api.NewAuthHandler(store)
		bookingHandler = api.NewBookingHandler(store)
		auth           = app.Group("/api")
		api            = app.Group("/api/v1", middleware.JWTAuthenticaion(userStore))
	)

	//auth
	auth.Post("/auth", authHandler.HandleAuth)
	// user
	api.Get("/user/:id", userHandler.HandleGetUser)
	api.Get("/user", userHandler.HandleGetUsers)
	api.Post("/user", userHandler.HandlePostUser)
	api.Put("/user/:id", userHandler.HandlePutUser)
	api.Delete("/user/:id", userHandler.HandleDeleteUser)
	// hotel
	api.Get("/hotel", hotelHandler.HandleGetHotels)
	api.Get("/hotel/:id", hotelHandler.HandleGetHotel)
	api.Get("/hotel/:id/rooms", hotelHandler.HandleGetHotelRooms)
	// room
	api.Get("/room", roomHandler.HandleGetRooms)
	// booking
	api.Post("/book/:id", bookingHandler.HandleBookRoom)
	api.Post("/booking/:id", bookingHandler.HandleGetBookings)
	api.Post("/booking", bookingHandler.HandleGetBooking)
	app.Listen(":8080")
}
