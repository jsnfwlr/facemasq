package vlans

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(router *bunrouter.Router) {
	router.WithGroup("/api/vlans", func(group *bunrouter.Group) {
		group.POST(``, Save)         // "SaveVLAN"
		group.DELETE(`/:ID`, Delete) // "DeleteVLAN"
	})
}
