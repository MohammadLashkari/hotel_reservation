package db

import (
	"context"
	"hotel-reservation/config"
	"hotel-reservation/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const bookingColl = "booking"

type BookingStore interface {
	GetAll(context.Context, bson.M) ([]*models.Booking, error)
	GetById(context.Context, string) (*models.Booking, error)
	Insert(context.Context, *models.Booking) (*models.Booking, error)
	Update(context.Context, string, bson.M) error
}

type MongoBookingStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoBookingStore(client *mongo.Client) *MongoBookingStore {
	return &MongoBookingStore{
		client:     client,
		collection: client.Database(config.Env.MONGO_DB_NAME).Collection(bookingColl),
	}
}

func (s *MongoBookingStore) GetAll(ctx context.Context, filter bson.M) ([]*models.Booking, error) {
	res, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var bookedRooms []*models.Booking
	if err := res.All(ctx, &bookedRooms); err != nil {
		return nil, err
	}
	return bookedRooms, nil
}

func (s *MongoBookingStore) Update(ctx context.Context, id string, newValue bson.M) error {
	objecId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	update := bson.M{"$set": newValue}
	_, err = s.collection.UpdateByID(ctx, objecId, update)
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoBookingStore) GetById(ctx context.Context, id string) (*models.Booking, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var (
		booking    models.Booking
		filterById = bson.M{"_id": objectId}
	)
	if err := s.collection.FindOne(ctx, filterById).Decode(&booking); err != nil {
		return nil, err
	}
	return &booking, nil
}
func (s *MongoBookingStore) Insert(ctx context.Context, booking *models.Booking) (*models.Booking, error) {
	resp, err := s.collection.InsertOne(ctx, booking)
	if err != nil {
		return nil, err
	}
	booking.Id = resp.InsertedID.(primitive.ObjectID)

	return booking, nil

}
