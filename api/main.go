package main

import (
	"facemasq/lib/db"
	"facemasq/lib/extensions"
	"facemasq/lib/logging"
	"facemasq/lib/network"
	"facemasq/lib/scans/iprange"
	"facemasq/lib/server"
	"facemasq/lib/translate"
	"flag"
	"time"
)

var Sleep bool

func init() {
	flag.BoolVar(&Sleep, "sleep", false, "Just sleep, don't start any goroutines")

	flag.Parse()
}

func main() {
	if Sleep {
		logging.Info("Running faceMasq as a sleeper")
		for {
			time.Sleep(time.Hour * 24)
		}
	}
	logging.Info("Running faceMasq as a daemon")

	// Initialize i18n engine
	err := translate.Start()
	if err != nil {
		logging.Panic(err)
	}

	// create DB connectio n
	err = db.Connect()
	if err != nil {
		logging.Panic(err)
	}
	logging.Info("Connected: %+v", db.DBEngine)

	// Prepare Routes
	server.Router = server.Init()

	// Load extensions
	_, err = extensions.LoadExtensions(server.Router.Bun)
	if err != nil {
		logging.Fatal("%v", err)
	}

	// Output Network summary
	network.ShowNetworkSummary()

	// Schedule the active ARP request scans if the network.Target is defined
	if network.Target != "" {
		logging.Info("Active Net scan running every %v", iprange.Frequency)
		iprange.Schedule()
	}

	// Run the API server
	server.Router.Run()
	if err != nil {
		logging.Fatal("%v", err)
	}
}
