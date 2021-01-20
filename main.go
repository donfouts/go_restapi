package main

import (
	"context"
	"encoding/json"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const dbName = "inventory"
const collectionName = "applications"
const port = 8000

//Create Struct
type App struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IsActive  bool
	Createdon string `json:"Createdon" bson:"Createdon,omitempty"`
	Appname   string `json:"Appname" bson:"Appname,omitempty"`
	Devowner  string `json:"Devowner" bson:"Devowner,omitempty"`
	Software  string `json:"Software" bson:"Software,omitempty"`
	Platform  string `json:"Platform" bson:"Platform,omitempty"`
	Farms     *Farm  `json:"Farms" bson:"Farm,omitempty"`
	Tags      *Tag   `json:"Tag" bson:"Tag,omitempty"`
}

type Farm struct {
	Fid  int    `json:"Fid" bson:"Fid"`
	Name string `json:"Name" bson:"Name,omitempty"`
}

type Tag struct {
	Tag string `json:"Tag" bson:"Tag,omitempty"`
}

func getApplication(c *fiber.Ctx) {
	collection, err := getMongoDbCollection(dbName, collectionName)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	var filter bson.M = bson.M{}

	if c.Params("id") != "" {
		id := c.Params("id")
		objID, _ := primitive.ObjectIDFromHex(id)
		filter = bson.M{"_id": objID}
	}

	var results []bson.M
	cur, err := collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	cur.All(context.Background(), &results)

	if results == nil {
		c.SendStatus(404)
		return
	}

	json, _ := json.Marshal(results)
	c.Send(json)
}

func createApplication(c *fiber.Ctx) {
	collection, err := getMongoDbCollection(dbName, collectionName)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	var app App
	json.Unmarshal([]byte(c.Body()), &app)

	res, err := collection.InsertOne(context.Background(), app)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	response, _ := json.Marshal(res)
	c.Send(response)
}

func updateApplication(c *fiber.Ctx) {
	collection, err := getMongoDbCollection(dbName, collectionName)
	if err != nil {
		c.Status(500).Send(err)
		return
	}
	var app App
	json.Unmarshal([]byte(c.Body()), &app)

	update := bson.M{
		"$set": app,
	}

	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	res, err := collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	response, _ := json.Marshal(res)
	c.Send(response)
}

func deleteApplication(c *fiber.Ctx) {
	collection, err := getMongoDbCollection(dbName, collectionName)

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	res, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	jsonResponse, _ := json.Marshal(res)
	c.Send(jsonResponse)
}

func main() {
	app := fiber.New()

	app.Get("/applications/:id?", getApplication)
	app.Post("/applications", createApplication)
	app.Put("/applications/:id", updateApplication)
	app.Delete("/applications/:id", deleteApplication)

	app.Listen(port)
}
