package db

import (
	"context"
	"hotel-reservation/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const roomColl = "rooms"

type RoomStore interface {
	Insert(context.Context, *models.Room) (*models.Room, error)
}

type MongoRoomStore struct {
	client     *mongo.Client
	collection *mongo.Collection
	HotelStore
}

func NewMongoRoomStore(client *mongo.Client, hotelStore HotelStore) *MongoRoomStore {
	return &MongoRoomStore{
		client:     client,
		collection: client.Database(DBNAME).Collection(roomColl),
		HotelStore: hotelStore,
	}
}

func (s *MongoRoomStore) InsertRoom(ctx context.Context, room *models.Room) (*models.Room, error) {
	res, err := s.collection.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}
	room.Id = res.InsertedID.(primitive.ObjectID)

	// update hotel
	filter := bson.M{"_id": room.HotelId}
	update := bson.M{"$push": bson.M{"rooms": room.Id}}
	s.HotelStore.Update(ctx, filter, update)
	return room, nil
}
