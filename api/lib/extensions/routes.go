package extensions

type RouteDefinition struct {
	Path    string
	Handler Handler
	Methods string
	Name    string
}

type Handler interface{}
