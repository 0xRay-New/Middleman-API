package database

import (
	"context"
	"mm-api/common"

	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx = context.TODO()
	MongoClient *mongo.Client = nil
)

var (
	MongoDbName = "Users-IndexedData"
	MongoUserCollectionName = "Users"
	MongoIndexedCollectionName = "Indexed-Data"
)

var (
	UserCollection *mongo.Collection
	IndexedCollection *mongo.Collection
)

func InitMongoDb() error {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(common.MongoDbConnectionURL).SetServerAPIOptions(serverAPIOptions)
	Mclient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println(err)
		return err
	}
	MongoClient = Mclient

	UserCollection = MongoClient.Database(MongoDbName).Collection(MongoUserCollectionName)
	IndexedCollection = MongoClient.Database(MongoDbName).Collection(MongoIndexedCollectionName)
	log.Println("started db")
	return nil
}