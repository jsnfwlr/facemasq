package routes

import (
	"log"
	"net/http"
	"os"

	"facemasq/lib/files"

	"github.com/rs/cors"
	"github.com/urfave/negroni/v3"
)

var Port string

func init() {
	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "6135"
	}
}

func Run() (err error) {
	var rootDir string
	rootDir, _ = files.GetAppRoot()
	log.Printf("Starting API server at localhost:%s from  %s\n", Port, rootDir)

	server := negroni.New()
	router := BuildRoutes()

	if os.Getenv("NETMASK") == "" {
		corsControl := cors.AllowAll()
		server.Use(corsControl)
	}

	if os.Getenv("VERBOSE") != "" {
		server.Use(negroni.NewLogger())
	}

	server.UseHandler(router.Mux)
	err = http.ListenAndServe(":"+Port, server)
	return
}
