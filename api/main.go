package main

import (
	"log"
	"os"

	"facemasq/lib/db"
	"facemasq/lib/logging"
	"facemasq/lib/netscan"
	"facemasq/lib/network"
	"facemasq/routes"
)

func init() {
	beVerbose := os.Getenv("VERBOSE")
	if beVerbose != "" {
		logging.Verbose = true
	}
}

func main() {
	log.Println("Running faceMasq as a daemon")
	err := db.Connect()
	if err != nil {
		panic(err)
	}

	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("Connected: %+v\n", db.DBEngine)

	network.ShowNetworkSummary()

	if os.Getenv("NETMASK") != "" {
		log.Printf("Active Net scan running every %v", netscan.Frequency)
		netscan.Schedule()
	}

	err = routes.Run()
	if err != nil {
		log.Fatalf("%v", err)
	}
}
