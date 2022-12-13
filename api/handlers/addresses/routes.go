package addresses

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(router *bunrouter.Router) {
	router.WithGroup("/api/address", func(group *bunrouter.Group) {
		group.POST(``, Save)     // "SaveAddress"
		group.DELETE(``, Delete) // "DeleteAddress"
	})
}
