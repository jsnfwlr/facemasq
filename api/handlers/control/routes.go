package control

import (
	"github.com/uptrace/bunrouter"
)

func GetRoutes(router *bunrouter.Router) {
	// webDir, err := files.GetDir("web")
	// if err != nil {
	// 	panic(err)
	// }
	// fileServer := http.FileServer(http.Dir(webDir))
	// router.GET(`/`, bunrouter.HTTPHandler(fileServer)) // "ServeUI"
	router.GET(`/*filename`, Static) // "ServeStatic"

	router.WithGroup("/api", func(group *bunrouter.Group) {
		group.GET(`/exit`, Exit)   // "APIExit"
		group.GET(`/state`, State) // "GetStatus"
	})
}
