package devicetypes

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(group *bunrouter.Group) {
	group.WithGroup("/deviceTypes", func(group *bunrouter.Group) {
		group.POST(``, Save)         // "SaveDeviceType"
		group.DELETE(`/:ID`, Delete) // "DeleteDeviceType"
	})
}
