package charts

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(router *bunrouter.Router) {
	router.WithGroup("/api/charts", func(group *bunrouter.Group) {
		group.GET(`/devicesovertime`, GetDevicesOverTime) // "GetDashboardChartData"
	})
}
