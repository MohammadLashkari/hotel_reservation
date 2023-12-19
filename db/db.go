package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DBNAME     = "hotel_reservation"
	DBURI      = "mongodb://localhost:27017"
	TESTDBNAME = "hotel_reservation_test"
)

type Store struct {
	UserStore
	HotelStore
	RoomStore
}

func InitMongo() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(DBURI))
	if err != nil {
		log.Fatal(err)
	}
	return client
}
