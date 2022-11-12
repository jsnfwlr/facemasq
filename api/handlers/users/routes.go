package users

import "facemasq/lib/extensions"

func GetRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/api/users`, Handler: SaveUser, Methods: "POST", Name: "SaveUser"},
		{Path: `/api/users/{ID:[0-9]+}`, Handler: DeleteUser, Methods: "DELETE", Name: "DeleteUser"}}
	return
}
