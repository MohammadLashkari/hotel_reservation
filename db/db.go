package db

import "go.mongodb.org/mongo-driver/bson/primitive"

const DBNAME = "hotel_reservation"

func ToObjectId(id string) primitive.ObjectID {
	objecId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	return objecId
}
