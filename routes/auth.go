package routes

import (
	"github.com/gorilla/mux"
	"github.com/tunardev/dev-post-server/controllers"
)

func AuthRoutes(router *mux.Router, controllers controllers.Controller) {

	router.HandleFunc("/api/auth/signup", controllers.Hello)
	router.HandleFunc("/api/auth/signin", controllers.Hello)
	router.HandleFunc("/api/auth/logout", controllers.Hello)
	router.HandleFunc("/api/auth/me", controllers.Hello)
	router.HandleFunc("/api/user/follow", controllers.Hello)

}