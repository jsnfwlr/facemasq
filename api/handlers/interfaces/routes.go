package interfaces

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(router *bunrouter.Router) {
	router.WithGroup("/api/interface", func(group *bunrouter.Group) {
		group.POST(``, Save)         // "SaveInterface"
		group.DELETE(`/:ID`, Delete) // "DeleteInterface"
	})
}
