package extensions

type RouteDefinition struct {
	Grouped    []GroupedRoute
	Individual []IndividualRoute
}

type IndividualRoute struct {
	Path    string
	Handler Handler
	Methods string
	Name    string
}

type GroupedRoute struct {
	Base   string
	Routes []IndividualRoute
}

type Handler interface{}
