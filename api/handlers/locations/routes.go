package locations

import "facemasq/lib/extensions"

func GetRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/api/locations`, Handler: Save, Methods: "POST", Name: "SaveLocation"},
		{Path: `/api/locations/{ID:[0-9]+}`, Handler: Delete, Methods: "DELETE", Name: "DeleteLocation"}}
	return
}
