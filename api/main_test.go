package main

import (
	"facemasq/lib/events"
	"facemasq/lib/extensions"
	"facemasq/lib/logging"
	"log"
	"testing"
	"time"

	"github.com/uptrace/bunrouter"
)

func TestLoadExtensions(t *testing.T) {
	var err error
	extensions.Manager, err = extensions.LoadExtensions()
	if err != nil {
		log.Fatalf("%v", err)
	}
	router := bunrouter.New()
	extensions.Manager.GetRoutes(router)

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
