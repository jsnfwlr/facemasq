package addresses

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(group *bunrouter.Group) {
	group.WithGroup("/address", func(group *bunrouter.Group) {
		group.POST(``, Save)     // "SaveAddress"
		group.DELETE(``, Delete) // "DeleteAddress"
	})
}
