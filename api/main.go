package main

import (
	"facemasq/lib/db"
	"facemasq/lib/extensions"
	"facemasq/lib/logging"
	"facemasq/lib/network"
	"facemasq/lib/scans/iprange"
	"facemasq/routes"
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

	err := db.Connect()
	if err != nil {
		logging.Panic(err)
	}
	logging.Info("Connected: %+v", db.DBEngine)

	router := routes.Init()
	_, err = extensions.LoadExtensions(router.Bun)
	if err != nil {
		logging.Fatal("%v", err)
	}

	network.ShowNetworkSummary()

	if network.Target != "" {
		logging.Info("Active Net scan running every %v", iprange.Frequency)
		iprange.Schedule()
	}

	router.Run()
	if err != nil {
		logging.Fatal("%v", err)
	}
}
