package models

type User struct {
	Id        string `bson:"_id" json:"id,omitempty"`
	FirstName string `bson:"firstName" json:"firtName"`
	LastName  string `bson:"lastName" json:"lastName"`
}
