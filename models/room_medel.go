package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RoomType int

const (
	SingleRoom RoomType = iota + 1
	DoubleRoom
	SeaSideRoom
	DeluxeRoom
)

type Room struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	HotelId   primitive.ObjectID `bson:"hotelId" json:"hotelId"`
	Type      RoomType           `bson:"type" json:"type"`
	BasePrice float32            `bson:"basePrice" json:"basePrice"`
	Price     float32            `bson:"Price" json:"Price"`
}
