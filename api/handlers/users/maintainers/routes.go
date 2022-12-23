package maintainers

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(group *bunrouter.Group) {
	group.WithGroup("/maintainers", func(group *bunrouter.Group) {
		group.POST(``, Save)         // "SaveMaintainer"
		group.DELETE(`/:ID`, Delete) // "DeleteMaintainer"
	})
}
