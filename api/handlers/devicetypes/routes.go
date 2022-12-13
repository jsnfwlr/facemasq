package devicetypes

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(router *bunrouter.Router) {
	router.WithGroup("/api/deviceTypes", func(group *bunrouter.Group) {
		group.POST(``, Save)         // "SaveDeviceType"
		group.DELETE(`/:ID`, Delete) // "DeleteDeviceType"
	})
}
