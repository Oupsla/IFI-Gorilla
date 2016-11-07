package main

import (
	"github.com/bnjjj/go_api/routes"
	"github.com/gorilla/mux"

	"net/http"
)

func main() {
	router := mux.NewRouter()

	// no strict slash
	router.StrictSlash(false)

	for _, route := range routes.GetUsersRoutes() {
		router.
			Methods(route.Method).
			Path("/api/v1" + route.Path).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	http.Handle("/", router)

	http.ListenAndServe(":8080", router)
}
