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
	GetAll(context.Context, bson.M) ([]*models.Room, error)
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
func (s *MongoRoomStore) GetAll(ctx context.Context, filter bson.M) ([]*models.Room, error) {
	res, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var rooms []*models.Room
	if err := res.All(ctx, &rooms); err != nil {
		return nil, err
	}
	return rooms, nil
}

func (s *MongoRoomStore) Insert(ctx context.Context, room *models.Room) (*models.Room, error) {
	res, err := s.collection.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}
	room.Id = res.InsertedID.(primitive.ObjectID)

	// update hotel rooms
	filter := bson.M{"_id": room.HotelId}
	update := bson.M{"$push": bson.M{"rooms": room.Id}}
	if err := s.HotelStore.Update(ctx, filter, update); err != nil {
		return nil, err
	}
	return room, nil
}
