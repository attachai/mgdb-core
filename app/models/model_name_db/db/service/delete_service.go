package service

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DeleteService is to handle Delete function
type DeleteService struct{}

//FindOneAndDelete is for Delete document
func (deleteservice DeleteService) FindOneAndDelete(filter interface{}, coll string) (interface{}, error, bool) {
	//check Collection is exist
	if !checkCollectionExist(coll) {
		return nil, nil, false
	}

	//create the context
	exp := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), exp)
	defer cancel()

	//select the collection
	collection := Database.Collection(coll)

	log.Printf("filter : %v", filter)

	result := collection.FindOneAndDelete(ctx, filter)
	if result.Err() != nil {
		return nil, result.Err(), true
	}

	id := filter.(map[string]interface{})["id"]

	doc := bson.M{}
	decodeErr := result.Decode(&doc)
	log.Printf("result : %v", decodeErr)

	return id, decodeErr, true
}

//FindOneAndUpdate is for update document
func (deleteservice DeleteService) FindOneAndUpdate(filter interface{}, arrayFilter []interface{}, update interface{}, coll string) (interface{}, error, bool) {
	//check Collection is exist
	if !checkCollectionExist(coll) {
		return nil, nil, false
	}

	//create the context
	exp := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), exp)
	defer cancel()

	//select the collection
	collection := Database.Collection(coll)

	//create an instance of an options and set the desired options
	upsert := true
	after := options.After
	arrayFilters := options.ArrayFilters{
		Filters: arrayFilter,
	}

	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
		ArrayFilters:   &arrayFilters,
	}

	log.Printf("arrayFilters : %v", arrayFilter)

	log.Printf("filter : %v", filter)
	log.Printf("update : %v", update)

	result := collection.FindOneAndUpdate(ctx, filter, update, &opt)
	if result.Err() != nil {
		return nil, result.Err(), true
	}

	id := filter.(map[string]interface{})["id"]

	doc := bson.M{}
	decodeErr := result.Decode(&doc)
	log.Printf("result : %v", decodeErr)

	return id, decodeErr, true
}

//DeleteMany is for Delete document
func (deleteservice DeleteService) DeleteMany(filter interface{}, coll string) (interface{}, error, bool) {
	//check Collection is exist
	if !checkCollectionExist(coll) {
		return nil, nil, false
	}

	//create the context
	exp := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), exp)
	defer cancel()

	//select the collection
	collection := Database.Collection(coll)

	log.Printf("filter : %v", filter)

	result, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		return nil, err, true
	}

	id := filter.(map[string]interface{})["id"]

	log.Printf("result : %v", result)
	log.Printf("err : %v", err)

	return id, err, true
}

//UpdateDocuments is to insert many document
func (deleteservice DeleteService) DeleteManyWithFilter(filter interface{}, data interface{}, coll string) (interface{}, error, bool) {
	//check Collection is exist
	if !checkCollectionExist(coll) {
		return nil, nil, false
	}

	//create the context
	exp := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), exp)
	defer cancel()

	//select the collection
	collection := Database.Collection(coll)

	result, err := collection.UpdateMany(ctx, filter, data)

	id := filter.(map[string]interface{})["id"]
	log.Printf("result names: %v", result)

	return id, err, true
}
