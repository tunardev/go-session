package routes

import (
	"github.com/gorilla/mux"
	"github.com/tunardev/go-session/controllers"
)

func Setup(router *mux.Router, controllers controllers.Controller) {
	router.HandleFunc("/", controllers.Hello)
	router.HandleFunc("/api/auth/signup", controllers.SignUp)
	router.HandleFunc("/api/auth/signin", controllers.SignIn)
	router.HandleFunc("/api/auth/logout", controllers.Logout)
	router.HandleFunc("/api/auth/me", controllers.Me)
}