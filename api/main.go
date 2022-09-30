package main

import (
	"os"
	"strconv"

	"facemasq/lib/db"
	"facemasq/lib/logging"
	"facemasq/lib/netscan"
	"facemasq/lib/network"
	"facemasq/routes"
)

func init() {
	var err error
	logging.Verbosity, err = strconv.Atoi(os.Getenv("VERBOSE"))
	if err != nil {
		logging.Verbosity = 0
	}
}

func main() {
	logging.Processln("Running faceMasq as a daemon")
	err := db.Connect()
	if err != nil {
		panic(err)
	}

	if err != nil {
		logging.Fatalf("%v", err)
	}
	logging.Processf("Connected: %+v\n", db.DBEngine)

	network.ShowNetworkSummary()

	if os.Getenv("NETMASK") != "" {
		logging.Processf("Active Net scan running every %v", netscan.Frequency)
		netscan.Schedule()
	}

	err = routes.Run()
	if err != nil {
		logging.Fatalf("%v", err)
	}
}
