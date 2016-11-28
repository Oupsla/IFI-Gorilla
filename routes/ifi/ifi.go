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
	}, routes.Route{
	  Name:        "2nd route",
	  Method:      "GET",
	  Path:        "/enigma", //L'url réel sera en fait /api/v1/myroute car tout est préfixé par /api/v1
	  HandlerFunc: ifi.GetEnigma,
	}, routes.Route{
	  Name:        "3nd route",
	  Method:      "POST",
	  Path:        "/enigma", //L'url réel sera en fait /api/v1/myroute car tout est préfixé par /api/v1
	  HandlerFunc: ifi.PostEnigma,
	},routes.Route{
	  Name:        "4nd route",
	  Method:      "GET",
	  Path:        "/plage", //L'url réel sera en fait /api/v1/myroute car tout est préfixé par /api/v1
	  HandlerFunc: ifi.GetPlage,
	},routes.Route{
	  Name:        "5nd route",
	  Method:      "POST",
	  Path:        "/plage", //L'url réel sera en fait /api/v1/myroute car tout est préfixé par /api/v1
	  HandlerFunc: ifi.PostPlage,
	},routes.Route{
	  Name:        "6nd route",
	  Method:      "DELETE",
	  Path:        "/plage", //L'url réel sera en fait /api/v1/myroute car tout est préfixé par /api/v1
	  HandlerFunc: ifi.RemovePlage,
	},routes.Route{
	  Name:        "7nd route",
	  Method:      "PUT",
	  Path:        "/plage", //L'url réel sera en fait /api/v1/myroute car tout est préfixé par /api/v1
	  HandlerFunc: ifi.EditPlage,
	},routes.Route{
	  Name:        "8nd route",
	  Method:      "GET",
	  Path:        "/present", //L'url réel sera en fait /api/v1/myroute car tout est préfixé par /api/v1
	  HandlerFunc: ifi.GetPresent,
	})

	return Routes
}
