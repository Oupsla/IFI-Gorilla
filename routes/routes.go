package routes

import (
	"net/http"
)

type Route struct {
	Method      string
	Path        string
	Name        string
	HandlerFunc http.HandlerFunc
}
