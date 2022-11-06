//go:build !extensions

package main

import (
	"facemasq/lib/db"
	"facemasq/lib/logging"
	"facemasq/lib/network"
	"facemasq/lib/scans/iprange"
	"facemasq/routes"
)

func main() {
	logging.= logging.New("", "")
	logging.System("Running faceMasq as a daemon")
	err := db.Connect()
	if err != nil {
		logging.Panic(err)
	}

	if err != nil {
		logging.Fatal("%v", err)
	}
	logging.System("Connected: %+v", db.DBEngine)

	network.ShowNetworkSummary()

	if network.Target != "" {
		logging.System("Active Net scan running every %v", iprange.Frequency)
		iprange.Schedule()
	}

	err = routes.Run()
	if err != nil {
		logging.Fatal("%v", err)
	}
}
