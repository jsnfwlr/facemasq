package statuses

import "facemasq/lib/extensions"

func GetRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/api/statuses`, Handler: Save, Methods: "POST", Name: "SaveStatus"},
		{Path: `/api/statuses/{ID:[0-9]+}`, Handler: Delete, Methods: "DELETE", Name: "DeleteStatus"},
	}
	return
}
