package interfaces

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(group *bunrouter.Group) {
	group.WithGroup("/interface", func(group *bunrouter.Group) {
		group.POST(``, Save)         // "SaveInterface"
		group.DELETE(`/:ID`, Delete) // "DeleteInterface"
	})
}
