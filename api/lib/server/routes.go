package server

import (
	"facemasq/handlers/addresses"
	"facemasq/handlers/architectures"
	"facemasq/handlers/categories"
	"facemasq/handlers/charts"
	"facemasq/handlers/control"
	"facemasq/handlers/devices"
	"facemasq/handlers/devicetypes"
	"facemasq/handlers/grids"
	"facemasq/handlers/hostnames"
	"facemasq/handlers/interfaces"
	"facemasq/handlers/interfacetypes"
	"facemasq/handlers/locations"
	"facemasq/handlers/meta"
	"facemasq/handlers/operatingsystems"
	"facemasq/handlers/params"
	"facemasq/handlers/statuses"
	"facemasq/handlers/trends"
	"facemasq/handlers/users"
	"facemasq/handlers/users/maintainers"
	"facemasq/handlers/vlans"
	"facemasq/lib/server/httperror"

	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

type RouterConfig struct {
	Bun *bunrouter.Router
}

func Init() (router *RouterConfig) {
	router = &RouterConfig{
		Bun: bunrouter.New(
			bunrouter.Use(reqlog.NewMiddleware()),
		),
	}
	return
}

func (router *RouterConfig) BuildRoutes() {

	group := router.Bun.NewGroup("/api", bunrouter.Use(httperror.ErrorHandler))

	charts.GetRoutes(group)              // Chart Routes
	grids.GetRoutes(group)               // Grid Routes
	trends.GetRoutes(group)              // Trend Routes
	devices.GetRoutes(group)             // Device Routes
	interfaces.GetRoutes(group)          // Interface Routes
	addresses.GetRoutes(group)           // Addresse Routes
	hostnames.GetRoutes(group)           // Hostname Routes
	meta.GetRoutes(group)                // User & App Settings Routes
	params.GetRoutes(group)              // Combined Params Routes
	categories.GetRoutes(group)          // Category Routes
	statuses.GetRoutes(group)            // Status Routes
	locations.GetRoutes(group)           // Location Routes
	maintainers.GetRoutes(group)         // Maintainer Routes
	devicetypes.GetRoutes(group)         // DeviceType Routes
	architectures.GetRoutes(group)       // Architecture Routes
	operatingsystems.GetRoutes(group)    // OperatingSystem Routes
	vlans.GetRoutes(group)               // VLAN Routes
	users.GetRoutes(group)               // User Routes
	interfacetypes.GetRoutes(group)      // InterfaceType Routes
	control.GetRoutes(router.Bun, group) // Control, UI, State Routes

	// For now, plugin routes go last, to prevent them over riding the built in routes.
	// This may change in the future, but will require some sort of plugin security, with user confirmation and signed binaries
	// if extensions.Manager.HasRoutes() {
	// 	extensions.Manager.GetRoutes(group)
	// }

	// return
}
