package devices

import "facemasq/lib/extensions"

func GetRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/api/device`, Handler: Save, Methods: "POST", Name: "SaveDevice"},
		{Path: `/api/device`, Handler: Delete, Methods: "DELETE", Name: "DeleteDevice"},
	}
	return
}
