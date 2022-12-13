package interfacetypes

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(router *bunrouter.Router) {
	router.WithGroup("/api/interfaceTypes", func(group *bunrouter.Group) {
		group.POST(``, SaveInterfaceType)         // "SaveInterfaceType"
		group.DELETE(`/:ID`, DeleteInterfaceType) // "DeleteInterfaceType"
	})
}
