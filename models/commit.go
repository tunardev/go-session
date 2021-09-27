package models

type Commit struct{
	UserID string `json:"userID" bson:"userID"`
	PostID string `json:"postID" bson:"postID"`
	Content string `json:"content" bson:"content"`
}