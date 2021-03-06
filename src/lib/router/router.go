package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/what-the-fake/src/lib/controllerapi"
	"github.com/what-the-fake/src/lib/repo"
)

var controller = &controllerapi.Controller{Repository: repo.Repository{}}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"Authentication",
		"POST",
		"/get-token",
		controller.GetToken,
	},
	Route{
		"SearchSites",
		"GET",
		"/search/{query}",
		//controllerapi.AuthenticationMiddleware(controller.SearchSites),
		controller.SearchSites,
	}}

// Get Product by {id}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
