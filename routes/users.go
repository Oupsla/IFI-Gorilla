package routes

import (
	"github.com/bnjjj/go_api/controllers"

	"net/http"
)

type Route struct {
	Method      string
	Path        string
	Name        string
	HandlerFunc http.HandlerFunc
}

func GetUsersRoutes() []Route {
	Routes := []Route{}

	Routes = append(Routes, Route{
		Name:        "Get all users",
		Method:      "GET",
		Path:        "/users",
		HandlerFunc: controllers.GetAll,
	})

	return Routes
}
