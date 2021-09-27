package routes

import (
	"github.com/gorilla/mux"
	"github.com/tunardev/dev-post-server/controllers"
)

func PostRoutes(router *mux.Router, controllers controllers.Controller) {
	
	router.HandleFunc("/api/posts", controllers.Hello)
	router.HandleFunc("/api/post", controllers.Hello)
	router.HandleFunc("/api/post/:id", controllers.Hello)
	router.HandleFunc("/api/post/:id/views", controllers.Hello)
	router.HandleFunc("/api/post/search", controllers.Hello)

}