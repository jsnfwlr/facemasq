package charts

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(group *bunrouter.Group) {
	group.WithGroup("/charts", func(group *bunrouter.Group) {
		group.GET(`/devicesovertime`, GetDevicesOverTime) // "GetDashboardChartData"
	})
}
