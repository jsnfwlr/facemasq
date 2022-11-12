package interfaces

import "facemasq/lib/extensions"

func GetRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/api/interface`, Handler: Save, Methods: "POST", Name: "SaveInterface"},
		{Path: `/api/interface`, Handler: Delete, Methods: "DELETE", Name: "DeleteInterface"},
	}
	return
}
