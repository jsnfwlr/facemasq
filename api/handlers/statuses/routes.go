package statuses

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(router *bunrouter.Router) {
	router.WithGroup("/api/statuses", func(group *bunrouter.Group) {
		group.POST(`/api/statuses`, Save)         // "SaveStatus"
		group.DELETE(`/api/statuses/:ID`, Delete) // "DeleteStatus"
	})
}
