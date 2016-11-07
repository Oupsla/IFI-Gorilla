package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Username  string        `bson: "username" json: "username"`
	Lastname  string        `bson: "lastname" json: "lastname"`
	Firstname string        `bson: "firstname" json: "firstname"`
	Email     string        `bson: "email" json: "email"`
}

// GetID returns the ID of an Spirit as a string
func (user *User) GetID() string {
	return user.ID.Hex()
}
