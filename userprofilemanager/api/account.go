package api

import (
	"time"
	"userprofilemanager/core"
	"userprofilemanager/services"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewUserProfile(username string, password string) bool {

	var checkUser []interface{} = core.GetData("username", username, "users")
	if len(checkUser) > 0 {
		return false
	}

	hashedPassword := services.HashString(password)
	// create new user
	user := core.Account{
		Id:        primitive.NewObjectID(),
		UserId:    services.HashString(username + time.Now().String()),
		CreatedAt: time.Now().UTC(),
		Username:  username,
		Password:  hashedPassword,
	}

	userProfile := core.UserProfile{
		Id:        primitive.NewObjectID(),
		UserId:    user.UserId,
		CreatedAt: time.Now().UTC(),
		Birthday:  time.Now().UTC(),
		Address:   "",
		Phone:     "",
	}

	// conver to interface
	insertUser := interface{}(user)
	return core.InsertOne(insertUser, "users") && core.InsertOne(userProfile, "profiles")
}

func CheckUser(username string, password string) (bool, string) {

	var result interface{} = core.GetDataOne("username", username, "users")

	if result == nil {

		return false, "User not found"
	}
	var user core.Account
	bsonBytes, _ := bson.Marshal(result)
	err := bson.Unmarshal(bsonBytes, &user)
	if err != nil {
		return false, "Failed to unmarshal user"
	}
	if user.Password == services.HashString(password) {

		return true, user.UserId
	}

	return false, "Password is incorrect"
}
