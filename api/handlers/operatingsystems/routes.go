package operatingsystems

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(router *bunrouter.Router) {
	router.WithGroup("/api/operatingSystems", func(group *bunrouter.Group) {
		group.POST(``, Save)         // "SaveOperatingSystem"
		group.DELETE(`/:ID`, Delete) // "DeleteOperatingSystem"
	})
}
