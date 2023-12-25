package db

import (
	"context"
	"hotel-reservation/config"
	"hotel-reservation/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const hotelColl = "hotels"

type HotelStore interface {
	Insert(context.Context, *models.Hotel) (*models.Hotel, error)
	GetAll(context.Context, bson.M) ([]*models.Hotel, error)
	GetById(context.Context, string) (*models.Hotel, error)
	Update(context.Context, bson.M, bson.M) error
}

type MongoHotelStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client) *MongoHotelStore {
	return &MongoHotelStore{
		client:     client,
		collection: client.Database(config.Env.MONGO_DB_NAME).Collection(hotelColl),
	}
}

func (s *MongoHotelStore) GetAll(ctx context.Context, filter bson.M) ([]*models.Hotel, error) {
	res, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var hotels []*models.Hotel
	if err := res.All(ctx, &hotels); err != nil {
		return nil, err
	}
	return hotels, nil
}

func (s *MongoHotelStore) GetById(ctx context.Context, id string) (*models.Hotel, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var (
		hotel      models.Hotel
		filterById = bson.M{"_id": objectId}
	)
	if err := s.collection.FindOne(ctx, filterById).Decode(&hotel); err != nil {
		return nil, err
	}
	return &hotel, nil
}

func (s *MongoHotelStore) Insert(ctx context.Context, hotel *models.Hotel) (*models.Hotel, error) {
	res, err := s.collection.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}
	hotel.Id = res.InsertedID.(primitive.ObjectID)
	return hotel, nil
}

func (s *MongoHotelStore) Update(ctx context.Context, filter, update bson.M) error {
	_, err := s.collection.UpdateOne(ctx, filter, update)
	return err
}
