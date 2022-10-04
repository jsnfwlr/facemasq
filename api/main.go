package main

import (
	"facemasq/lib/db"
	"facemasq/lib/logging"
	"facemasq/lib/netscan"
	"facemasq/lib/network"
	"facemasq/routes"
)

func main() {
	logging.Processln("Running faceMasq as a daemon")
	err := db.Connect()
	if err != nil {
		logging.Panic(err)
	}

	if err != nil {
		logging.Fatalf("%v", err)
	}
	logging.Processf("Connected: %+v\n", db.DBEngine)

	network.ShowNetworkSummary()

	if network.Target != "" {
		logging.Processf("Active Net scan running every %v", netscan.Frequency)
		netscan.Schedule()
	}

	err = routes.Run()
	if err != nil {
		logging.Fatalf("%v", err)
	}
}
