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

func (user *User) GetID() string {
	return user.ID.Hex()
}
