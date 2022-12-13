package hostnames

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(router *bunrouter.Router) {
	router.WithGroup("/api/hostname", func(group *bunrouter.Group) {
		group.POST(``, Save)     // "SaveHostname"
		group.DELETE(``, Delete) // "DeleteHostname"
	})
}
