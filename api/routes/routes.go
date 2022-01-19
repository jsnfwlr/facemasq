package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jsnfwlr/facemasq/api/handlers/control"
	"github.com/jsnfwlr/facemasq/api/handlers/devices"
	"github.com/jsnfwlr/facemasq/api/handlers/dnsmasq"
	"github.com/jsnfwlr/facemasq/api/handlers/params"
	"github.com/jsnfwlr/facemasq/api/handlers/settings"
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
	router.staticRoutes()

	return
}

func (router *Router) controlRoutes() {
	router.Mux.HandleFunc(`/exit`, control.Exit).Methods("GET")
}

func (router *Router) uiRoutes() {
	router.Mux.Handle(`/`, http.FileServer(http.Dir("../ui"))).Methods("GET")
}

func (router *Router) masqRoutes() {
	router.Mux.HandleFunc(`/dhcp`, dnsmasq.WriteDHCPConfig).Methods("GET")
	router.Mux.HandleFunc(`/dns`, dnsmasq.WriteDNSConfig).Methods("GET")
}

func (router *Router) deviceRoutes() {
	router.Mux.HandleFunc(`/records/all`, devices.GetAll).Methods("GET")
	router.Mux.HandleFunc(`/records/active`, devices.GetActive).Methods("GET")
	router.Mux.HandleFunc(`/records/chart`, devices.GetDashboardChartData).Methods("GET")
	router.Mux.HandleFunc(`/records/unknown`, devices.GetUnknown).Methods("GET")
	router.Mux.HandleFunc(`/records/trends`, devices.GetTrendData).Methods("GET")
	router.Mux.HandleFunc(`/records/investigate`, devices.InvestigateAddresses).Methods("POST")

	router.Mux.HandleFunc(`/records/device`, devices.SaveDevice).Methods("POST")
	router.Mux.HandleFunc(`/records/interface`, devices.SaveInterface).Methods("POST")
	router.Mux.HandleFunc(`/records/address`, devices.SaveAddress).Methods("POST")
	router.Mux.HandleFunc(`/records/hostname`, devices.SaveHostname).Methods("POST")

}

func (router *Router) settingsRoutes() {
	router.Mux.HandleFunc(`/settings/{userID:[0-9]+}`, settings.GetUserSettings).Methods("GET")
	router.Mux.HandleFunc(`/settings`, settings.GetAppSettings).Methods("GET")
	router.Mux.HandleFunc(`/setting/{userID:[0-9]+}`, settings.SaveUserSetting).Methods("POST")
	router.Mux.HandleFunc(`/setting}`, settings.SaveAppSetting).Methods("POST")
}

func (router *Router) paramRoutes() {
	router.Mux.HandleFunc(`/params`, params.GetAll).Methods("GET")
	router.Mux.HandleFunc(`/categories`, params.SaveCategory).Methods("POST")
	router.Mux.HandleFunc(`/categories/{ID:[0-9]+}`, params.DeleteCategory).Methods("DELETE")
	router.Mux.HandleFunc(`/statuses`, params.SaveStatus).Methods("POST")
	router.Mux.HandleFunc(`/statuses/{ID:[0-9]+}`, params.DeleteStatus).Methods("DELETE")
	router.Mux.HandleFunc(`/locations`, params.SaveLocation).Methods("POST")
	router.Mux.HandleFunc(`/locations/{ID:[0-9]+}`, params.DeleteLocation).Methods("DELETE")
	router.Mux.HandleFunc(`/maintainers`, params.SaveMaintainer).Methods("POST")
	router.Mux.HandleFunc(`/maintainers/{ID:[0-9]+}`, params.DeleteMaintainer).Methods("DELETE")
	router.Mux.HandleFunc(`/deviceTypes`, params.SaveDeviceType).Methods("POST")
	router.Mux.HandleFunc(`/deviceTypes/{ID:[0-9]+}`, params.DeleteDeviceType).Methods("DELETE")
	router.Mux.HandleFunc(`/architectures`, params.SaveArchitecture).Methods("POST")
	router.Mux.HandleFunc(`/architectures/{ID:[0-9]+}`, params.DeleteArchitecture).Methods("DELETE")
	router.Mux.HandleFunc(`/operatingSystems`, params.SaveOperatingSystem).Methods("POST")
	router.Mux.HandleFunc(`/operatingSystems/{ID:[0-9]+}`, params.DeleteOperatingSystem).Methods("DELETE")
	router.Mux.HandleFunc(`/vlans`, params.SaveVLAN).Methods("POST")
	router.Mux.HandleFunc(`/vlans/{ID:[0-9]+}`, params.DeleteVLAN).Methods("DELETE")
	router.Mux.HandleFunc(`/users`, params.SaveUser).Methods("POST")
	router.Mux.HandleFunc(`/users/{ID:[0-9]+}`, params.DeleteUser).Methods("DELETE")
	router.Mux.HandleFunc(`/interfaceTypes`, params.SaveInterfaceType).Methods("POST")
	router.Mux.HandleFunc(`/interfaceTypes/{ID:[0-9]+}`, params.DeleteInterfaceType).Methods("DELETE")
}

func (router *Router) staticRoutes() {
	router.Mux.HandleFunc(`/{filename:[a-zA-Z0-9=\.\/]+}`, control.Static).Methods("GET")
}
