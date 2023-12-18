package models

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost      = 12
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPasswordLen  = 8
)

type User struct {
	Id                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"EncryptedPassword" json:"-"`
}

type InsertUserPrams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UpdateUserPrams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func NewUserFromPrams(prams InsertUserPrams) (*User, error) {
	encPass, err := bcrypt.GenerateFromPassword([]byte(prams.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         prams.FirstName,
		LastName:          prams.LastName,
		Email:             prams.Email,
		EncryptedPassword: string(encPass),
	}, nil
}

func (p UpdateUserPrams) ToBson() bson.M {
	m := bson.M{}
	if len(p.FirstName) > 0 {
		m["firstName"] = p.FirstName
	}
	if len(p.LastName) > 0 {
		m["lastName"] = p.LastName
	}
	return m
}

func (p InsertUserPrams) Validate() map[string]string {
	errors := map[string]string{}
	if len(p.FirstName) < minFirstNameLen {
		errors["firstName"] =
			fmt.Sprintf("firt name should be at least %d characters", minFirstNameLen)

	}
	if len(p.LastName) < minLastNameLen {
		errors["lastName"] =
			fmt.Sprintf("last name should be at least %d characters", minLastNameLen)
	}
	if len(p.Password) < minPasswordLen {
		errors["password"] =
			fmt.Sprintf("password should be at least %d characters", minPasswordLen)
	}
	if !isEmailValid(p.Email) {
		errors["email"] = "invalid email"
	}
	return errors
}

func isEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}
