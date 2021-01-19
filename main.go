package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"./models"
	"./helper"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = helper.ConnectDB()

func main() {

	//Init Router
	r := mux.NewRouter()

  	// arrange our route
	r.HandleFunc("/api/apps", getapps).Methods("GET")
	r.HandleFunc("/api/apps/{id}", getapp).Methods("GET")
	r.HandleFunc("/api/apps", createapp).Methods("POST")
	r.HandleFunc("/api/apps/{id}", updateapp).Methods("PUT")
	r.HandleFunc("/api/apps/{id}", deleteapp).Methods("DELETE")

  	// set our port address
	log.Fatal(http.ListenAndServe(":8000", r))

}

func getapps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// we created app array
	var apps []models.App

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var app models.App
		// & character returns the memory address of the following variable.
		err := cur.Decode(&app) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		apps = append(apps, app)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(apps) // encode similar to serialize process.
}

func getapp(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var app models.App
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&app)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(app)
}

func createapp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var app models.App

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&app)

	// insert our app model.
	result, err := collection.InsertOne(context.TODO(), app)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func updateapp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	//Get id from parameters
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var app models.App
	
	// Create filter
	filter := bson.M{"_id": id}

	// Read update model from body request
	_ = json.NewDecoder(r.Body).Decode(&app)

	// prepare update model.
	update := bson.D{
		{"$set", bson.D{
			{"isActive", app.IsActive},
			{"createdon", app.Createdon},
			{"appname", app.Appname},
			{"devowner", app.Devowner},
			{"software", app.Software},
			{"platform", app.Platform},		
			{"farms", bson.D{
				{"id", app.Farms.Fid},
				{"name", app.Farms.Name},
				},
			},
			{"tags", bson.D{
				{"tag", app.Tags},
				},
			},
		},
	},
}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&app)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	app.ID = id

	json.NewEncoder(w).Encode(app)
}

func deleteapp(w http.ResponseWriter, r *http.Request) {
	// Set header
	w.Header().Set("Content-Type", "application/json")

	// get params
	var params = mux.Vars(r)

	// string to primitve.ObjectID
	id, err := primitive.ObjectIDFromHex(params["id"])

	// prepare filter.
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}