package grids

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(group *bunrouter.Group) {
	group.WithGroup("/grids", func(group *bunrouter.Group) {
		group.POST(`/dosier`, GetDeviceDosier)          // "GetDeviceDosier"
		group.GET(`/alldevices`, GetAllDevices)         // "GetAllDevices"
		group.GET(`/activedevices`, GetActiveDevices)   // "GetActiveDevices"
		group.GET(`/unknowndevices`, GetUnknownDevices) // "GetUnknownDevices"
	})
	group.WithGroup("/ws/grids", func(group *bunrouter.Group) {
		group.GET(`/grids`, GetRecentlyChangedDevices) // "WSRecentlyChangedDevices"
	})

}
