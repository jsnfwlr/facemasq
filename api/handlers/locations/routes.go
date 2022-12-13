package locations

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(router *bunrouter.Router) {
	router.WithGroup("/api/locations", func(group *bunrouter.Group) {
		group.POST(``, Save)         // "SaveLocation"
		group.DELETE(`/:ID`, Delete) // "DeleteLocation"
	})
}
