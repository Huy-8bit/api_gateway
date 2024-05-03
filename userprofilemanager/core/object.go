package core

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Environment struct {
	MongoURI string
	DBName   string
}

type Account struct {
	Id        primitive.ObjectID `bson:"_id"`
	UserId    string             `bson:"userId"`
	CreatedAt time.Time          `bson:"createdAt"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
}

type UserProfile struct {
	Id        primitive.ObjectID `bson:"_id"`
	UserId    string             `bson:"userId"`
	FullName  string             `bson:"fullName"`
	CreatedAt time.Time          `bson:"createdAt"`
	Birthday  time.Time          `bson:"birthday"`
	Address   string             `bson:"address"`
	Email     string             `bson:"email"`
	Phone     string             `bson:"phone"`
}
