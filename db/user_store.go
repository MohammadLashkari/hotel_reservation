package db

import (
	"context"
	"fmt"
	"hotel-reservation/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userColl = "users"

type Dropper interface {
	Drop(context.Context) error
}

type UserStore interface {
	Dropper
	GetById(context.Context, string) (*models.User, error)
	GetByEmail(context.Context, string) (*models.User, error)
	GetAll(context.Context, bson.M) ([]*models.User, error)
	Insert(context.Context, *models.User) (*models.User, error)
	Delete(context.Context, string) error
	Update(context.Context, string, models.UpdateUserPrams) error
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

func (s *MongoUserStore) GetById(ctx context.Context, id string) (*models.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var (
		user       models.User
		filterById = bson.M{"_id": objectId}
	)
	if err := s.collection.FindOne(ctx, filterById).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *MongoUserStore) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var (
		user          models.User
		filterByEmail = bson.M{"email": email}
	)
	if err := s.collection.FindOne(ctx, filterByEmail).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *MongoUserStore) GetAll(ctx context.Context, filter bson.M) ([]*models.User, error) {
	cur, err := s.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var users []*models.User
	if err := cur.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (s *MongoUserStore) Insert(ctx context.Context, user *models.User) (*models.User, error) {
	res, err := s.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.Id = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (s *MongoUserStore) Delete(ctx context.Context, id string) error {
	objecId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	// ToDo: maybe its a good idea to handle if we did not delete any user
	_, err = s.collection.DeleteOne(ctx, bson.M{"_id": objecId})
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoUserStore) Update(ctx context.Context, id string, params models.UpdateUserPrams) error {
	objecId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	update := bson.M{"$set": params}
	_, err = s.collection.UpdateByID(ctx, objecId, update)
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoUserStore) Drop(ctx context.Context) error {
	fmt.Println("---droping")
	return s.collection.Drop(ctx)
}
