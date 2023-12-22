package models

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoomType int

const (
	SingleRoom RoomType = iota + 1
	DoubleRoom
	SeaSideRoom
	DeluxeRoom
)

type Room struct {
	Id      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	HotelId primitive.ObjectID `bson:"hotelId" json:"hotelId"`
	Type    RoomType           `bson:"type" json:"type"`
	Price   float32            `bson:"price" json:"price"`
}

type BookRoomParams struct {
	FromDate   time.Time `bson:"fromDate" json:"fromDate"`
	TillDate   time.Time `bson:"tillDate" json:"tillDate"`
	NumPersons int       `bson:"numPersons" json:"NumPersons"`
}

func (p BookRoomParams) Validate() error {
	if time.Now().After(p.FromDate) || time.Now().After(p.TillDate) {
		return errors.New("cannot book a room in the past")
	}
	return nil
}
