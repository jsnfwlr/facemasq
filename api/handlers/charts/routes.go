package charts

import "facemasq/lib/extensions"

func GetRoutes() (routes []extensions.RouteDefinition) {
	routes = []extensions.RouteDefinition{
		{Path: `/api/charts/devicesovertime`, Handler: GetDevicesOverTime, Methods: "GET", Name: "GetDashboardChartData"},
	}
	return
}
