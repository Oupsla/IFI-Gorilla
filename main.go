package main

import (
	"github.com/oupsla/IFI-Gorilla/routes"
	"github.com/oupsla/IFI-Gorilla/routes/ifi"

	"github.com/gorilla/mux"

	"fmt"
	"net/http"
)

func getAllRoutes() []routes.Route {
	Routes := []routes.Route{}

	Routes = append(Routes, ifi.GetRoutes()...)

	return Routes
}

func main() {
	router := mux.NewRouter()

	// no strict slash
	router.StrictSlash(false)

	for _, route := range getAllRoutes() {
		router.
			Methods(route.Method).
			Path("/api/v1" + route.Path).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	http.Handle("/", router)
	fmt.Println("Server listening on localhost:8080")
	http.ListenAndServe(":8080", router)

}
