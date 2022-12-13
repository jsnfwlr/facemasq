package routes

import (
	"testing"
)

// type table struct {
// 	name   string
// 	method string
// }

func TestRoutes(t *testing.T) {
	// router := BuildRoutes()

	// list := []table{
	// 	{name: "APIExit", method: "GET"},
	// 	{name: "Exit", method: "GET"},
	// 	{name: "WriteDHCPConfigGet", method: "GET"},
	// 	{name: "WriteDHCPConfigPut", method: "PUT"},
	// 	{name: "WriteDNSConfigGet", method: "GET"},
	// 	{name: "WriteDNSConfigPut", method: "PUT"},
	// 	{name: "GetAllDevices", method: "GET"},
	// 	{name: "GetActiveDevices", method: "GET"},
	// 	{name: "GetDashboardChartData", method: "GET"},
	// 	{name: "GetUnknownDevices", method: "GET"},
	// 	{name: "GetDeviceTrendData", method: "GET"},
	// 	{name: "InvestigateAddresses", method: "POST"},
	// 	{name: "SaveDevice", method: "POST"},
	// 	{name: "SaveInterface", method: "POST"},
	// 	{name: "SaveAddress", method: "POST"},
	// 	{name: "SaveHostname", method: "POST"},
	// 	{name: "GetUserSettings", method: "GET"},
	// 	{name: "GetAppSettings", method: "GET"},
	// 	{name: "SaveUserSetting", method: "POST"},
	// 	{name: "SaveAppSetting", method: "POST"},
	// 	{name: "GetStatus", method: "GET"},
	// 	{name: "GetAllParams", method: "GET"},
	// 	{name: "SaveCategory", method: "POST"},
	// 	{name: "DeleteCategory", method: "DELETE"},
	// 	{name: "SaveStatus", method: "POST"},
	// 	{name: "DeleteStatus", method: "DELETE"},
	// 	{name: "SaveLocation", method: "POST"},
	// 	{name: "DeleteLocation", method: "DELETE"},
	// 	{name: "SaveMaintainer", method: "POST"},
	// 	{name: "DeleteMaintainer", method: "DELETE"},
	// 	{name: "SaveDeviceType", method: "POST"},
	// 	{name: "DeleteDeviceType", method: "DELETE"},
	// 	{name: "SaveArchitecture", method: "POST"},
	// 	{name: "DeleteArchitecture", method: "DELETE"},
	// 	{name: "SaveOperatingSystem", method: "POST"},
	// 	{name: "DeleteOperatingSystem", method: "DELETE"},
	// 	{name: "SaveVLAN", method: "POST"},
	// 	{name: "DeleteVLAN", method: "DELETE"},
	// 	{name: "SaveUser", method: "POST"},
	// 	{name: "DeleteUser", method: "DELETE"},
	// 	{name: "SaveInterfaceType", method: "POST"},
	// 	{name: "DeleteInterfaceType", method: "DELETE"},
	// 	{name: "ServeUI", method: "GET"},
	// 	{name: "ServeStatic", method: "GET"},
	// }
	// for i := range list {
	// 	err := router.Mux.Get(list[i].name).GetError()
	// 	if err != nil {
	// 		t.Errorf("Error building route: %v", err)
	// 		continue
	// 	}
	// 	result, err := router.Mux.Get(list[i].name).GetMethods()
	// 	if err != nil {
	// 		t.Errorf("Could not get methods for %s", list[i].name)
	// 	}
	// 	if result[0] != list[i].method {
	// 		t.Errorf("Route %s using incorrect method %s != %s", list[i].name, result, list[i].method)
	// 	}
	// }
}
