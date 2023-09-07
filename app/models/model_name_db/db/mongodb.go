package db

import (
	"context"

	"github.com/attachai/core/app/models/model_name_db/db/service"
	"github.com/attachai/core/utils"

	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connect is for get mongo driver connection
func Connect() {

	connectionString := utils.ViperEnvVariable("Host")
	dbName := utils.ViperEnvVariable("Name")
	userDb := utils.ViperEnvVariable("User")
	passDb := utils.ViperEnvVariable("Password")
	// Database Config
	credential := options.Credential{
		Username: userDb,
		Password: passDb,
	}
	// clientOptions := options.Client().ApplyURI(connectionString)
	clientOptions := options.Client().ApplyURI(connectionString).SetAuth(credential)
	client, err := mongo.NewClient(clientOptions)

	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	//Cancel context to avoid memory leak
	defer cancel()

	// Ping our db connection
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	//Connect to the database
	db := client.Database(dbName)
	service.DBConnection(db)

	return
}

func CollectionList(db *mongo.Database) []string {
	//Check collection is not empty
	filter := bson.D{{}}
	names, err := db.ListCollectionNames(context.TODO(), filter)
	if err != nil {
		// Handle error
		log.Printf("Failed to get coll names: %v", err)
		return nil
	}
	return names
}
