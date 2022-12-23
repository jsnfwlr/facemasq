package control

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(router *bunrouter.Router, group *bunrouter.Group) {
	// webDir, err := files.GetDir("web")
	// if err != nil {
	// 	panic(err)
	// }
	// fileServer := http.FileServer(http.Dir(webDir))
	// router.GET(`/`, bunrouter.HTTPHandler(fileServer)) // "ServeUI"
	router.GET(`/*filename`, Static) // "ServeStatic"

	group.GET(`/exit`, Exit)   // "APIExit"
	group.GET(`/state`, State) // "GetStatus
}
