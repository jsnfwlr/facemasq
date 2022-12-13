package devices

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(router *bunrouter.Router) {
	router.WithGroup("/api/device", func(group *bunrouter.Group) {
		group.POST(``, Save)         // "SaveDevice"
		group.DELETE(`/:ID`, Delete) // "DeleteDevice"
		// group.DELETE(``, Delete) // "DeleteDevice"
	})
}
