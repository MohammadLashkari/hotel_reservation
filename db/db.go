package db

import (
	"context"
	"hotel-reservation/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	UserStore
	HotelStore
	RoomStore
	BookingStore
}

func InitMongo() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Env.MONGO_DB_URL))
	if err != nil {
		log.Fatalf("unable to connect mongodb : %e", err)
	}
	return client
}
