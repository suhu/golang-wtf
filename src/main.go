package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/what-the-fake/src/lib/router"
)

func main() {

	// Get the "PORT" env variable
	routerEngine := router.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST"})

	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(allowedOrigins, allowedMethods)(routerEngine)))

}
