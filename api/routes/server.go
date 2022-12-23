package routes

import (
	"net/http"
	"os"

	"facemasq/lib/files"
	"facemasq/lib/logging"

	"github.com/uptrace/bunrouter/extra/reqlog"
)

var Port string

func init() {
	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "6135"
	}
}

func (router *Router) Run() (err error) {
	var rootDir string
	rootDir, _ = files.GetAppRoot()
	logging.Info("Starting API server at localhost:%s from  %s", Port, rootDir)

	router.BuildRoutes()

	if os.Getenv("VERBOSE") != "" {
		router.Bun.Use(reqlog.NewMiddleware(
		// reqlog.WithEnabled(false),
		// reqlog.FromEnv("BUNDEBUG"),
		))
	}

	handler := http.Handler(router.Bun)
	handler = CORSHandler{Next: handler}
	err = http.ListenAndServe(":"+Port, handler)
	return
}
