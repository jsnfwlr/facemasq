package maintainers

import "facemasq/lib/extensions"

func GetRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/api/maintainers`, Handler: Save, Methods: "POST", Name: "SaveMaintainer"},
		{Path: `/api/maintainers/{ID:[0-9]+}`, Handler: Delete, Methods: "DELETE", Name: "DeleteMaintainer"}}
	return
}
