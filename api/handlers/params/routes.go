package params

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(group *bunrouter.Group) {
	group.WithGroup("/params", func(group *bunrouter.Group) {
		group.GET(``, GetAllParams) // "GetAllParams"
	})
}
