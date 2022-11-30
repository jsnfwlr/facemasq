package main

import (
	"facemasq/lib/events"
	"facemasq/lib/extensions"
	"facemasq/lib/logging"
	"log"
	"testing"
	"time"
)

func TestLoadExtensions(t *testing.T) {
	var err error
	extensions.Manager, err = extensions.LoadPlugins()
	if err != nil {
		log.Fatalf("%v", err)
	}
	routes := extensions.Manager.GetRoutes()
	for r := range routes {
		logging.Info(routes[r].Name)
	}

	eventList, err := events.List()
	if err != nil {
		logging.Error("Error with event: %v", err)
	}

	for e := range eventList {
		logging.Info(eventList[e])
	}

	extensions.Manager.GetCoordinator().Listen("device:after:change", func(e events.Event) {
		logging.Info("main: %+v", e)
	})

	err = extensions.Manager.GetCoordinator().Emit(events.Event{Kind: "device:after:change"})
	if err != nil {
		logging.Error("Error with event: %v", err)
	}
	time.Sleep(3 * time.Second)

}
