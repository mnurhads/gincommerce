package database

//
import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/?serverSelectionTimeoutMS=5000&connectTimeoutMS=10000&3t.uriVersion=3&3t.connection.name=Laragon+-+imported+on+Aug+11%2C+2022&3t.alwaysShowAuthDB=true&3t.alwaysShowDBFromUserRole=true"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("failed to connect to mongodb")
		return nil
	}

	// This doesn't work, even if you don't have mongodb running this line still prints.
	fmt.Println("Successfully Connected to the mongodb")
	return client
}

var Client *mongo.Client = DBSet()

func UserData(client *mongo.Client, CollectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("Ecommerce").Collection(CollectionName)
	return collection

}

func ProductData(client *mongo.Client, CollectionName string) *mongo.Collection {
	var productcollection *mongo.Collection = client.Database("Ecommerce").Collection(CollectionName)
	return productcollection
}

func LogData(client *mongo.Client, CollectionName string) *mongo.Collection {
	var LogCollection *mongo.Collection = client.Database("Ecommerce").Collection(CollectionName)
	return LogCollection
}

