package main

import (
	"context"
	"fmt"
	"hotel-reservation/config"
	"hotel-reservation/db"
	"hotel-reservation/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {

	client := db.InitMongo()
	ctx := context.Background()
	client.Database(config.Env.MONGO_DB_NAME).Drop(ctx)

	hotelStore := db.NewMongoHotelStore(client)
	roomStore := db.NewMongoRoomStore(client, hotelStore)
	hotel := models.Hotel{
		Name:     "gho",
		Location: "kerman",
		Rating:   4,
		Rooms:    []primitive.ObjectID{},
	}
	rooms := []models.Room{
		{
			Type:  models.DeluxeRoom,
			Price: 189.9,
		},
		{
			Type:  models.SingleRoom,
			Price: 99.9,
		},
		{
			Type:  models.SeaSideRoom,
			Price: 120.0,
		},
	}
	insertedHotel, err := hotelStore.Insert(ctx, &hotel)
	if err != nil {
		fmt.Println(err)
	}
	for _, room := range rooms {
		room.HotelId = insertedHotel.Id
		roomStore.Insert(ctx, &room)
	}
}
