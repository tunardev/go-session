package routes

import (
	"github.com/gorilla/mux"
	"github.com/tunardev/dev-post-server/controllers"
)

func Setup(router *mux.Router, controllers controllers.Controller) {
	router.HandleFunc("/", controllers.Hello)

	PostRoutes(router, controllers)
	UserRoutes(router, controllers)
}