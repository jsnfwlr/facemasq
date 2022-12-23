package users

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(group *bunrouter.Group) {
	group.WithGroup("/users", func(group *bunrouter.Group) {
		group.POST(``, SaveUser)         // "SaveUser"
		group.DELETE(`/:ID`, DeleteUser) // "DeleteUser"
	})
}
