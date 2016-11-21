package ifi

import (
	"github.com/oupsla/IFI-Gorilla/controllers/ifi"
	"github.com/oupsla/IFI-Gorilla/routes"
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
