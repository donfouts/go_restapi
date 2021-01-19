package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"net/http"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)
// mongouser
// A67qtTuPT4knq3Z
// ConnectDB : This is helper function to connect mongoDB


// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
func ConnectDB() *mongo.Collection {

	
	// Set client options
	// clientOptions := options.Client().ApplyURI("your_cluster_endpoint")

	// Connect to MongoDB
	// client, err := mongo.Connect(context.TODO(), clientOptions)

	// copied from mongodb atlas
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://mongouser:A67qtTuPT4knq3Z@cluster0.hmru5.mongodb.net/inventory?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	defer client.Disconnect(ctx)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		// Can't connect to Mongo server
		log.Fatal(err)
	}
	
	fmt.Println("Connected to MongoDB!")
	collection := client.Database("apps").Collection("inventory")
	return collection
}

// ErrorResponse : This is error model.
type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

func GetError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}