package trends

import "facemasq/lib/extensions"

func GetRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/api/trends/connections`, Handler: GetConnectionTrends, Methods: "GET", Name: "GetConnectionTrends"},
	}
	return
}
