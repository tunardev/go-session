package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/rs/cors"
	"github.com/tunardev/go-session/controllers"
	db "github.com/tunardev/go-session/database"
	"github.com/tunardev/go-session/models"
	"github.com/tunardev/go-session/routes"
)

var (
    key = []byte(os.Getenv("SECRET_KEY"))
    session = sessions.NewCookieStore(key)
)

func main() {
    router := mux.NewRouter()
    db.ConnectDb()
    defer db.Close()    

    controller := controllers.Controller{ Database: db.Db, Session: session }

    routes.Setup(router, controller)

    router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        err := models.Error{
            Message: "404 not found",
            Success: false,
        }

        json.NewEncoder(w).Encode(err)
    })

    handler := cors.Default().Handler(router)
    if err := http.ListenAndServe(":8000", handler); err != nil {
        panic(err)
    }
    
}