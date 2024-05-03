package core

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetData(filed string, dataFind string, collectionName string) []interface{} {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(GetEnvrionment("MONGO_URI")))
	if err != nil {
		log.Println(err)
	}

	db := client.Database(
		GetEnvrionment("DB_NAME"),
	)

	condition := bson.M{filed: dataFind}
	cur, err := db.Collection(collectionName).Find(context.Background(), condition)
	if err != nil {
		log.Println(err)
	}

	var result []interface{}
	if err := cur.All(context.Background(), &result); err != nil {
		log.Println(err)
	}

	return result
}

func GetDataOne(filed string, dataFind string, collectionName string) interface{} {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(GetEnvrionment("MONGO_URI")))

	if err != nil {

		log.Println(err)
	}

	db := client.Database(
		GetEnvrionment("DB_NAME"),
	)

	condition := bson.M{filed: dataFind}

	cur := db.Collection(collectionName).FindOne(context.Background(), condition)
	if err != nil {
		log.Println(err)
	}

	var result interface{}
	if err := cur.Decode(&result); err != nil {
		log.Println(err)
	}
	return result
}

func GetDataWithID(id primitive.ObjectID, collectionName string) []interface{} {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(GetEnvrionment("MONGO_URI")))
	if err != nil {
		log.Println(err)
	}

	db := client.Database(
		GetEnvrionment("DB_NAME"),
	)

	condition := bson.M{"_id": id}
	cur, err := db.Collection(collectionName).Find(context.Background(), condition)
	if err != nil {
		log.Println(err)
	}

	var result []interface{}
	if err := cur.All(context.Background(), &result); err != nil {
		log.Println(err)
	}

	return result
}

func InsertOne(dataInsert interface{}, collectionName string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(GetEnvrionment("MONGO_URI")))
	if err != nil {
		log.Println(err)
	}

	db := client.Database(
		GetEnvrionment("DB_NAME"),
	)
	res, err := db.Collection(collectionName).InsertOne(context.Background(), dataInsert)
	if err != nil {
		log.Println(err)
		return false
	}

	log.Printf("Inserted %v", res.InsertedID)

	return true
}

func InsertMany(dataInserts []interface{}, collectionName string) bool {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(GetEnvrionment("MONGO_URI")))
	if err != nil {
		log.Println(err)
	}

	db := client.Database(
		GetEnvrionment("DB_NAME"),
	)

	res, err := db.Collection(collectionName).InsertMany(context.Background(), dataInserts)
	if err != nil {
		log.Println(err)

		return false
	}

	// 2 documents inserted
	log.Printf("%v documents inserted", len(res.InsertedIDs))
	return true
}

func UpdateWithID(dataUpdate interface{}, collectionName string, id primitive.ObjectID) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(GetEnvrionment("MONGO_URI")))
	if err != nil {
		log.Println(err)
	}

	db := client.Database(
		GetEnvrionment("DB_NAME"),
	)

	condition := bson.M{"_id": id}
	update := bson.M{"$set": dataUpdate}
	_, err = db.Collection(collectionName).UpdateOne(context.Background(), condition, update)
	if err != nil {
		log.Println(err)
		return false
	}

	return true

}
