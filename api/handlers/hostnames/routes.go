package hostnames

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(group *bunrouter.Group) {
	group.WithGroup("/hostname", func(group *bunrouter.Group) {
		group.POST(``, Save)     // "SaveHostname"
		group.DELETE(``, Delete) // "DeleteHostname"
	})
}
