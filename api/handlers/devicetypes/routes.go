package devicetypes

import "facemasq/lib/extensions"

func GetRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/api/deviceTypes`, Handler: Save, Methods: "POST", Name: "SaveDeviceType"},
		{Path: `/api/deviceTypes/{ID:[0-9]+}`, Handler: Delete, Methods: "DELETE", Name: "DeleteDeviceType"}}
	return
}
