package locations

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(group *bunrouter.Group) {
	group.WithGroup("/locations", func(group *bunrouter.Group) {
		group.POST(``, Save)         // "SaveLocation"
		group.DELETE(`/:ID`, Delete) // "DeleteLocation"
	})
}
