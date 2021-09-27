package routes

import (
	"github.com/gorilla/mux"
	"github.com/tunardev/dev-post-server/controllers"
)

func CommitRoutes(router *mux.Router, controllers controllers.Controller) {

	router.HandleFunc("/api/commit/:id", controllers.Hello)
	router.HandleFunc("/api/commit/:id/like", controllers.Hello)

}