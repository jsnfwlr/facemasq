package params

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(router *bunrouter.Router) {
	router.WithGroup("/api/params", func(group *bunrouter.Group) {
		group.GET(``, GetAllParams) // "GetAllParams"
	})
}
