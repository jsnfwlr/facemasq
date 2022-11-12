package grids

import "facemasq/lib/extensions"

func GetRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/api/grids/dosier`, Handler: GetDeviceDosier, Methods: "POST", Name: "GetDeviceDosier"},

		{Path: `/api/grids/alldevices`, Handler: GetAllDevices, Methods: "GET", Name: "GetAllDevices"},
		{Path: `/api/grids/activedevices`, Handler: GetActiveDevices, Methods: "GET", Name: "GetActiveDevices"},
		{Path: `/api/grids/unknowndevices`, Handler: GetUnknownDevices, Methods: "GET", Name: "GetUnknownDevices"},

		{Path: `/ws/grids/changed`, Handler: GetRecentlyChangedDevices, Methods: "WS", Name: "WSRecentlyChangedDevices"},
	}
	return
}
