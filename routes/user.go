package routes

import (
	"github.com/gorilla/mux"
	"github.com/tunardev/dev-post-server/controllers"
)

func UserRoutes(router *mux.Router, controllers controllers.Controller) {

	router.HandleFunc("/api/user/follow", controllers.Hello)

}