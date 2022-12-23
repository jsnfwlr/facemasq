package vlans

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(group *bunrouter.Group) {
	group.WithGroup("/vlans", func(group *bunrouter.Group) {
		group.POST(``, Save)         // "SaveVLAN"
		group.DELETE(`/:ID`, Delete) // "DeleteVLAN"
	})
}
