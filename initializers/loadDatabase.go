package initializers

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var PostCollection *mongo.Collection

func LoadDatabase() {
	mongoDBURI := os.Getenv("MONGODB_URL")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoDBURI).SetServerAPIOptions(serverAPI)

	mongoClient, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	err = mongoClient.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal("Couldn't connect to mongoDB")
	}

	db := mongoClient.Database("my-database")
	PostCollection = db.Collection("posts")
}
