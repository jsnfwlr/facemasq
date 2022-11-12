package control

import (
	"facemasq/lib/extensions"
	"net/http"
)

func GetRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/api/exit`, Handler: Exit, Methods: "GET", Name: "APIExit"},
		{Path: `/exit`, Handler: Exit, Methods: "GET", Name: "Exit"},

		{Path: `/`, Handler: http.FileServer(http.Dir("../web")), Methods: "GET", Name: "ServeUI"},

		{Path: `/state`, Handler: State, Methods: "GET", Name: "GetStatus"},
	}
	return
}

func GetStaticRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/{filename:[a-zA-Z0-9=\.\/]+}`, Handler: Static, Methods: "GET", Name: "ServeStatic"}}
	return
}
