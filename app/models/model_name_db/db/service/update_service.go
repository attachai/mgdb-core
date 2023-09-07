package service

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//UpdateService is to handle update function
type UpdateService struct{}

//FindOneAndUpdate is for update document
func (updateservice UpdateService) FindOneAndUpdate(filter interface{}, arrayFilter []interface{}, update interface{}, coll string) (interface{}, error, bool) {
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

//FindOneAndReplace is for replace document
func (updateservice UpdateService) FindOneAndReplace(filter interface{}, update interface{}, coll string) (interface{}, error, bool) {
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
	opt := options.FindOneAndReplaceOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	log.Printf("filter : %v", filter)
	log.Printf("update : %v", update)

	result := collection.FindOneAndReplace(ctx, filter, update, &opt)
	if result.Err() != nil {
		return nil, result.Err(), true
	}

	id := filter.(map[string]interface{})["id"]

	doc := bson.M{}
	decodeErr := result.Decode(&doc)
	log.Printf("result : %v", decodeErr)

	return id, decodeErr, true
}

//UpdateDocuments is to insert many document
func (updateservice UpdateService) UpdateManyDocuments(condition interface{}, data interface{}, coll string) (interface{}, error, bool) {
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

	result, err := collection.UpdateMany(ctx, condition, data)

	id := condition.(map[string]interface{})["id"]
	log.Printf("result names: %v", result)

	return id, err, true
}
