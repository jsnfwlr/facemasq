package categories

import "facemasq/lib/extensions"

func GetRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/api/categories`, Handler: Save, Methods: "POST", Name: "SaveCategory"},
		{Path: `/api/categories/{ID:[0-9]+}`, Handler: Delete, Methods: "DELETE", Name: "DeleteCategory"},
	}
	return
}
