package service

import (
	"context"
	"log"
)

//CreateService is to handle create function relation db query
type CreateService struct{}

//InsertOneDocument is to insert one document
func (createservice CreateService) InsertOneDocument(data map[string]interface{}, collection string) (interface{}, error, bool) {
	if !checkCollectionExist(collection) {
		return nil, nil, false
	}
	_, err := Database.Collection(collection).InsertOne(context.TODO(), data)
	id := data["id"]
	return id, err, true
}

//InsertManyDocuments is to insert many document
func (createservice CreateService) InsertManyDocuments(data []interface{}, collection string) (interface{}, error, bool) {
	if !checkCollectionExist(collection) {
		return nil, nil, false
	}
	_, err := Database.Collection(collection).InsertMany(context.TODO(), data)
	var listid []string
	for _, result := range data {
		id := result.(map[string]interface{})["id"]
		listid = append(listid, id.(string))
	}
	return listid, err, true
}

//UpdateDocuments is to insert many document
func (createservice CreateService) UpdateDocuments(condition interface{}, data interface{}, collection string) (interface{}, error, bool) {
	if !checkCollectionExist(collection) {
		return nil, nil, false
	}
	result, err := Database.Collection(collection).UpdateOne(context.TODO(), condition, data)
	id := condition.(map[string]interface{})["id"]
	log.Printf("result names: %v", result)
	return id, err, true
}
