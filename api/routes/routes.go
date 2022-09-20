package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"facemasq/handlers/control"
	"facemasq/handlers/devices"
	"facemasq/handlers/dnsmasq"
	"facemasq/handlers/params"
	"facemasq/handlers/settings"
)

type Router struct {
	Mux *mux.Router
}

func BuildRoutes() (router Router) {
	router = Router{
		Mux: mux.NewRouter(),
	}

	router.controlRoutes()
	router.uiRoutes()
	router.masqRoutes()
	router.deviceRoutes()
	router.settingsRoutes()
	router.paramRoutes()
	router.statusRoutes()
	router.staticRoutes()

	return
}

func (router *Router) controlRoutes() {
	router.Mux.HandleFunc(`/api/exit`, control.Exit).Methods("GET").Name("APIExit")
	router.Mux.HandleFunc(`/exit`, control.Exit).Methods("GET").Name("Exit")
}

func (router *Router) masqRoutes() {
	router.Mux.HandleFunc(`/api/dhcp`, dnsmasq.WriteDHCPConfig).Methods("GET").Name("WriteDHCPConfigGet")
	router.Mux.HandleFunc(`/api/dhcp`, dnsmasq.WriteDHCPConfig).Methods("PUT").Name("WriteDHCPConfigPut")
	router.Mux.HandleFunc(`/api/dns`, dnsmasq.WriteDNSConfig).Methods("GET").Name("WriteDNSConfigGet")
	router.Mux.HandleFunc(`/api/dns`, dnsmasq.WriteDNSConfig).Methods("GET").Name("WriteDNSConfigPut")
}

func (router *Router) deviceRoutes() {
	router.Mux.HandleFunc(`/api/records/all`, devices.GetAll).Methods("GET").Name("GetAllDevices")
	router.Mux.HandleFunc(`/api/records/active`, devices.GetActive).Methods("GET").Name("GetActiveDevices")
	router.Mux.HandleFunc(`/api/records/chart`, devices.GetDashboardChartData).Methods("GET").Name("GetDashboardChartData")
	router.Mux.HandleFunc(`/api/records/unknown`, devices.GetUnknown).Methods("GET").Name("GetUnknownDevices")
	router.Mux.HandleFunc(`/api/records/trends`, devices.GetTrendData).Methods("GET").Name("GetDeviceTrendData")
	router.Mux.HandleFunc(`/api/records/investigate`, devices.InvestigateAddresses).Methods("POST").Name("InvestigateAddresses")

	router.Mux.HandleFunc(`/api/records/device`, devices.SaveDevice).Methods("POST").Name("SaveDevice")
	router.Mux.HandleFunc(`/api/records/interface`, devices.SaveInterface).Methods("POST").Name("SaveInterface")
	router.Mux.HandleFunc(`/api/records/address`, devices.SaveAddress).Methods("POST").Name("SaveAddress")
	router.Mux.HandleFunc(`/api/records/hostname`, devices.SaveHostname).Methods("POST").Name("SaveHostname")
}

func (router *Router) settingsRoutes() {
	router.Mux.HandleFunc(`/api/settings/{userID:[0-9]+}`, settings.GetUserSettings).Methods("GET").Name("GetUserSettings")
	router.Mux.HandleFunc(`/api/settings`, settings.GetAppSettings).Methods("GET").Name("GetAppSettings")
	router.Mux.HandleFunc(`/api/setting/{userID:[0-9]+}`, settings.SaveUserSetting).Methods("POST").Name("SaveUserSetting")
	router.Mux.HandleFunc(`/api/setting`, settings.SaveAppSetting).Methods("POST").Name("SaveAppSetting")
}

func (router *Router) statusRoutes() {
	router.Mux.HandleFunc(`/status`, control.Status).Methods("GET").Name("GetStatus")
}

func (router *Router) paramRoutes() {
	router.Mux.HandleFunc(`/api/params`, params.GetAll).Methods("GET").Name("GetAllParams")
	router.Mux.HandleFunc(`/api/categories`, params.SaveCategory).Methods("POST").Name("SaveCategory")
	router.Mux.HandleFunc(`/api/categories/{ID:[0-9]+}`, params.DeleteCategory).Methods("DELETE").Name("DeleteCategory")
	router.Mux.HandleFunc(`/api/statuses`, params.SaveStatus).Methods("POST").Name("SaveStatus")
	router.Mux.HandleFunc(`/api/statuses/{ID:[0-9]+}`, params.DeleteStatus).Methods("DELETE").Name("DeleteStatus")
	router.Mux.HandleFunc(`/api/locations`, params.SaveLocation).Methods("POST").Name("SaveLocation")
	router.Mux.HandleFunc(`/api/locations/{ID:[0-9]+}`, params.DeleteLocation).Methods("DELETE").Name("DeleteLocation")
	router.Mux.HandleFunc(`/api/maintainers`, params.SaveMaintainer).Methods("POST").Name("SaveMaintainer")
	router.Mux.HandleFunc(`/api/maintainers/{ID:[0-9]+}`, params.DeleteMaintainer).Methods("DELETE").Name("DeleteMaintainer")
	router.Mux.HandleFunc(`/api/deviceTypes`, params.SaveDeviceType).Methods("POST").Name("SaveDeviceType")
	router.Mux.HandleFunc(`/api/deviceTypes/{ID:[0-9]+}`, params.DeleteDeviceType).Methods("DELETE").Name("DeleteDeviceType")
	router.Mux.HandleFunc(`/api/architectures`, params.SaveArchitecture).Methods("POST").Name("SaveArchitecture")
	router.Mux.HandleFunc(`/api/architectures/{ID:[0-9]+}`, params.DeleteArchitecture).Methods("DELETE").Name("DeleteArchitecture")
	router.Mux.HandleFunc(`/api/operatingSystems`, params.SaveOperatingSystem).Methods("POST").Name("SaveOperatingSystem")
	router.Mux.HandleFunc(`/api/operatingSystems/{ID:[0-9]+}`, params.DeleteOperatingSystem).Methods("DELETE").Name("DeleteOperatingSystem")
	router.Mux.HandleFunc(`/api/vlans`, params.SaveVLAN).Methods("POST").Name("SaveVLAN")
	router.Mux.HandleFunc(`/api/vlans/{ID:[0-9]+}`, params.DeleteVLAN).Methods("DELETE").Name("DeleteVLAN")
	router.Mux.HandleFunc(`/api/users`, params.SaveUser).Methods("POST").Name("SaveUser")
	router.Mux.HandleFunc(`/api/users/{ID:[0-9]+}`, params.DeleteUser).Methods("DELETE").Name("DeleteUser")
	router.Mux.HandleFunc(`/api/interfaceTypes`, params.SaveInterfaceType).Methods("POST").Name("SaveInterfaceType")
	router.Mux.HandleFunc(`/api/interfaceTypes/{ID:[0-9]+}`, params.DeleteInterfaceType).Methods("DELETE").Name("DeleteInterfaceType")
}

func (router *Router) uiRoutes() {
	router.Mux.Handle(`/`, http.FileServer(http.Dir("../ui"))).Methods("GET").Name("ServeUI")
}

func (router *Router) staticRoutes() {
	router.Mux.HandleFunc(`/{filename:[a-zA-Z0-9=\.\/]+}`, control.Static).Methods("GET").Name("ServeStatic")
}
