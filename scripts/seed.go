package main

import (
	"context"
	"fmt"
	"hotel-reservation/db"
	"hotel-reservation/models"
)

func main() {

	client := db.Init()

	hotelStore := db.NewMongoHotelStore(client)
	roomStore := db.NewMongoRoomStore(client, hotelStore)
	hotel := models.Hotel{
		Name:     "gho",
		Location: "kerman",
	}
	room := models.Room{
		Type:      models.SingleRoom,
		BasePrice: 99.9,
	}
	insertedHotel, err := hotelStore.Insert(
		context.Background(),
		&hotel,
	)
	if err != nil {
		fmt.Println(err)
	}
	room.HotelId = insertedHotel.Id
	insertedRoom, err := roomStore.InsertRoom(
		context.Background(),
		&room,
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(insertedHotel)
	fmt.Println(insertedRoom)
}
