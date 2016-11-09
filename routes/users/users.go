package users

import (
	"github.com/bnjjj/go_api/controllers/users"
	"github.com/bnjjj/go_api/routes"
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
