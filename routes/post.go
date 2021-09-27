package routes

import (
	"github.com/gorilla/mux"
	"github.com/tunardev/dev-post-server/controllers"
)

func PostRoutes(router *mux.Router, controllers controllers.Controller) {
	router.HandleFunc("/post", controllers.Hello)
}