package routes

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
	"facemasq/lib/httperror"

	"github.com/uptrace/bunrouter"
)

type Router struct {
	Bun *bunrouter.Router
}

func Init() (router *Router) {
	router = &Router{
		Bun: bunrouter.New(),
	}
	return
}

func (router *Router) BuildRoutes() {

	router.Bun.WithGroup("/api", func(group *bunrouter.Group) {
		group.Use(httperror.Handler)
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
	})

	// For now, plugin routes go last, to prevent them over riding the built in routes.
	// This may change in the future, but will require some sort of plugin security, with user confirmation and signed binaries
	// if extensions.Manager.HasRoutes() {
	// 	extensions.Manager.GetRoutes(group)
	// }

	// return
}
