package users

import (
	"github.com/oupsla/IFI-Gorilla/controllers/users"
	"github.com/oupsla/IFI-Gorilla/routes"
)

func GetRoutes() []routes.Route {
	Routes := []routes.Route{}

	Routes = append(Routes, routes.Route{
		Name:        "Get all users",
		Method:      "GET",
		Path:        "/users",
		HandlerFunc: users.GetAll,
	})

	return Routes
}
