package api

import (
	"time"
	"userprofilemanager/core"

	"go.mongodb.org/mongo-driver/bson"
)

func UpdateProfile(fullName string, birthday time.Time, address string, email string, phone string, id string) bool {
	var checkUser interface{} = core.GetDataOne("userId", id, "profiles")
	if checkUser == nil {
		return false
	}
	var userProfile core.UserProfile
	bsonBytes, _ := bson.Marshal(checkUser)
	err := bson.Unmarshal(bsonBytes, &userProfile)
	if err != nil {
		return false
	}

	var dataUpdate interface{} = core.UserProfile{
		Id:        userProfile.Id,
		UserId:    userProfile.UserId,
		CreatedAt: userProfile.CreatedAt,
		FullName:  fullName,
		Birthday:  birthday,
		Address:   address,
		Email:     email,
		Phone:     phone,
	}

	return core.UpdateWithID(dataUpdate, "profiles", userProfile.Id)

}

func ViewProfile(id string) core.UserProfile {
	var checkUser interface{} = core.GetDataOne("userId", id, "profiles")
	if checkUser == nil {
		return core.UserProfile{}
	}

	var userProfile core.UserProfile
	bsonBytes, _ := bson.Marshal(checkUser)
	err := bson.Unmarshal(bsonBytes, &userProfile)
	if err != nil {
		return core.UserProfile{}
	}
	return userProfile

}
