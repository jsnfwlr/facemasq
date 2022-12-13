package trends

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(router *bunrouter.Router) {
	router.WithGroup("/api/trends", func(group *bunrouter.Group) {
		group.GET(`/connections`, GetConnectionTrends) // "GetConnectionTrends"
	})
}
