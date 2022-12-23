package categories

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(group *bunrouter.Group) {
	group.WithGroup("/categories", func(group *bunrouter.Group) {
		group.POST(``, Save)         // "SaveCategory"
		group.DELETE(`/:ID`, Delete) // "DeleteCategory"
	})
}
