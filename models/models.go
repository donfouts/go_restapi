package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Create Struct
type App struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IsActive   bool
	Createdon  string         `json:"Createdon" bson:"Createdon,omitempty"`
	Appname  string           `json:"Appname" bson:"Appname,omitempty"`
	Devowner  string          `json:"Devowner" bson:"Devowner,omitempty"`
	Software  string          `json:"Software" bson:"Software,omitempty"`
	Platform  string          `json:"Platform" bson:"Platform,omitempty"`
	Farms *Farm               `json:"Farms" bson:"Farm,omitempty"`
	Tags  *Tag                `json:"Tag" bson:"Tag,omitempty"`
}

type Farm struct {
	Fid	int       `json:"Fid" bson:"Fid"`
	Name string   `json:"Name" bson:"Name,omitempty"`
}

type Tag struct {
	Tag	string    `json:"Tag" bson:"Tag,omitempty"`
}