package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/attachai/mgdb-core/app/models/model_name_db/db/service"
	structService "github.com/attachai/mgdb-core/app/structs"

	"fmt"

	"github.com/attachai/mgdb-core/app/models/model_name_db/structs"
	"github.com/attachai/mgdb-core/packages/logging"
	"github.com/attachai/mgdb-core/utils"

	cnst "github.com/attachai/mgdb-core/app/cnst"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// ReadController is for insert logic
type ReadController struct{}

func (auth *ReadController) FindDocumentObj(jsonPost structService.JsonService) (bool, interface{}) {

	byteArray, err := json.Marshal(jsonPost)
	if err != nil {
		logging.Logger(cnst.Error, err)
	}
	//fmt.Println("byteArray=", string(byteArray))
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("{}")))
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(byteArray))

	return auth.FindDocument(c, false)
}

// / count = true ; only count document non get data
func (auth *ReadController) FindDocumentObjCount(jsonPost structService.JsonService, count bool) (bool, interface{}) {

	byteArray, err := json.Marshal(jsonPost)
	if err != nil {
		logging.Logger(cnst.Error, err)
	}
	//fmt.Println("byteArray=", string(byteArray))
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte("{}")))
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(byteArray))
	os.Setenv("collection", jsonPost.Collection)
	return auth.FindDocument(c, count)
}

// FindDocument is for Document insert
func (auth *ReadController) FindDocument(c *gin.Context, count bool) (bool, interface{}) {
	var resultStatus bool
	var resultData interface{}
	var jsonbody structs.Jsonbody
	//Check if jsonbody is not following struck format
	if err := c.ShouldBindJSON(&jsonbody); err != nil {
		fmt.Println(err)
		c.JSON(401, err)
		return resultStatus, resultData
	}
	//Check if data is empty
	if jsonbody.Projection == nil {
		c.JSON(401, gin.H{"error": "'Projection': required field is not set"})
		return resultStatus, resultData
	}

	limit := jsonbody.Limit
	skip := limit * (jsonbody.Offset - 1)

	//Projection
	pro, e := jsonbody.Projection.(map[string]interface{})
	if e {
	}
	projection, date, aggregate := projectionSet(pro, jsonbody.Timezone)

	//Condition
	con, e := jsonbody.Condition.(map[string]interface{})
	if e {
	}
	condition := utils.ConvertOperators(con).(map[string]interface{})

	//arrayFilter
	// arr, err := jsonbody.ArrayFilter.(map[string]interface{})
	// if err {}
	// arrayFilter := mapString(arr).(map[string]interface{})
	// fmt.Println(arrayFilter)

	//find with aggregate
	if aggregate {
		pipeline := []bson.M{}

		condition = bson.M{"$match": condition}
		projection = bson.M{"$project": projection}
		sort := bson.M{"$sort": jsonbody.Sort}
		limits := bson.M{"$limit": limit}
		skips := bson.M{"$skip": skip}
		addFields := bson.M{"$addFields": date}

		if len(date) != 0 && len(con) != 0 {
			pipeline = []bson.M{condition, projection, addFields, sort, skips, limits}
		} else if len(date) == 0 && len(con) != 0 {
			pipeline = []bson.M{condition, projection, sort, skips, limits}
		} else if len(date) != 0 && len(con) == 0 {
			pipeline = []bson.M{projection, addFields, sort, skips, limits}
		} else {
			pipeline = []bson.M{projection, sort, skips, limits}
		}

		fmt.Println("Aggregate(): ", pipeline)

		userservice := service.ReadService{}
		result, err, collection := userservice.AggregateDocument(pipeline, jsonbody.Collection)
		if err != nil || !collection {
			if !collection {
				c.JSON(500, gin.H{"statusCode": cnst.HTTP500, "message": "The following item haven't gotten", "errors": "Collection not found!"})
			} else {
				c.JSON(500, gin.H{"statusCode": cnst.HTTP500, "message": "The following item haven't gotten", "errors": err.Error()})
			}
		} else {
			// c.JSON(200, gin.H{"statusCode": setting.AppSetting.HTTP200, "message": "The following items have gotten successfully", "results": result})
			resultStatus = true
			resultData = result

		}

		//find document
	} else {
		filter := condition
		if len(date) != 0 {
			projection["last_updated"] = date
		}
		fmt.Println("Find() filter: ", filter)
		fmt.Println("Find() projection: ", projection)

		userservice := service.ReadService{}
		if count == true {
			result, _, collection := userservice.FindDocumentCount(filter, projection, jsonbody.Collection, jsonbody.Sort, int64(limit), int64(skip))
			if collection == true {
				return true, result
			} else {
				return false, nil
			}
		}
		result, err, collection := userservice.FindDocument(filter, projection, jsonbody.Collection, jsonbody.Sort, int64(limit), int64(skip))
		if err != nil || !collection {
			if !collection {
				c.JSON(500, gin.H{"statusCode": cnst.HTTP500, "message": "The following item haven't gotten", "errors": "Collection not found!"})
			} else {
				c.JSON(500, gin.H{"statusCode": cnst.HTTP500, "message": "The following item haven't gotten", "errors": err.Error()})
			}
		} else {
			// c.JSON(200, gin.H{"statusCode": setting.AppSetting.HTTP200, "message": "The following items have gotten successfully", "results": result})
			resultStatus = true
			resultData = result

		}
	}
	return resultStatus, resultData
}

// projectionSet is for setup projection data
func projectionSet(p interface{}, timezone string) (bson.M, bson.M, bool) {
	projection := bson.M{}
	date := bson.M{}
	dateList := bson.M{}
	aggregate := false
	project, err := p.(map[string]interface{})
	if err {
	}
	//set default 0 for document's id (MongoDB)
	if len(project) != 0 {
		projection["_id"] = 0
	}
	//set projection data
	for key, result := range project {
		if result == "date" {
			date = bson.M{
				"$dateToString": bson.M{
					"date":     "$" + key,
					"timezone": timezone,
					"format":   "%Y-%m-%dT%H:%M:%S.%L%z",
				},
			}
			dateList[key] = date
		} else {
			projection[key] = result
		}

		if result == 0.0 {
			if key == "_id" {
				continue
			}
			aggregate = true
		}
	}
	return projection, dateList, aggregate
}

// //conditionSet is for setup condition data
// func conditionSet(c interface{}) (bson.M) {
// 	//if the argument is not a map, ignore it
//     condition, ok := c.(map[string]interface{})

//     if !ok {
//         return nil
//     }
// 	// fmt.Println(condition)
//     for _, v := range condition {
// 		// fmt.Println(k)

//         // key match
// 		// if k != utils.MapStr(k) {
// 			// condition[utils.MapStr(k)] = v
// 		// 	delete(condition, k)
// 		// }

//         // if the value is a map, search recursively
//         if m, ok := v.(map[string]interface{}); ok {
//             conditionSet(m)
//         }
//         // if the value is an array, search recursively
//         // from each element
//         if va, ok := v.([]interface{}); ok {
//             for _, a := range va {
//                 conditionSet(a)
//             }
//         }
//     }

//     return condition
// }

// jsonCondition := `
// 					{
// 						"id": "5bf142459b72e12b2b1b2cd1",
// 						"$or": [
// 							{
// 								"sizes": {
// 									"$elemMatch": {
// 										"id": "5bf142459b72e12b2b1b2af2",
// 										"quantity": {
// 											"$gt": 0
// 										}
// 									},
// 									"$and": [
// 										{
// 											"sizes.uk": "7"
// 										},
// 										{
// 											"sizes.quantity": 0
// 										}
// 									]
// 								}
// 							},
// 							{
// 								"colors": {
// 									"$all": [
// 										"Black",
// 										"White"
// 									]
// 								}
// 							},
// 							{
// 								"sizes": {
// 									"$in": [
// 										"M",
// 										"L"
// 									]
// 								}
// 							}
// 						]
// 					}
// 				`
// 	var condition map[string]interface{}
// 	json.Unmarshal([]byte(jsonCondition), &condition)
