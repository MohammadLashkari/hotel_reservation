package main

import (
	"hotel-reservation/api"
	"hotel-reservation/config"
	"hotel-reservation/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	client := db.InitMongo()

	var fiberConfig = fiber.Config{
		ErrorHandler: api.ErrorHandler,
	}
	app := fiber.New(fiberConfig)

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
		apiV1          = app.Group("/api/v1", api.JWTAuthenticaion(userStore))
		admin          = apiV1.Group("/admin", api.AdminAuth)
	)

	//auth
	auth.Post("/auth", authHandler.HandleAuth)
	// user
	apiV1.Get("/user/:id", userHandler.HandleGetUser)
	apiV1.Get("/user", userHandler.HandleGetUsers)
	apiV1.Post("/user", userHandler.HandlePostUser)
	apiV1.Put("/user/:id", userHandler.HandlePutUser)
	apiV1.Delete("/user/:id", userHandler.HandleDeleteUser)
	// hotel
	apiV1.Get("/hotel", hotelHandler.HandleGetHotels)
	apiV1.Get("/hotel/:id", hotelHandler.HandleGetHotel)
	apiV1.Get("/hotel/:id/rooms", hotelHandler.HandleGetHotelRooms)
	// room
	apiV1.Get("/room", roomHandler.HandleGetRooms)
	// booking
	apiV1.Post("/book/:id", bookingHandler.HandleBookRoom)
	apiV1.Get("/booking/:id", bookingHandler.HandleGetBooking)
	apiV1.Get("/booking/:id/cancel", bookingHandler.HandleCancelBooking)
	// admin
	admin.Get("/booking", bookingHandler.HandleGetBookings)
	app.Listen(config.Env.HTTP_LISTEN_ADDRESS)
}
