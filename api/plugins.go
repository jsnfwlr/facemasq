//go:build extensions

package main

import (
	"facemasq/lib/events"
	"facemasq/lib/extensions"
	"facemasq/lib/logging"
	"log"
	"time"
)

func main() {
	var err error
	extensions.Extensions, err = extensions.LoadPlugins()
	if err != nil {
		log.Fatalf("%v", err)
	}
	routes := extensions.Extensions.GetRoutes()
	for r := range routes {
		logging.System(routes[r].Name)
	}

	eventList, err := events.List()
	if err != nil {
		logging.Error("Error with event: %v", err)
	}

	for e := range eventList {
		logging.System(eventList[e])
	}

	extensions.Extensions.GetCoordinator().Listen("device:after:change", func(e events.Event) {
		logging.System("main: %+v", e)
	})

	err = extensions.Extensions.GetCoordinator().Emit(events.Event{Kind: "device:after:change"})
	if err != nil {
		logging.Error("Error with event: %v", err)
	}
	time.Sleep(3 * time.Second)

}
