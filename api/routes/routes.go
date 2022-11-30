package routes

import (
	"net/http"

	"github.com/gorilla/mux"

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
	"facemasq/lib/extensions"
)

type Router struct {
	Mux *mux.Router
}

func BuildRoutes() (router Router) {
	router = Router{
		Mux: mux.NewRouter(),
	}

	var routes []extensions.RouteDefinition
	routes = append(routes, control.GetRoutes()...)          // Control, UI, State Routes
	routes = append(routes, charts.GetRoutes()...)           // Chart Routes
	routes = append(routes, grids.GetRoutes()...)            // Grid Routes
	routes = append(routes, trends.GetRoutes()...)           // Trend Routes
	routes = append(routes, devices.GetRoutes()...)          // Device Routes
	routes = append(routes, interfaces.GetRoutes()...)       // Interface Routes
	routes = append(routes, addresses.GetRoutes()...)        // Addresse Routes
	routes = append(routes, hostnames.GetRoutes()...)        // Hostname Routes
	routes = append(routes, meta.GetRoutes()...)             // User & App Settings Routes
	routes = append(routes, params.GetRoutes()...)           // Combined Params Routes
	routes = append(routes, categories.GetRoutes()...)       // Category Routes
	routes = append(routes, statuses.GetRoutes()...)         // Status Routes
	routes = append(routes, locations.GetRoutes()...)        // Location Routes
	routes = append(routes, maintainers.GetRoutes()...)      // Maintainer Routes
	routes = append(routes, devicetypes.GetRoutes()...)      // DeviceType Routes
	routes = append(routes, architectures.GetRoutes()...)    // Architecture Routes
	routes = append(routes, operatingsystems.GetRoutes()...) // OperatingSystem Routes
	routes = append(routes, vlans.GetRoutes()...)            // VLAN Routes
	routes = append(routes, users.GetRoutes()...)            // User Routes
	routes = append(routes, interfacetypes.GetRoutes()...)   // InterfaceType Routes
	routes = append(routes, control.GetStaticRoutes()...)    // Static Routes

	// For now, plugin routes go last, to prevent them over riding the built in routes.
	// This may change in the future, but will require some sort of plugin security, with user confirmation and signed binaries
	routes = append(routes, extensions.Manager.GetRoutes()...)

	for r := range routes {
		if handlerFunc, ok := routes[r].Handler.(func(http.ResponseWriter, *http.Request)); ok {
			if routes[r].Methods != "WS" {
				router.Mux.HandleFunc(routes[r].Path, handlerFunc).Methods(routes[r].Methods).Name(routes[r].Name)
			}
		} else {
			handler := routes[r].Handler.(http.Handler)
			router.Mux.Handle(routes[r].Path, handler).Methods(routes[r].Methods).Name(routes[r].Name)
		}
	}

	return
}
