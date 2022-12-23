package interfacetypes

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(group *bunrouter.Group) {
	group.WithGroup("/interfaceTypes", func(group *bunrouter.Group) {
		group.POST(``, SaveInterfaceType)         // "SaveInterfaceType"
		group.DELETE(`/:ID`, DeleteInterfaceType) // "DeleteInterfaceType"
	})
}
