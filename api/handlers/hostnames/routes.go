package hostnames

import "facemasq/lib/extensions"

func GetRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/api/hostname`, Handler: Save, Methods: "POST", Name: "SaveHostname"},
		{Path: `/api/hostname`, Handler: Delete, Methods: "DELETE", Name: "DeleteHostname"},
	}
	return
}
