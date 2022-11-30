package main

import (
	"facemasq/lib/db"
	"facemasq/lib/extensions"
	"facemasq/lib/logging"
	"facemasq/lib/network"
	"facemasq/lib/scans/iprange"
	"facemasq/routes"
)

func main() {
	logging.Info("Running faceMasq as a daemon")

	err := db.Connect()
	if err != nil {
		logging.Panic(err)
	}
	logging.Info("Connected: %+v", db.DBEngine)

	_, err = extensions.LoadPlugins()
	if err != nil {
		logging.Fatal("%v", err)
	}

	network.ShowNetworkSummary()

	if network.Target != "" {
		logging.Info("Active Net scan running every %v", iprange.Frequency)
		iprange.Schedule()
	}

	err = routes.Run()
	if err != nil {
		logging.Fatal("%v", err)
	}
}
