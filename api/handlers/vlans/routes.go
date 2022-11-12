package vlans

import "facemasq/lib/extensions"

func GetRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/api/vlans`, Handler: Save, Methods: "POST", Name: "SaveVLAN"},
		{Path: `/api/vlans/{ID:[0-9]+}`, Handler: Delete, Methods: "DELETE", Name: "DeleteVLAN"}}
	return
}
