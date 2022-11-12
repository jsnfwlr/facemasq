package addresses

import "facemasq/lib/extensions"

func GetRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/api/address`, Handler: Save, Methods: "POST", Name: "SaveAddress"},
		{Path: `/api/address`, Handler: Delete, Methods: "DELETE", Name: "DeleteAddress"},
	}
	return
}
