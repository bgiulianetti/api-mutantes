package api

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"mutant",
		"POST",
		"/mutant",
		DetectMutant,
	},
	Route{
		"Stats",
		"GET",
		"/stats",
		Stats,
	},
	Route{
		"health",
		"GET",
		"/health",
		Health,
	},
}
