package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//GetMongoDbConnection get connection of mongodb
func GetMongoDbConnection() (*mongo.Client, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://dsoinventory:lTdCql8EjuMUr8G5n79M4LsiLHG9lCRab5IkCVKgygomlqrnVJRIGeWQi9e1hpl6l6ElMewQnRt5NXFIGlWKIA==@dsoinventory.mongo.cosmos.azure.com:10255/?ssl=true&replicaSet=globaldb&retrywrites=false&maxIdleTimeMS=120000&appName=@dsoinventory@").SetDirect(true))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}

func getMongoDbCollection(DbName string, CollectionName string) (*mongo.Collection, error) {
	client, err := GetMongoDbConnection()

	if err != nil {
		return nil, err
	}

	collection := client.Database(DbName).Collection(CollectionName)

	return collection, nil
}
