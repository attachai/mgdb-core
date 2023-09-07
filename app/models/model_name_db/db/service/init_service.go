package service

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//InitService is to handle create function relation db query
type InitService struct{}

var Database *mongo.Database

//DBConnection ..
func DBConnection(c *mongo.Database) {
	Database = c
}

//checkCollectionExist is to check collection exist or not
func checkCollectionExist(collection string) bool {
	filter := bson.D{{}}
	collectionList, err := Database.ListCollectionNames(context.TODO(), filter)
	if err != nil {
		// Handle error
		log.Printf("Failed to get coll names: %v", err)
		return false
	}
	for _, name := range collectionList {
		if name == collection {
			return true
		}
	}
	return false
}
