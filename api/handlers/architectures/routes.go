package architectures

import "facemasq/lib/extensions"

func GetRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/api/architectures`, Handler: Save, Methods: "POST", Name: "SaveArchitecture"},
		{Path: `/api/architectures/{ID:[0-9]+}`, Handler: Delete, Methods: "DELETE", Name: "DeleteArchitecture"}}
	return
}
