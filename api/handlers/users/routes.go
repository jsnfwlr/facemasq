package users

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(router *bunrouter.Router) {
	router.WithGroup("/api/users", func(group *bunrouter.Group) {
		group.POST(``, SaveUser)         // "SaveUser"
		group.DELETE(`/:ID`, DeleteUser) // "DeleteUser"
	})
}
