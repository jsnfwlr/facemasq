package meta

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(group *bunrouter.Group) {
	group.WithGroup("/setting", func(group *bunrouter.Group) {
		group.POST(`/:UserID`, SaveUserSetting) // "SaveUserSetting"
		group.POST(``, SaveAppSetting)          // "SaveAppSetting"
	})

	group.WithGroup("/settings", func(group *bunrouter.Group) {
		group.GET(`/:UserID`, GetUserSettings) // "GetUserSettings"
		group.GET(``, GetAppSettings)          // "GetAppSettings"
	})
}
