package architectures

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(router *bunrouter.Router) {
	router.WithGroup("/api/architectures", func(group *bunrouter.Group) {
		group.POST(``, Save)         // "SaveArchitecture"
		group.DELETE(`/:ID`, Delete) // "DeleteArchitecture"
	})
}
