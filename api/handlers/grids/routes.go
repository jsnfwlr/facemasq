package grids

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(router *bunrouter.Router) {
	router.WithGroup("/api/grids", func(group *bunrouter.Group) {
		group.POST(`/dosier`, GetDeviceDosier)          // "GetDeviceDosier"
		group.GET(`/alldevices`, GetAllDevices)         // "GetAllDevices"
		group.GET(`/activedevices`, GetActiveDevices)   // "GetActiveDevices"
		group.GET(`/unknowndevices`, GetUnknownDevices) // "GetUnknownDevices"
	})
	// 	router.WS(`/ws/grids/changed`, GetRecentlyChangedDevices) // "WSRecentlyChangedDevices"
}
