package interfacetypes

import "facemasq/lib/extensions"

func GetRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/api/interfaceTypes`, Handler: SaveInterfaceType, Methods: "POST", Name: "SaveInterfaceType"},
		{Path: `/api/interfaceTypes/{ID:[0-9]+}`, Handler: DeleteInterfaceType, Methods: "DELETE", Name: "DeleteInterfaceType"}}
	return
}
