package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserId     primitive.ObjectID `bson:"userId" json:"userId"`
	RoomId     primitive.ObjectID `bson:"roomId" json:"roomId"`
	NumPersons int                `bson:"numPersons" json:"numPersons"`
	FromDate   time.Time          `bson:"fromDate" json:"fromDate"`
	TillDate   time.Time          `bson:"tillDate" json:"tillDate"`
	Canceled   bool               `bson:"canceled" json:"canceled"`
}
