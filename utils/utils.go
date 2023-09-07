package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"time"

	oid "github.com/coolbed/mgo-oid"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// const charset = "abcdefghijklmnopqrstuvwxyz" +
// 	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const charset = "0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomCode(length int) string {
	return StringWithCharset(length, charset)
}

//GenerateID function is for Generate Document id
func GenerateID(prefix string) string {
	objectID := oid.NewOID()
	// fmt.Println("object id:", objectID.String())
	// fmt.Println("object timestamp", objectID.Timestamp())
	return prefix + objectID.String()
}

func GenerateUUID() string {
	objectID := uuid.New()
	// fmt.Println("object id:", objectID.String())
	// fmt.Println("object timestamp", objectID.Timestamp())
	return objectID.String()
}

type StrSlice struct {
	Str []string
}

//convertToSlice is to convert string to array
func ConvertToSlice(s string) []string {
	var str []string
	err := json.Unmarshal([]byte(s), &str)
	if err != nil {
		log.Fatal(err)
	}
	return str
}

//ConvertOperators is to convert Query Operators
func ConvertOperators(data interface{}) interface{} {
	if reflect.ValueOf(data).Kind() == reflect.Slice {
		d := reflect.ValueOf(data)
		tmpData := make([]interface{}, d.Len())
		returnSlice := make([]interface{}, d.Len())
		for i := 0; i < d.Len(); i++ {
			tmpData[i] = d.Index(i).Interface()
		}
		for i, v := range tmpData {
			returnSlice[i] = ConvertOperators(v)
		}
		return returnSlice
	} else if reflect.ValueOf(data).Kind() == reflect.Map {
		d := reflect.ValueOf(data)
		tmpData := make(map[string]interface{})
		for _, k := range d.MapKeys() {
			typeOfValue := reflect.TypeOf(d.MapIndex(k).Interface()).Kind()
			if typeOfValue == reflect.Map || typeOfValue == reflect.Slice {
				tmpData[MapOperators(k.String())] = ConvertOperators(d.MapIndex(k).Interface())
			} else {
				tmpData[MapOperators(k.String())] = d.MapIndex(k).Interface()
			}
		}
		return tmpData
	}
	return data
}

//MapOperators for set Query Operators
func MapOperators(str string) string {
	switch str {
	//---MongoDB Query Operators---//
	//comparison
	case "eq":
		return "$eq"
	case "gt":
		return "$gt"
	case "gte":
		return "$gte"
	case "in":
		return "$in"
	case "lt":
		return "$lt"
	case "lte":
		return "$lte"
	case "ne":
		return "$ne"
	case "nin":
		return "$nin"
	//logical
	case "and":
		return "$and"
	case "not":
		return "$not"
	case "nor":
		return "$nor"
	case "or":
		return "$or"
	//element
	case "exists":
		return "$exists"
	case "type":
		return "$type"
	//evaluation
	case "expr":
		return "$expr"
	case "jsonSchema":
		return "$jsonSchema"
	case "mod":
		return "$mod"
	case "regex":
		return "$regex"
	case "regexNo":
		return "regex"
	case "text":
		return "$text"
	case "where":
		return "$where"
	//geospatial
	case "geoIntersects":
		return "$geoIntersects"
	case "geoWithin":
		return "$geoWithin"
	case "near":
		return "$near"
	case "nearSphere":
		return "$nearSphere"
	//array
	case "all":
		return "$all"
	case "elemMatch":
		return "$elemMatch"
	case "size":
		return "$size"
	//bitwise
	case "bitsAllClear":
		return "$bitsAllClear"
	case "bitsAllSet":
		return "$bitsAllSet"
	case "bitsAnyClear":
		return "$bitsAnyClear"
	case "bitsAnySet":
		return "$bitsAnySet"

	//---MongoDB -Index of Expression Operators---//
	case "abs":
		return "$abs"
	case "accumulator":
		return "$accumulator"
	case "acos":
		return "$acos"
	case "acosh":
		return "$acosh"
	case "add":
		return "$add"
	case "addToSet":
		return "$addToSet"
	case "allElementsTrue":
		return "$allElementsTrue"
	// case "and":
	// 	return "$and"
	case "anyElementTrue":
		return "$anyElementTrue"
	case "arrayElemAt":
		return "$arrayElemAt"
	case "arrayToObject":
		return "$arrayToObject"
	case "asin":
		return "$asin"
	case "asinh":
		return "$asinh"
	case "atan":
		return "$atan"
	case "atan2":
		return "$atan2"
	case "atanh":
		return "$atanh"
	case "avg":
		return "$avg"
	case "binarySize":
		return "$binarySize"
	case "bsonSize":
		return "$bsonSize"
	case "ceil":
		return "$ceil"
	case "cmp":
		return "$cmp"
	case "concat":
		return "$concat"
	case "concatArrays":
		return "$concatArrays"
	case "cond":
		return "$cond"
	case "convert":
		return "$convert"
	case "cos":
		return "$cos"
	case "dateFromParts":
		return "$dateFromParts"
	case "dateFromString":
		return "$dateFromString"
	case "dateToParts":
		return "$dateToParts"
	case "dateToString":
		return "$dateToString"
	case "dayOfMonth":
		return "$dayOfMonth"
	case "dayOfWeek":
		return "$dayOfWeek"
	case "dayOfYear":
		return "$dayOfYear"
	case "degreesToRadians":
		return "$degreesToRadians"
	case "divide":
		return "$divide"
	// case "eq":
	// 	return "$eq"
	case "exp":
		return "$exp"
	case "filter":
		return "$filter"
	// case "first":
	// 	return "$first"  (array)
	// case "first":
	// 	return "$first"  (accumulator)
	case "floor":
		return "$floor"
	case "function":
		return "$function"
	// case "gt":
	// 	return "$gt"
	// case "gte":
	// 	return "$gte"
	case "hour":
		return "$hour"
	case "ifNull":
		return "$ifNull"
	// case "in":
	// 	return "$in"
	case "indexOfArray":
		return "$indexOfArray"
	case "indexOfBytes":
		return "$indexOfBytes"
	case "indexOfCP":
		return "$indexOfCP"
	case "isArray":
		return "$isArray"
	case "isNumber":
		return "$isNumber"
	case "isoDayOfWeek":
		return "$isoDayOfWeek"
	case "isoWeek":
		return "$isoWeek"
	case "isoWeekYear":
		return "$isoWeekYear"
	// case "last":
	// 	return "$last"  (array)
	// case "last":
	// 	return "$last"  (accumulator)
	case "let":
		return "$let"
	case "literal":
		return "$literal"
	case "ln":
		return "$ln"
	case "log":
		return "$log"
	case "log10":
		return "$log10"
	// case "lt":
	// 	return "$lt"
	// case "lte":
	// 	return "$lte"
	case "ltrim":
		return "$ltrim"
	case "map":
		return "$map"
	case "max":
		return "$max"
	case "mergeObjects":
		return "$mergeObjects"
	case "meta":
		return "$meta"
	case "millisecond":
		return "$millisecond"
	case "min":
		return "$min"
	case "minute":
		return "$minute"
	// case "mod":
	// 	return "$mod"
	case "month":
		return "$month"
	case "multiply":
		return "$multiply"
	// case "ne":
	// 	return "$ne"
	// case "not":
	// 	return "$not"
	case "objectToArray":
		return "$objectToArray"
	// case "or":
	// 	return "$or"
	case "pow":
		return "$pow"
	case "push":
		return "$push"
	case "radiansToDegrees":
		return "$radiansToDegrees"
	case "range":
		return "$range"
	case "reduce":
		return "$reduce"
	case "regexFind":
		return "$regexFind"
	case "regexFindAll":
		return "$regexFindAll"
	case "regexMatch":
		return "$regexMatch"
	case "replaceOne":
		return "$replaceOne"
	case "replaceAll":
		return "$replaceAll"
	case "reverseArray":
		return "$reverseArray"
	case "round":
		return "$round"
	case "rtrim":
		return "$rtrim"
	case "second":
		return "$second"
	case "setDifference":
		return "$setDifference"
	case "setEquals":
		return "$setEquals"
	case "setIntersection":
		return "$setIntersection"
	case "setIsSubset":
		return "$setIsSubset"
	case "setUnion":
		return "$setUnion"
	case "sin":
		return "$sin"
	// case "size":
	// 	return "$size"
	case "slice":
		return "$slice"
	case "split":
		return "$split"
	case "sqrt":
		return "$sqrt"
	case "stdDevPop":
		return "$stdDevPop"
	case "stdDevSamp":
		return "$stdDevSamp"
	case "strLenBytes":
		return "$strLenBytes"
	case "strLenCP":
		return "$strLenCP"
	case "strcasecmp":
		return "$strcasecmp"
	case "substr":
		return "$substr"
	case "substrBytes":
		return "$substrBytes"
	case "substrCP":
		return "$substrCP"
	case "subtract":
		return "$subtract"
	case "sum":
		return "$sum"
	case "switch":
		return "$switch"
	case "tan":
		return "$tan"
	case "toBool":
		return "$toBool"
	case "toDate":
		return "$toDate"
	case "toDecimal":
		return "$toDecimal"
	case "toDouble":
		return "$toDouble"
	case "toInt":
		return "$toInt"
	case "toLong":
		return "$toLong"
	case "toLower":
		return "$toLower"
	case "toObjectId":
		return "$toObjectId"
	case "toString":
		return "$toString"
	case "toUpper":
		return "$toUpper"
	case "trim":
		return "$trim"
	case "trunc":
		return "$trunc"
	// case "type":
	// 	return "$type"
	case "week":
		return "$week"
	case "year":
		return "$year"
	case "zip":
		return "$zip"

	//---MongoDB Update Operators---//
	//Fields
	case "currentDate":
		return "$currentDate"
	case "inc":
		return "$inc"
	// case "min":
	// 	return "$min"
	// case "max":
	// 	return "$max"
	case "mul":
		return "$mul"
	case "rename":
		return "$rename"
	case "set":
		return "$set"
	case "setOnInsert":
		return "$setOnInsert"
	case "unset":
		return "$unset"
	//Array
	// case "addToSet":
	// 	return "$addToSet"
	case "pop":
		return "$pop"
	case "pull":
		return "$pull"
	// case "push":
	// 	return "$push"
	case "pullAll":
		return "$pullAll"
	//Modifiers
	case "each":
		return "$each"
	case "position":
		return "$position"
	// case "slice":
	// 	return "$slice"
	case "sort":
		return "$sort"
	//Bitwise
	case "bit":
		return "$bit"

	}
	return str
}

func GetEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func GetUrl(ipAddress string, apiGroup string, apiEndpoint string) string {
	var url string
	if ipAddress != "" {
		url = ipAddress
		if apiGroup != "" {
			url = url + apiGroup
		}
		if apiEndpoint != "" {
			url = url + apiEndpoint
		}
	}

	return url
}

// use viper package to read .env file
// return the value of the key
func ViperEnvVariable(key string) string {

	// SetConfigFile explicitly defines the path, name and extension of the config file.
	// Viper will use this and not check any of the config paths.
	// .env - It will search for the .env file in the current directory
	// viper.SetConfigFile(".env")

	// Find and read the config file
	// err := viper.ReadInConfig()

	// if err != nil {
	// 	log.Fatalf("Error while reading config file %s", err)
	// }

	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error
	value := viper.Get(key)

	return fmt.Sprint(value)
}

func ContainInSlice(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func NewBoolean(state bool) *bool {
	if state == true {
		state = true
	} else if state == false {
		state = false
	}
	return &state
}