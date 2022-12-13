package categories

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(router *bunrouter.Router) {
	router.WithGroup("/api/categories", func(group *bunrouter.Group) {
		group.POST(``, Save)         // "SaveCategory"
		group.DELETE(`/:ID`, Delete) // "DeleteCategory"
	})
}
