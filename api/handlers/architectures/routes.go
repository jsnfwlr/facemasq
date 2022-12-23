package architectures

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(group *bunrouter.Group) {
	group.WithGroup("/architectures", func(group *bunrouter.Group) {
		group.POST(``, Save)         // "SaveArchitecture"
		group.DELETE(`/:ID`, Delete) // "DeleteArchitecture"
	})
}
