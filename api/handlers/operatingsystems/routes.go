package operatingsystems

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(group *bunrouter.Group) {
	group.WithGroup("/operatingSystems", func(group *bunrouter.Group) {
		group.POST(``, Save)         // "SaveOperatingSystem"
		group.DELETE(`/:ID`, Delete) // "DeleteOperatingSystem"
	})
}
