package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2"
)

type Controller struct{
	Database mgo.Session
	Session *sessions.CookieStore
}

func (c Controller) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}