package routes

import (
	"github.com/uptrace/bunrouter"

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

	control.GetRoutes(router.Bun)          // Control, UI, State Routes
	charts.GetRoutes(router.Bun)           // Chart Routes
	grids.GetRoutes(router.Bun)            // Grid Routes
	trends.GetRoutes(router.Bun)           // Trend Routes
	devices.GetRoutes(router.Bun)          // Device Routes
	interfaces.GetRoutes(router.Bun)       // Interface Routes
	addresses.GetRoutes(router.Bun)        // Addresse Routes
	hostnames.GetRoutes(router.Bun)        // Hostname Routes
	meta.GetRoutes(router.Bun)             // User & App Settings Routes
	params.GetRoutes(router.Bun)           // Combined Params Routes
	categories.GetRoutes(router.Bun)       // Category Routes
	statuses.GetRoutes(router.Bun)         // Status Routes
	locations.GetRoutes(router.Bun)        // Location Routes
	maintainers.GetRoutes(router.Bun)      // Maintainer Routes
	devicetypes.GetRoutes(router.Bun)      // DeviceType Routes
	architectures.GetRoutes(router.Bun)    // Architecture Routes
	operatingsystems.GetRoutes(router.Bun) // OperatingSystem Routes
	vlans.GetRoutes(router.Bun)            // VLAN Routes
	users.GetRoutes(router.Bun)            // User Routes
	interfacetypes.GetRoutes(router.Bun)   // InterfaceType Routes

	// For now, plugin routes go last, to prevent them over riding the built in routes.
	// This may change in the future, but will require some sort of plugin security, with user confirmation and signed binaries
	// if extensions.Manager.HasRoutes() {
	// 	extensions.Manager.GetRoutes(router.Bun)
	// }

	// return
}
