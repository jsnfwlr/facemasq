package maintainers

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(router *bunrouter.Router) {
	router.WithGroup("/api/maintainers", func(group *bunrouter.Group) {
		group.POST(``, Save)         // "SaveMaintainer"
		group.DELETE(`/:ID`, Delete) // "DeleteMaintainer"
	})
}
