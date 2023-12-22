package db

import (
	"context"
	"hotel-reservation/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const bookingColl = "booking"

type BookingStore interface {
	GetAll(context.Context) ([]*models.Booking, error)
	Insert(context.Context, *models.Booking) (*models.Booking, error)
}

type MongoBookingStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoBookingStore(client *mongo.Client) *MongoBookingStore {
	return &MongoBookingStore{
		client:     client,
		collection: client.Database(DBNAME).Collection(bookingColl),
	}
}

func (s *MongoBookingStore) GetAll(ctx context.Context) ([]*models.Booking, error) {
	res, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var bookedRooms []*models.Booking
	if err := res.All(ctx, &bookedRooms); err != nil {
		return nil, err
	}
	return bookedRooms, nil
}
func (s *MongoBookingStore) Insert(ctx context.Context, booking *models.Booking) (*models.Booking, error) {
	resp, err := s.collection.InsertOne(ctx, booking)
	if err != nil {
		return nil, err
	}
	booking.Id = resp.InsertedID.(primitive.ObjectID)

	return booking, nil

}
