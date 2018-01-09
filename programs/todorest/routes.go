package main

import "github.com/gorilla/mux"
import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		IndexHandler,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoHandler,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		ShowHandler,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		CreateHandler,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
