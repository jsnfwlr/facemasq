package devices

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(group *bunrouter.Group) {
	group.WithGroup("/device", func(group *bunrouter.Group) {
		group.POST(``, Save)         // "SaveDevice"
		group.DELETE(`/:ID`, Delete) // "DeleteDevice"
		// group.DELETE(``, Delete) // "DeleteDevice"
	})
}
