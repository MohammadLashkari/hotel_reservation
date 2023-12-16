package db

import (
	"context"
	"hotel-reservation/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userColl = "users"

type UserStore interface {
	GetUserById(context.Context, string) (*models.User, error)
}

type MongoUserStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client:     client,
		collection: client.Database(DBNAME).Collection(userColl),
	}
}

func (s *MongoUserStore) GetUserById(ctx context.Context, id string) (*models.User, error) {
	objecId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var user models.User
	if err := s.collection.FindOne(ctx, bson.M{"_id": objecId}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
