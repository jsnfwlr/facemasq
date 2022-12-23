package trends

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(group *bunrouter.Group) {
	group.WithGroup("/trends", func(group *bunrouter.Group) {
		group.GET(`/connections`, GetConnectionTrends) // "GetConnectionTrends"
	})
}
