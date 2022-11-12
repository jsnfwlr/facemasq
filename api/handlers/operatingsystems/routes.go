package operatingsystems

import "facemasq/lib/extensions"

func GetRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/api/operatingSystems`, Handler: Save, Methods: "POST", Name: "SaveOperatingSystem"},
		{Path: `/api/operatingSystems/{ID:[0-9]+}`, Handler: Delete, Methods: "DELETE", Name: "DeleteOperatingSystem"}}
	return
}
