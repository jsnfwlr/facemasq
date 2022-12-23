package statuses

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(group *bunrouter.Group) {
	group.WithGroup("/statuses", func(group *bunrouter.Group) {
		group.POST(`/api/statuses`, Save)         // "SaveStatus"
		group.DELETE(`/api/statuses/:ID`, Delete) // "DeleteStatus"
	})
}
