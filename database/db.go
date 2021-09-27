package db

import (
	"os"

	"gopkg.in/mgo.v2"
)

var Db mgo.Session

func ConnectDb() {
	url := os.Getenv("MONGODB_URL")
	db, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	Db = *db
}

func  Close() {
	Db.Close()
}

