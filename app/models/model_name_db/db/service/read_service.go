package service

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ReadService is to handle create function relation db query
type ReadService struct{}

//AggregateDocument is to read documents
func (readservice ReadService) AggregateDocument(data []bson.M, collection string) (interface{}, error, bool) {
	if !checkCollectionExist(collection) {
		return nil, nil, false
	}
	result := []bson.M{}
	opts := options.Aggregate()
	cur, err := Database.Collection(collection).Aggregate(context.TODO(), data, opts)
	if err != nil {
		return result, err, true
	}
	//Map result to slice
	for cur.Next(context.TODO()) {
		t := bson.M{}
		err := cur.Decode(&t)
		if err != nil {
			return result, err, true
		}
		result = append(result, t)
	}
	cur.Close(context.TODO())
	if len(result) == 0 {
		return result, mongo.ErrNoDocuments, true
	}

	return result, err, true
}

func (readservice ReadService) FindDocument(filter bson.M, projection bson.M, collection string, sort interface{}, limit int64, skip int64) (interface{}, error, bool) {
	if !checkCollectionExist(collection) {
		return nil, nil, false
	}

	result := []bson.M{}
	opts := options.Find()
	opts.SetProjection(projection)
	opts.SetSort(sort)
	opts.SetLimit(limit)
	opts.SetSkip(skip)
	cur, err := Database.Collection(collection).Find(context.TODO(), filter, opts)
	if err != nil {
		return result, err, true
	}
	//Map result to slice
	for cur.Next(context.TODO()) {
		t := bson.M{}
		err := cur.Decode(&t)
		if err != nil {
			return result, err, true
		}
		result = append(result, t)
	}
	cur.Close(context.TODO())
	if len(result) == 0 {
		return result, mongo.ErrNoDocuments, true
	}

	return result, err, true
}

func (readservice ReadService) FindDocumentCount(filter bson.M, projection bson.M, collection string, sort interface{}, limit int64, skip int64) (interface{}, error, bool) {
	if !checkCollectionExist(collection) {
		return nil, nil, false
	}
	// fmt.Println("limit=  = ", limit)
	// result := []bson.M{}

	cotp := options.Count()
	if limit != 0 {
		cotp.SetLimit(limit)
	}
	cotp.SetSkip(skip)

	no, err := Database.Collection(collection).CountDocuments(context.TODO(), filter, cotp)
	if err != nil {
		fmt.Println("count search error : ", err.Error())
		return nil, err, false
	}
	// fmt.Println("read = = == = count = ", no)

	return no, err, true
}
