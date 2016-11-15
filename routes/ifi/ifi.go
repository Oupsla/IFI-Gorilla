package ifi

import (
	"github.com/bnjjj/go_api/controllers/ifi"
	"github.com/bnjjj/go_api/routes"
)

func GetRoutes() []routes.Route {
	Routes := []routes.Route{}

	Routes = append(Routes, routes.Route{
		Name:        "Get welcome message",
		Method:      "GET",
		Path:        "/welcome",
		HandlerFunc: ifi.Welcome,
	})

	return Routes
}
