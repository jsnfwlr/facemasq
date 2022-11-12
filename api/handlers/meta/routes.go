package meta

import "facemasq/lib/extensions"

func GetRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{

		{Path: `/api/settings/{userID:[0-9]+}`, Handler: GetUserSettings, Methods: "GET", Name: "GetUserSettings"},
		{Path: `/api/setting/{userID:[0-9]+}`, Handler: SaveUserSetting, Methods: "POST", Name: "SaveUserSetting"},

		{Path: `/api/settings`, Handler: GetAppSettings, Methods: "GET", Name: "GetAppSettings"},
		{Path: `/api/setting`, Handler: SaveAppSetting, Methods: "POST", Name: "SaveAppSetting"},
	}
	return
}
