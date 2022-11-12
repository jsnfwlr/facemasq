package params

import "facemasq/lib/extensions"

func GetRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/api/params`, Handler: GetAllParams, Methods: "GET", Name: "GetAllParams"},
	}
	return
}
