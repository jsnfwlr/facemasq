package meta

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(router *bunrouter.Router) {
	router.WithGroup("/api/setting", func(group *bunrouter.Group) {
		group.POST(`/:UserID`, SaveUserSetting) // "SaveUserSetting"
		group.POST(``, SaveAppSetting)          // "SaveAppSetting"
	})

	router.WithGroup("/api/settings", func(group *bunrouter.Group) {
		group.GET(`/:UserID`, GetUserSettings) // "GetUserSettings"
		group.GET(``, GetAppSettings)          // "GetAppSettings"
	})
}
