package main

import (
	"log"
	"net/http"
	"os"

	"github.com/urfave/negroni"

	"github.com/jsnfwlr/facemasq/api/lib/db"
	"github.com/jsnfwlr/facemasq/api/lib/netscan"
	"github.com/jsnfwlr/facemasq/api/routes"
)

var port string

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "6135"
	}
}

func main() {
	log.Println("Starting facemasq")
	err := db.Connect("../data/", "network.sqlite")
	if err != nil {
		log.Fatalf("%v", err)
	}

	netscan.Schedule()

	err = runServer()
	if err != nil {
		log.Fatalf("%v", err)
	}
}

func runServer() (err error) {
	cwd, _ := os.Getwd()
	log.Printf("Starting API server in %s\n", cwd)

	server := negroni.New()
	router := routes.BuildRoutes()

	server.UseHandler(router.Mux)
	err = http.ListenAndServe(":"+port, server)
	return
}
