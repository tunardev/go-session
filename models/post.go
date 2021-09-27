package models

import "gopkg.in/mgo.v2/bson"

type Post struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	UserID string `json:"userID" bson:"userID"`
	Title string `json:"title" bson:"title"`
	Content string `json:"content" bson:"content"`
	Category []string `json:"category" bson:"category"`
	Views string `json:"views" bson:"views"`
	Commits []Commit `json:"commits" bson:"commits"`
}