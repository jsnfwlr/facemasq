package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"facemasq/handlers/control"
	"facemasq/handlers/devices"
	"facemasq/handlers/params"
	"facemasq/handlers/settings"
	"facemasq/lib/extensions"
)

type Router struct {
	Mux *mux.Router
}

func BuildRoutes() (router Router) {
	router = Router{
		Mux: mux.NewRouter(),
	}

	var routes []extensions.RouteDefinition

	// For now, plugin routes go first, to prevent them over riding the built in routes.
	// This may change in the future, but will require some sort of plugin security, with user confirmation and signed binaries
	routes = extensions.Extensions.GetRoutes()

	routes = append(staticRoutes(), routes...)
	routes = append(statusRoutes(), routes...)
	routes = append(paramRoutes(), routes...)
	routes = append(settingsRoutes(), routes...)
	routes = append(deviceRoutes(), routes...)
	routes = append(uiRoutes(), routes...)
	routes = append(controlRoutes(), routes...)

	for r := range routes {
		if handlerFunc, ok := routes[r].Handler.(func(http.ResponseWriter, *http.Request)); ok {
			if routes[r].Methods != "WS" {
				router.Mux.HandleFunc(routes[r].Path, handlerFunc).Methods(routes[r].Methods).Name(routes[r].Name)
			}
		} else {
			handler := routes[r].Handler.(http.Handler)
			router.Mux.Handle(routes[r].Path, handler).Methods(routes[r].Methods).Name(routes[r].Name)
		}
	}

	return
}

func controlRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/api/exit`, Handler: control.Exit, Methods: "GET", Name: "APIExit"},
		{Path: `/exit`, Handler: control.Exit, Methods: "GET", Name: "Exit"},
	}
	return
}

func deviceRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/api/records/all`, Handler: devices.GetAll, Methods: "GET", Name: "GetAllDevices"},
		{Path: `/api/records/active`, Handler: devices.GetActive, Methods: "GET", Name: "GetActiveDevices"},
		{Path: `/api/records/chart`, Handler: devices.GetDashboardChartData, Methods: "GET", Name: "GetDashboardChartData"},
		{Path: `/api/records/unknown`, Handler: devices.GetUnknown, Methods: "GET", Name: "GetUnknownDevices"},
		{Path: `/api/records/trends`, Handler: devices.GetTrendData, Methods: "GET", Name: "GetDeviceTrendData"},
		{Path: `/api/records/investigate`, Handler: devices.InvestigateAddresses, Methods: "POST", Name: "InvestigateAddresses"},
		{Path: `/api/records/device`, Handler: devices.SaveDevice, Methods: "POST", Name: "SaveDevice"},
		{Path: `/api/records/interface`, Handler: devices.SaveInterface, Methods: "POST", Name: "SaveInterface"},
		{Path: `/api/records/address`, Handler: devices.SaveAddress, Methods: "POST", Name: "SaveAddress"},
		{Path: `/api/records/hostname`, Handler: devices.SaveHostname, Methods: "POST", Name: "SaveHostname"},

		{Path: `/ws/records/changed`, Handler: devices.GetRecentChanges, Methods: "WS", Name: "WSRecentlyChangedDevices"},
	}
	return
}

func settingsRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/api/settings/{userID:[0-9]+}`, Handler: settings.GetUserSettings, Methods: "GET", Name: "GetUserSettings"},
		{Path: `/api/settings`, Handler: settings.GetAppSettings, Methods: "GET", Name: "GetAppSettings"},
		{Path: `/api/setting/{userID:[0-9]+}`, Handler: settings.SaveUserSetting, Methods: "POST", Name: "SaveUserSetting"},
		{Path: `/api/setting`, Handler: settings.SaveAppSetting, Methods: "POST", Name: "SaveAppSetting"},
	}
	return
}

func statusRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/status`, Handler: control.Status, Methods: "GET", Name: "GetStatus"},
	}
	return
}

func paramRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/api/params`, Handler: params.GetAll, Methods: "GET", Name: "GetAllParams"},
		{Path: `/api/categories`, Handler: params.SaveCategory, Methods: "POST", Name: "SaveCategory"},
		{Path: `/api/categories/{ID:[0-9]+}`, Handler: params.DeleteCategory, Methods: "DELETE", Name: "DeleteCategory"},
		{Path: `/api/statuses`, Handler: params.SaveStatus, Methods: "POST", Name: "SaveStatus"},
		{Path: `/api/statuses/{ID:[0-9]+}`, Handler: params.DeleteStatus, Methods: "DELETE", Name: "DeleteStatus"},
		{Path: `/api/locations`, Handler: params.SaveLocation, Methods: "POST", Name: "SaveLocation"},
		{Path: `/api/locations/{ID:[0-9]+}`, Handler: params.DeleteLocation, Methods: "DELETE", Name: "DeleteLocation"},
		{Path: `/api/maintainers`, Handler: params.SaveMaintainer, Methods: "POST", Name: "SaveMaintainer"},
		{Path: `/api/maintainers/{ID:[0-9]+}`, Handler: params.DeleteMaintainer, Methods: "DELETE", Name: "DeleteMaintainer"},
		{Path: `/api/deviceTypes`, Handler: params.SaveDeviceType, Methods: "POST", Name: "SaveDeviceType"},
		{Path: `/api/deviceTypes/{ID:[0-9]+}`, Handler: params.DeleteDeviceType, Methods: "DELETE", Name: "DeleteDeviceType"},
		{Path: `/api/architectures`, Handler: params.SaveArchitecture, Methods: "POST", Name: "SaveArchitecture"},
		{Path: `/api/architectures/{ID:[0-9]+}`, Handler: params.DeleteArchitecture, Methods: "DELETE", Name: "DeleteArchitecture"},
		{Path: `/api/operatingSystems`, Handler: params.SaveOperatingSystem, Methods: "POST", Name: "SaveOperatingSystem"},
		{Path: `/api/operatingSystems/{ID:[0-9]+}`, Handler: params.DeleteOperatingSystem, Methods: "DELETE", Name: "DeleteOperatingSystem"},
		{Path: `/api/vlans`, Handler: params.SaveVLAN, Methods: "POST", Name: "SaveVLAN"},
		{Path: `/api/vlans/{ID:[0-9]+}`, Handler: params.DeleteVLAN, Methods: "DELETE", Name: "DeleteVLAN"},
		{Path: `/api/users`, Handler: params.SaveUser, Methods: "POST", Name: "SaveUser"},
		{Path: `/api/users/{ID:[0-9]+}`, Handler: params.DeleteUser, Methods: "DELETE", Name: "DeleteUser"},
		{Path: `/api/interfaceTypes`, Handler: params.SaveInterfaceType, Methods: "POST", Name: "SaveInterfaceType"},
		{Path: `/api/interfaceTypes/{ID:[0-9]+}`, Handler: params.DeleteInterfaceType, Methods: "DELETE", Name: "DeleteInterfaceType"},
	}
	return
}

func uiRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/`, Handler: http.FileServer(http.Dir("../web")), Methods: "GET", Name: "ServeUI"},
	}
	return
}

func staticRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/{filename:[a-zA-Z0-9=\.\/]+}`, Handler: control.Static, Methods: "GET", Name: "ServeStatic"},
	}
	return
}
