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

var (
  port string
  dbEngine db.Engine
)

func init() {
	var err error
	port = os.Getenv("PORT")
	if port == "" {
		port = "6135"
	}
	dbEngineType := os.Getenv("DBENGINE")
	switch dbEngineType {
		case "", "sqlite":
			dbEngine = db.Engine{
				EngineType: "sqlite",
				DataRoot: "../data",
				DataFile: "network.sqlite",
			}
//		case "mysql", "mariadb":
//			log.Fatal("Not yet supported")
//		case "postgresql":
//			log.Fatal("Not yet supported")
		default:
			log.Fatal("Not yet supported")
	}
	if err != nil {
	    log.Fatalf("%v", err.Error())
	}
}

func main() {
	var err error
	log.Println("Running faceMasq as a daemon")
	db.Conn, err = dbEngine.Connect()
	if err != nil {
		log.Fatalf("%v", err)
	}
//	db.Conn = dbEngine.Conn
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
