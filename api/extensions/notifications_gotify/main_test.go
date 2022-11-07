package main

import (
	"facemasq/lib/events"
	"facemasq/lib/events/mapi"
	"testing"
)

func TestMessage(t *testing.T) {
	// payload := make(mapi.Mapi)
	payload := mapi.Mapi{
		"title":    "faceMasq: Intruder Detected",
		"message":  "MAC: 00:08:48:BA:D1:31\nIP: 192.168.1.5\nPrevious Sessions: 32\nMost Recent: 2006-01-02 15:04",
		"priority": 1,
	}

	e := events.Event{
		Kind:    "send:notifcation:intruder",
		Message: "",
		Payload: payload,
	}
	SendMessage(e)

}
